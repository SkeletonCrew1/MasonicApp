from django.conf import settings
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
    email = models.EmailField(blank=True, default="")
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
class Post(models.Model):
    id = models.BigAutoField(primary_key=True)
    title = models.CharField(max_length=200)
    image_url = models.URLField(blank=True, default="")
    description = models.TextField(blank=True, default="")
    latitude = models.FloatField()
    longitude = models.FloatField()
    status = models.CharField(max_length=10, choices=POST_STATUS_CHOICES, default="pending")
    created_at = models.DateTimeField(auto_now_add=True)
    
    class Meta:
        ordering = ["-created_at"]
    def __str__(self):
        return self.title
    def to_dict(self):
        return {
            "id": self.id,
            "title": self.title,
            "image_url": self.image_url,
            "description": self.description,
            "latitude": self.latitude,
            "longitude": self.longitude,
            "status": self.status,
            "created_at": self.created_at.isoformat(),
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
    
class Broadcast(models.Model):
    MESSAGE_STATUS_CHOICES = [
        ("queued", "Queued"),
        ("sent", "Sent"),
        ("partial", "Partial"),
        ("failed", "Failed"),
    ]
    message = models.TextField()
    message_status = models.CharField(
        max_length=20, choices=MESSAGE_STATUS_CHOICES, default="queued"
    )
    sender = models.ForeignKey(
        settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, null=True, blank=True,
        related_name="broadcasts",
    )
    created_at = models.DateTimeField(auto_now_add=True)
    class Meta:
        ordering = ["-created_at"]
    def to_dict(self):
        return {
            "id": self.id,
            "message": self.message,
            "message_status": self.message_status,
            "sender_id": self.sender_id,
            "created_at": self.created_at.isoformat(),
        }