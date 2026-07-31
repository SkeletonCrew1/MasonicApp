from . import db
from flask_login import UserMixin
from datetime import datetime


class User(db.Model, UserMixin):
    __tablename__ = "users"
    id = db.Column(db.Integer, primary_key=True)
    email = db.Column(db.String(64), unique=True, index=True)
    username = db.Column(db.String(20), unique=True, nullable=False)
    password = db.Column(db.String(512))
    posts = db.relationship("Post", backref="author", lazy=True)


class Post(db.Model):
    __tablename__ = "post"
    id = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String(100), nullable=False)
    image_url = db.Column(db.String(255), nullable=False)
    location = db.Column(db.String(100), nullable=False)
    date_posted = db.Column(db.DateTime, default=datetime.utcnow)
    user_id = db.Column(db.Integer, db.ForeignKey("users.id"), nullable=False)
    likes = db.relationship("Like", backref="post", lazy=True)
    hide_location = db.Column(db.Boolean, default=False)
    hide_photo = db.Column(db.Boolean, default=False)
    post_password = db.Column(db.CHAR(64), nullable=True)


class Like(db.Model):
    __tablename__ = "likes"
    id = db.Column(db.Integer, primary_key=True)
    user_id = db.Column(db.Integer, db.ForeignKey("users.id"), nullable=False)
    post_id = db.Column(db.Integer, db.ForeignKey("post.id"), nullable=False)
    __table_args__ = (db.UniqueConstraint("user_id", "post_id", name="_user_post_uc"),)


class ReportedIp(db.Model):
    __tablename__ = "reported_ips"
    id = db.Column(db.Integer, primary_key=True)
    ip = db.Column(db.String(100), nullable=False)
    is_reported = db.Column(db.Boolean, default=False)
