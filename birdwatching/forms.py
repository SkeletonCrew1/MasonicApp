from flask_wtf import FlaskForm
from flask_wtf.file import FileField, FileAllowed, FileRequired
from wtforms import BooleanField, StringField, PasswordField, SubmitField
from wtforms.validators import DataRequired, Email, Length, Optional


class RegistrationForm(FlaskForm):
    email = StringField("Email", validators=[DataRequired(), Email()])
    username = StringField("Username:", validators=[DataRequired()])
    password = PasswordField("Password", validators=[DataRequired(), Length(min=6)])
    submit = SubmitField("Register")


class LoginForm(FlaskForm):
    email = StringField("Email", validators=[DataRequired(), Email()])
    password = PasswordField("Password", validators=[DataRequired()])
    submit = SubmitField("Login")


class PostForm(FlaskForm):
    title = StringField("Bird Name", validators=[DataRequired()])
    location = StringField("Location", validators=[DataRequired()])
    picture = FileField(
        "Upload Picture",
        validators=[
            FileRequired(message="Please upload a picture! (jpg or png)"),
            FileAllowed(["jpg", "png", "jpeg"]),
        ],
    )
    hide_location = BooleanField("Hide my location")
    hide_photo = BooleanField("Hide my photo")
    post_password = PasswordField("Post Password (Optional)", validators=[Optional()])
    submit = SubmitField("Post")


class UnlockForm(FlaskForm):
    password = PasswordField("Password", validators=[DataRequired()])
    submit = SubmitField("Unlock")


class ReportIpForm(FlaskForm):
    ip_address = StringField("Target IP Address", validators=[DataRequired()])
    submit = SubmitField("Ban IP")
