from django.urls import path
from . import views
urlpatterns = [
    path("users/search/", views.users_search),
    path("users/<int:user_id>/promote/", views.user_promote),
    path("broadcast/", views.broadcast),
    path("invite/", views.invite),
    path("ban/", views.ban_ip),
    path("bans/", views.bans_list),
    path("delete-all/", views.delete_all),
]