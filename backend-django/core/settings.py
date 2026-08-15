import os
from pathlib import Path


FRONTEND_PAGE = os.getenv("FRONTEND_SERVICE_URL")
MAIL_SERVICE_URL = os.getenv("MAIL_SERVICE_URL")


BASE_DIR = Path(__file__).resolve().parent.parent

DEBUG = True
ALLOWED_HOSTS = ["*"]
INSTALLED_APPS = [
    "django.contrib.admin",
    "django.contrib.auth",
    "django.contrib.contenttypes",
    "django.contrib.sessions",
    "django.contrib.messages",
    "django.contrib.staticfiles",
    "corsheaders",
    "api",
]
MIDDLEWARE = [
    "corsheaders.middleware.CorsMiddleware",
    "django.middleware.security.SecurityMiddleware",
    "django.contrib.sessions.middleware.SessionMiddleware",
    "django.middleware.common.CommonMiddleware",
    "django.middleware.csrf.CsrfViewMiddleware",
    "django.contrib.auth.middleware.AuthenticationMiddleware",
    "django.contrib.messages.middleware.MessageMiddleware",
    "django.middleware.clickjacking.XFrameOptionsMiddleware",
]
ROOT_URLCONF = "core.urls"
TEMPLATES = [
    {
        "BACKEND": "django.template.backends.django.DjangoTemplates",
        "DIRS": [],
        "APP_DIRS": True,
        "OPTIONS": {
            "context_processors": [
                "django.template.context_processors.debug",
                "django.template.context_processors.request",
                "django.contrib.auth.context_processors.auth",
                "django.contrib.messages.context_processors.messages",
            ],
        },
    },
]
WSGI_APPLICATION = "core.wsgi.application"

DATABASES = {
    "default": {
        "ENGINE": "django.db.backends.postgresql",
        "NAME": os.environ.get("USERS_DB_NAME"),
        "USER": os.environ.get("USERS_DB_USER"),
        "PASSWORD": os.environ.get("USERS_DB_PASSWORD"),
        "HOST": os.environ.get("USERS_DB_HOST"),
        "PORT": os.environ.get("USERS_DB_PORT"),
    },
    "map-db": {
        "ENGINE": "django.db.backends.postgresql",
        "NAME": os.environ.get("MAP_DB_NAME"),
        "USER": os.environ.get("MAP_DB_USER"),
        "PASSWORD": os.environ.get("MAP_DB_PASSWORD"),
        "HOST": os.environ.get("MAP_DB_HOST"),
        "PORT": os.environ.get("MAP_DB_PORT"),
    },
    "voting-db": {
        "ENGINE": "django.db.backends.postgresql",
        "NAME": os.environ.get("VOTING_DB_NAME"),
        "USER": os.environ.get("POSTGRES_USER"),
        "PASSWORD": os.environ.get("POSTGRES_PASSWORD"),
        "HOST": os.environ.get("VOTING_DB_HOST"),
        "PORT": os.environ.get("DB_PORT"),
    }
}

AUTH_PASSWORD_VALIDATORS = []
LANGUAGE_CODE = "en-us"
TIME_ZONE = "UTC"
USE_I18N = True
USE_TZ = True
STATIC_URL = "static/"
DEFAULT_AUTO_FIELD = "django.db.models.BigAutoField"
CORS_ALLOW_ALL_ORIGINS = True
