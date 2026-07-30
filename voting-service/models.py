from flask_sqlalchemy import SQLAlchemy


db = SQLAlchemy()


class User(db.Model):
    __tablename__ = 'users'

    user_id = db.Column('userid', db.Integer, primary_key=True, autoincrement=True)
    user_display_name = db.Column('userdisplayname', db.String(40), nullable=False)
    user_password = db.Column('userpassword', db.Text, nullable=False)
    user_status = db.Column('userstatus', db.String(40), nullable=False)
    user_email = db.Column('useremail', db.Text, nullable=False)
    user_is_inquisitor = db.Column('userisinquisitor', db.Boolean, nullable=False)


class Voting(db.Model):
    __bind_key__ = 'voting-db'
    __tablename__ = 'votings'

    voting_id = db.Column('votingid', db.Integer, primary_key=True, autoincrement=True)
    voting_subject = db.Column('votingsubject', db.String(40), nullable=False)
    voting_category = db.Column('votingcategory', db.String(40), nullable=False)
    votes = db.relationship('Vote', backref='voting', cascade='all, delete-orphan', lazy=True)


class Vote(db.Model):
    __bind_key__ = 'voting-db'
    __tablename__ = 'votes'

    voting_id = db.Column(
        'votingid',
        db.Integer,
        db.ForeignKey('votings.votingid', ondelete='CASCADE'),
        primary_key=True
    )
    voter_id = db.Column('voterid', db.Integer, primary_key=True)