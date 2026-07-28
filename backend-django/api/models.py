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
            "id": self.userid,
            "username": self.userdisplayname,
            "status": self.userstatus,
            "email": self.useremail,
            "is_inquizitor": self.userisinquisitor
        }
    
    class Meta:
        db_table = "users"


    
class BannedIP(models.Model):
    userid = models.BigAutoField(primary_key=True)
    bannedip = models.CharField(max_length=40)

    def to_dict(self):
        return {"ip": self.ip_address}

    
class Invite(models.Model):
    email = models.EmailField()
    created_at = models.DateTimeField(auto_now_add=True)

    def to_dict(self):
        return {
            "id": self.messageid,
            "message": self.messagecontext,
            "status": self.messagestatus,
            "created_at": self.created_at.isoformat()
        }

    class Meta:
        db_table = "messages"
