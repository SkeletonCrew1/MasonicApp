from django.contrib import admin

from .models import BannedIP, Invite, Profile

admin.site.register(Profile)
admin.site.register(BannedIP)
admin.site.register(Invite)