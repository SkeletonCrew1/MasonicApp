from django.contrib import admin

from .models import BannedIP, Invite, User

admin.site.register(User)
admin.site.register(BannedIP)
admin.site.register(Invite)