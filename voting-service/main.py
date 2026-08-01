from flask import Flask, request, make_response
from config import MAIN_DATABASE_URL, VOTING_DATABASE_URL
from models import db, Voting, Vote, User
from sqlalchemy.exc import IntegrityError

app = Flask(__name__)

app.config['SQLALCHEMY_DATABASE_URI'] = MAIN_DATABASE_URL

app.config['SQLALCHEMY_BINDS'] = {
    'voting-db': VOTING_DATABASE_URL
}

db.init_app(app)

@app.route("/create_voting", methods=['POST'])
def create_voting():
    data = request.get_json()
    voting_subject = data.get("voting_subject")
    voting_category = data.get("voting_category")
    subject_data = User.query.filter_by(user_display_name=voting_subject).first()
    subject_status = subject_data.user_status

    if subject_status == "gold" and voting_category == "promote":
        return make_response({"message": "Current user already has gold status"}, 200)

    new_voting = Voting(
        voting_subject=voting_subject,
        voting_category=voting_category,
        subject_status=subject_status
        )
    db.session.add(new_voting)
    db.session.commit()
    return make_response({"message": "New voting was created"}, 200)


@app.route("/vote", methods=['POST'])
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
        return make_response({"error": "Vote already exist"}, 400)
    return make_response({"message": "Your vote was added"}, 200)


@app.route("/get_votings", methods=['POST'])
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
        return make_response({"message": "Current user status doesn't exist"}, 200)
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


@app.route("/sumarize_votings", methods=['GET']) # temporary for testing
def sumarize_votings():
    votings_list = []
    votings_data = Voting.query.all()
    for voting_data in votings_data:
        voting_id = voting_data.voting_id
        voting_subject = voting_data.voting_subject
        voting_category = voting_data.voting_category
        subject_status = voting_data.subject_status
        votes_count = len(list(voting_data.votes))

        if voting_category == "exclude":
            pass

        voting_info = {
            "voting_id": voting_id,
            "voting_subject": voting_subject,
            "voting_category": voting_category,
            "subject_status": subject_status,
            "votes_count": votes_count
        }
        votings_list.append(voting_info)
    return make_response({"votings": votings_list}, 200)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=4242)
