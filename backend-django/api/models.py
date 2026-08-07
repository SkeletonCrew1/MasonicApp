from django.db import models


class User(models.Model):

    status_choices = [
    ("bronze", "Bronze"),
    ("silver", "Silver"),
    ("golden", "Golden"),
    ]  

    userid = models.BigAutoField(primary_key=True)
    userdisplayname = models.CharField(max_length=150, unique=True)
    userstatus = models.CharField(max_length=10, choices=status_choices, default="bronze")
    useremail = models.CharField(max_length=150, unique=True)

    
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
    ipid = models.BigAutoField(primary_key=True)
    bannedip = models.CharField(max_length=40, db_column="bannedip")
    
    class Meta:
        db_table = "bannedips"


class Message(models.Model):
    messageid = models.BigAutoField(primary_key=True)
    messagecontext = models.CharField(max_length=200)
    messagerecieverstatus = models.CharField(max_length=40)
    messagestatus = models.CharField(max_length=20)
   # senderid = models.BigAutoField(primary_key=True) 
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



class WhiteList(models.Model):
    invitationid = models.BigAutoField(primary_key=True)
    invitedemail = models.CharField(max_length=150, unique=True)

    class Meta:
        db_table = "whitelist"
