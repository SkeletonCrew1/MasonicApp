from django.db import models


class User(models.Model):

    status_choices = [
    ("bronze", "Bronze"),
    ("silver", "Silver"),
    ("gold", "Gold"),
    ]  

    userid = models.BigAutoField(primary_key=True)
    userdisplayname = models.CharField(max_length=150, unique=True)
    userstatus = models.CharField(max_length=10, choices=status_choices, default="bronze")

    def __str__(self):
        return self.username
    
    def to_dict(self):
        return {
            "id": self.id,
            "username": self.username,
            "status": self.status,
            "is_banned": self.is_banned,
        }
    
    class Meta:
        db_table = "users"


    
class BannedIP(models.Model):
    ip_address = models.GenericIPAddressField(unique=True)
    created_at = models.DateTimeField(auto_now_add=True)
    def __str__(self):
        return self.ip_address
    def to_dict(self):
        return {"ip": self.ip_address}

    
class Invite(models.Model):
    email = models.EmailField()
    created_at = models.DateTimeField(auto_now_add=True)
    def __str__(self):
        return self.email
