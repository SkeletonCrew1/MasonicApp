from flask import Flask, request, make_response
from config import MAIN_DATABASE_URL, VOTING_DATABASE_URL
from models import db, Voting, Vote

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

    new_voting = Voting(voting_subject=voting_subject, voting_category=voting_category)
    db.session.add(new_voting)
    db.session.commit()
    return make_response({"success": "New voting was created"}, 200)


@app.route("/vote", methods=['POST'])
def add_vote():
    data = request.get_json()
    voting_id = data.get("voting_id")
    voter_id = data.get("voter_id")

    added_vote = Vote(voting_id=voting_id, voter_id=voter_id)
    db.session.add(added_vote)
    db.session.commit()
    return make_response({"success": "Your vote was added"}, 200)


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=4242)
