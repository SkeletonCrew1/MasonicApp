from django.db import models
from django.conf import settings


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

    
class BannedIP(models.Model):
    ip_address = models.GenericIPAddressField(unique=True)
    created_at = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return self.ip_address
    
    def to_dict(self):
        return {"ip": self.ip_address}
<<<<<<< HEAD


=======
    
>>>>>>> origin/fix-2-20-Implement-broadcast-backend
class Invite(models.Model):
    email = models.EmailField()
    created_at = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return self.email
<<<<<<< HEAD


=======
    
class Broadcast(models.Model):
    message = models.TextField()
    statuses = models.JSONField(default=list)
    recipients_count = models.PositiveIntegerField(default=0)
    delivered_count = models.PositiveIntegerField(default=0)
    failed_count = models.JSONField(default=list)
    sent_by = models.ForeignKey(
        settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, null=True, blank=True
    )
    created_at = models.DateTimeField(auto_now_add=True)
    class Meta:
        ordering = ["-created_at"]
        
    def to_dict(self):
        return {
            "id": self.id,
            "message": self.message,
            "recipient_count": self.recipients_count,
            "delivered_count": self.delivered_count,
            "failed_emails": self.failed_emails,
            "created_at": self.created_at.isoformat(),
        }
    
>>>>>>> origin/fix-2-20-Implement-broadcast-backend
