from django.contrib import admin

from .models import BannedIP, User


admin.site.register(User)
admin.site.register(BannedIP)
