import requests
from flask import Flask, request, make_response
from config import MAIN_DATABASE_URL, VOTING_DATABASE_URL, MAIL_SERVICE_URL, PROMOTION_URL,FRONTEND_SERVICE_URL
from models import db, Voting, Vote, User, Blacklist
from sqlalchemy.exc import IntegrityError
from apscheduler.schedulers.background import BackgroundScheduler
from apscheduler.triggers.cron import CronTrigger
from pytz import timezone
from flask_cors import CORS


app = Flask(__name__)

CORS(app, supports_credentials=True,origins = [FRONTEND_SERVICE_URL])

app.config['SQLALCHEMY_DATABASE_URI'] = MAIN_DATABASE_URL

app.config['SQLALCHEMY_BINDS'] = {
    'voting-db': VOTING_DATABASE_URL
}

db.init_app(app)

scheduler = BackgroundScheduler()

@app.route("/api/voting/create_voting", methods=['POST'])
def create_voting():
    data = request.get_json()
    voting_subject = data.get("voting_subject")
    voting_category = data.get("voting_category")
    subject_data = User.query.filter_by(user_display_name=voting_subject).first()

    if subject_data is None:
        return make_response({"error": "Selected user not found"}, 404)

    subject_status = subject_data.user_status

    voting_exist = Voting.query.filter(
            (Voting.voting_subject == voting_subject) & (Voting.voting_category == voting_category)
            ).first()
    
    if voting_exist is not None:
        return make_response({"error": "This voting already exist"}, 409)

    if subject_status == "gold" and voting_category == "promote":
        return make_response({"message": "Current user already has gold status"}, 304)

    new_voting = Voting(
        voting_subject=voting_subject,
        voting_category=voting_category,
        subject_status=subject_status
        )
    db.session.add(new_voting)
    db.session.commit()
    return make_response({"message": "New voting was created"}, 200)


@app.route("/api/voting/vote", methods=['POST'])
def add_vote():
    data = request.get_json()
    voting_id = data.get("voting_id")
    voter_id = data.get("voter_id")
    try:
        added_vote = Vote(voting_id=voting_id, voter_id=voter_id)
        db.session.add(added_vote)
        db.session.commit()
    except IntegrityError:
        db.session.rollback()
        return make_response({"error": "Vote already exist"}, 409)
    return make_response({"message": "Your vote was added"}, 200)


@app.route("/api/voting/get_votings", methods=['POST'])
def get_all_votings():
    data = request.get_json()
    user_id = data.get("user_id")
    viewer_status = data.get("status")

    if viewer_status == "bronze":
        all_votings = Voting.query.filter_by(voting_category="exclude").all()
    elif viewer_status == "silver":
        all_votings = Voting.query.filter(
            (Voting.subject_status == "bronze") | (Voting.voting_category == "exclude")
            ).all()
    elif viewer_status == "gold":
        all_votings = Voting.query.all()
    else:
        return make_response({"error": "Current user status doesn't exist"}, 404)
    votings_list = []

    for voting in all_votings:
        voting_id = voting.voting_id
        category = voting.voting_category
        username = voting.voting_subject
        subject_status = voting.subject_status

        if db.session.get(Vote, (voting_id, user_id)) is not None:
            is_approved = True
        else:
            is_approved = False
        voting_info = {
            "voting_id": voting_id,
            "category": category,
            "username": username,
            "is_approved": is_approved,
            "subject_status": subject_status
        }
        votings_list.append(voting_info)
    return make_response({"votings": votings_list}, 200)


def promote(user_name: str, email: str):
    body = {
        "userdisplayname": user_name
    }
    response = requests.post(PROMOTION_URL, json=body)

    if int(response.status_code) == 200:
        notification = {
            "dest": [email],
            "subject": "Promotion",
            "body": "Congrats! Your was promoted to higher status."
        }
        requests.post(MAIL_SERVICE_URL, json=notification)


# This endpoint gives ability to trigger function without scheduler
# @app.route("/summarize", methods=['GET'])
def summarize_votings():
    with app.app_context():
        all_voters_count = len(list(User.query.all()))
        bronze_voters_count = len(list(User.query.filter(
            (User.user_status == "silver") | (User.user_status == "gold")
            ).all()))
        silver_voters_count = len(list(User.query.filter_by(user_status="gold").all()))

        if bronze_voters_count == 0:
            bronze_voters_count = 1
        if silver_voters_count == 0:
            silver_voters_count = 1

        votings_data = Voting.query.all()
        for voting_data in votings_data:
            voting_subject = str(voting_data.voting_subject)
            subject_data = User.query.filter_by(user_display_name=voting_subject).first()
            voting_category = voting_data.voting_category
            subject_status = voting_data.subject_status
            votes_count = len(list(voting_data.votes))

            if voting_category == "exclude":
                if (votes_count / all_voters_count) * 100 > 80:
                    try:
                        email_to_ban = Blacklist(banned_email=subject_data.user_email)
                        db.session.add(email_to_ban)
                        db.session.commit()
                    except IntegrityError:
                        db.session.rollback()
            elif voting_category == "promote":
                if subject_status == "bronze":
                    if (votes_count / bronze_voters_count) * 100 >= 51:
                        promote(voting_subject, subject_data.user_email)
                elif subject_status == "silver":
                    if (votes_count / silver_voters_count) * 100 >= 51:
                        promote(voting_subject, subject_data.user_email)

        db.session.execute(db.delete(Voting))
        db.session.commit()


if __name__ == "__main__":
    kyiv_tz = timezone('Europe/Kyiv')
    scheduler.add_job(
        func=summarize_votings,
        trigger=CronTrigger(hour=23, minute=59, timezone=kyiv_tz),
        id='votings_sumarize'
        )
    scheduler.start()
    app.run(host="0.0.0.0", port=4242)