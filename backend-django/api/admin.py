from django.contrib import admin

from .models import BannedIP, Invite, Post, Profile

admin.site.register(Profile)
admin.site.register(Post)
admin.site.register(BannedIP)
admin.site.register(Invite)