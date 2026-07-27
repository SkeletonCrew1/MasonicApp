from django.urls import path
from . import views


urlpatterns = [
    path("broadcast/", views.broadcast),
    path("invite/", views.invite),
    path("ban/", views.ban_ip),
    path("bans/", views.bans_list),
    path("delete-all/", views.delete_all),
    path("promotion", views.user_promotion),
]