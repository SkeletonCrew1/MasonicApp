from flask import Blueprint, render_template, redirect, url_for, flash, request
from flask_login import login_user, logout_user, login_required, current_user
from werkzeug.security import generate_password_hash, check_password_hash
from werkzeug.utils import secure_filename
from . import db, s3, BUCKET_NAME, login_manager
from .models import User, Post, Like, ReportedIp
from forms import RegistrationForm, LoginForm, PostForm, UnlockForm, ReportIpForm
import hashlib

app = Blueprint("app", __name__)


@login_manager.user_loader
def load_user(user_id):
    return User.query.get(int(user_id))


def get_presigned_url(filename):
    """Generates a temporary URL for private S3 objects."""
    try:
        url = s3.generate_presigned_url(
            "get_object",
            Params={"Bucket": BUCKET_NAME, "Key": filename},
            ExpiresIn=3600,
        )
        return url
    except Exception:
        return None


@app.route("/")
def index():
    posts = Post.query.all()
    for post in posts:
        filename = post.image_url.split("/")[-1]
        post.presigned_url = get_presigned_url(filename)
    return render_template("index.html", posts=posts)


@app.route("/register", methods=["GET", "POST"])
def register():
    form = RegistrationForm()
    if form.validate_on_submit():
        hashed_pw = generate_password_hash(form.password.data)
        user = User.query.filter_by(email=form.email.data).first()
        if user:
            flash("This email already registered", "danger")

        else:
            db.session.add(
                User(
                    email=form.email.data,
                    username=form.username.data,
                    password=hashed_pw,
                )
            )
            db.session.commit()
            return redirect(url_for("app.login"))
    return render_template("register.html", form=form)


@app.route("/login", methods=["GET", "POST"])
def login():
    form = LoginForm()
    if form.validate_on_submit():
        user = User.query.filter_by(email=form.email.data).first()
        if user and check_password_hash(user.password, form.password.data):
            login_user(user)
            return redirect(url_for("app.index"))
        else:
            flash("Incorrect password", "danger")

    return render_template("login.html", form=form)


@app.route("/logout")
def logout():
    logout_user()
    return redirect(url_for("app.index"))


@app.route("/create_post", methods=["GET", "POST"])
@login_required
def create_post():
    form = PostForm()
    if form.validate_on_submit():
        hashed_post_pw = None

        if form.hide_photo.data and form.post_password.data:
            hashed_post_pw = generate_password_hash(form.post_password.data)

        file = form.picture.data
        image_url = None
        filename = "id" + str(len(Post.query.all())) + secure_filename(file.filename)
        s3.upload_fileobj(
            file, BUCKET_NAME, filename, ExtraArgs={"ContentType": file.content_type}
        )
        image_url = f"https://{BUCKET_NAME}.s3.amazonaws.com/{filename}"

        new_post = Post(
            title=form.title.data,
            location=form.location.data,
            image_url=image_url,
            user_id=current_user.id,
            hide_location=form.hide_location.data,
            hide_photo=form.hide_photo.data,
            post_password=hashed_post_pw,
        )
        db.session.add(new_post)
        db.session.commit()
        return redirect(url_for("app.index"))
    return render_template("create_post.html", form=form)


@app.route("/post/<int:post_id>", methods=["GET", "POST"])
def post_detail(post_id):
    post = Post.query.get_or_404(post_id)
    filename = post.image_url.split("/")[-1]
    secure_url = get_presigned_url(filename)

    if post.post_password:
        form = UnlockForm()
        if form.validate_on_submit():
            if check_password_hash(post.post_password, form.password.data):
                return render_template(
                    "post.html", post=post, image_url=secure_url, unlocked=True
                )
            else:
                flash("Incorrect password", "danger")

        return render_template("unlock.html", form=form, post=post)
    return render_template("post.html", post=post, image_url=secure_url, unlocked=True)


@app.route("/like/<int:post_id>")
@login_required
def like_post(post_id):
    existing_like = Like.query.filter_by(
        user_id=current_user.id, post_id=post_id
    ).first()
    if not existing_like:
        db.session.add(Like(user_id=current_user.id, post_id=post_id))
    else:
        db.session.delete(existing_like)
    db.session.commit()

    return redirect(url_for("app.post_detail", post_id=post_id))


@app.route("/hunter/report", methods=["GET", "POST"])
@login_required
def hunter_report():
    form = ReportIpForm()
    if form.validate_on_submit():
        target_ip = form.ip_address.data
        protected_ip = ("0.0.0.0", "127.0.0.1")
        if target_ip in protected_ip:
            flash("Incorrect ip", "danger")
            return render_template("report_user.html", form=form)

        hashed_ip = hashlib.sha256(target_ip.encode("utf-8")).hexdigest()

        existing_record = ReportedIp.query.filter_by(ip=hashed_ip).first()
        if existing_record:
            flash("ip is already added", "danger")
            existing_record.is_reported = True
            return render_template("report_user.html", form=form)

        else:
            new_ban = ReportedIp(ip=hashed_ip, is_reported=True)
            db.session.add(new_ban)
        db.session.commit()
        return redirect(url_for("app.index"))

    return render_template("report_user.html", form=form)


@app.before_request
def block_reported_ips():
    db.create_all()
    visitor_ip = request.remote_addr
    ip_record = ReportedIp.query.filter_by(ip=visitor_ip).first()

    if ip_record and ip_record.is_reported:
        return redirect("https://zakon.rada.gov.ua/laws/show/3325-17")
