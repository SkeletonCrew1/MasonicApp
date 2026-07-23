from django.db import models


STATUS_CHOICES = [
    ("bronze", "Bronze"),
    ("silver", "Silver"),
    ("golden", "Golden"),
]

POST_STATUS_CHOICES = [
    ("pending", "Pending"),
    ("approved", "Approved"),
]


class Profile(models.Model):
    id = models.BigAutoField(primary_key=True)
    username = models.CharField(max_length=150, unique=True)
    status = models.CharField(max_length=10, choices=STATUS_CHOICES, default="bronze")
    is_banned = models.BooleanField(default=False)
    created_at = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return self.username
    
    def to_dict(self):
        return {
            "id": self.id,
            "username": self.username,
            "status": self.status,
            "is_banned": self.is_banned,
        }

    
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


