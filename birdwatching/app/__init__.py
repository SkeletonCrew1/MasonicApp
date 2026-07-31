from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from flask_login import LoginManager
import boto3
import os
from dotenv import load_dotenv

db = SQLAlchemy()
login_manager = LoginManager()


def create_app():
    global s3, BUCKET_NAME

    app = Flask(__name__)
    load_dotenv()

    app.config["SECRET_KEY"] = os.environ.get("SECRET_KEY")
    app.config["SQLALCHEMY_DATABASE_URI"] = os.environ.get("SQLALCHEMY_DATABASE_URI", "sqlite:////app/bird.db")
    app.config["SQLALCHEMY_TRACK_MODIFICATIONS"] = False
    db.init_app(app)
    login_manager.init_app(app)
    login_manager.login_view = "app.login"

    s3 = boto3.client(
        "s3",
        aws_access_key_id=os.environ.get("AWS_ACCESS_KEY"),
        aws_secret_access_key=os.environ.get("AWS_SECRET_KEY"),
    )
    BUCKET_NAME = os.environ.get("S3_BUCKET")

    from .routes import app as app_blueprint

    app.register_blueprint(app_blueprint)

    return app
