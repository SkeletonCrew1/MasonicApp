from django.urls import path
from . import views
urlpatterns = [
    path("posts/", views.posts_list),
    path("posts/<int:post_id>/approve/", views.post_approve),
    path("users/search/", views.users_search),
    path("users/<int:user_id>/promote/", views.user_promote),
    path("broadcast/", views.broadcast),
    path("broadcasts/", views.broadcasr_history),
    path("invite/", views.invite),
    path("ban/", views.ban_ip),
    path("bans/", views.bans_list),
    path("delete-all/", views.delete_all),
]