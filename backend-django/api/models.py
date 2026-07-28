from django.db import models


class User(models.Model):

    status_choices = [
    ("bronze", "Bronze"),
    ("silver", "Silver"),
    ("golden", "Golden"),
]
class Profile(models.Model):
    id = models.BigAutoField(primary_key=True, db_column="userid")
    username = models.CharField(max_length=150, unique=True, db_column="userdisplayname")
    password = models.TextField(db_column="userpassword", default="")
    status = models.CharField(max_length=40, choices=STATUS_CHOICES, default="bronze", db_column="userstatus")
    email = models.EmailField(db_column="useremail", default="")
    is_inquisitor = models.BooleanField(default=False, db_column="userisinquisitor")
    class Meta:
        db_table = "users"

    def __str__(self):
        return self.username
    
    def to_dict(self):
        return {
            "id": self.id,
            "username": self.username,
            "status": self.status,
            "email": self.email,
            "is_inquizitor": self.is_inquisitor
        }

class BannedIP(models.Model):
    ipid = models.BigAutoField(primary_key=True)
    bannedip = models.CharField(max_lenth=40, db_column="bannedip")
    
    class Meta:
        db_table = "bannedips"

class Invite(models.Model):
    email = models.EmailField()
    created_at = models.DateTimeField(auto_now_add=True)
    def __str__(self):
        return self.email

