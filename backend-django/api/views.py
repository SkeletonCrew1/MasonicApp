import json
from django.core.mail import get_connection, send_mail
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from django.views.decorators.http import require_http_methods
from .models import BannedIP, Invite, Post, Profile
from .permissions import golden_required

def _body(request):
    if not request.body:
        return {}
    try:
        return json.loads(request.body)
    except json.JSONDecodeError:
        return {}
# ---------- Posts ----------
@csrf_exempt
@require_http_methods(["GET", "POST"])
def posts_list(request):
    if request.method == "GET":
        posts = Post.objects.all()
        return JsonResponse([p.to_dict() for p in posts], safe=False)
    data = _body(request)
    title = (data.get("title") or "").strip()
    if not title:
        return JsonResponse({"error": "Title is required"}, status=400)
    if data.get("latitude") is None or data.get("longitude") is None:
        return JsonResponse({"error": "Latitude and longitude are required"}, status=400)
    post = Post.objects.create(
        title=title,
        image_url=data.get("image_url") or "",
        description=data.get("description") or "",
        latitude=data["latitude"],
        longitude=data["longitude"],
    )
    return JsonResponse(post.to_dict(), status=201)
@csrf_exempt
@require_http_methods(["POST", "PATCH"])
def post_approve(request, post_id):
    try:
        post = Post.objects.get(id=post_id)
    except Post.DoesNotExist:
        return JsonResponse({"error": "Post not found"}, status=404)
    post.status = "approved"
    post.save(update_fields=["status"])
    return JsonResponse(post.to_dict())
# ---------- Users / Profiles ----------
@require_http_methods(["GET"])
def users_search(request):
    query = request.GET.get("q", "").strip()
    if not query:
        return JsonResponse([], safe=False)
    profiles = Profile.objects.filter(username__icontains=query)[:20]
    return JsonResponse([p.to_dict() for p in profiles], safe=False)

@csrf_exempt
@require_http_methods(["POST"])
def user_promote(request, user_id):
    data = _body(request)
    new_status = data.get("status")
    if new_status not in ("bronze", "silver", "golden"):
        return JsonResponse({"error": "Invalid status"}, status=400)
    try:
        profile = Profile.objects.get(id=user_id)
    except Profile.DoesNotExist:
        return JsonResponse({"error": "User not found"}, status=404)
    profile.status = new_status
    profile.save(update_fields=["status"])
    return JsonResponse(profile.to_dict())
# ---------- Broadcast ----------
@csrf_exempt
@golden_required
@require_http_methods(["POST"])
def broadcast(request):
    data = _body(request)
    message = (data.get("message") or "").strip()
    statuses = data.get("statuses") or []
    if not message:
        return JsonResponse({"error": "Message is required"}, status=400)
    if not statuses:
        return JsonResponse({"error": "At least one status must be selected"}, status=400)

    recipients = (
        Profile.objects.filter(status__in=statuses)
        .exclude(email="")
        .values_list("email", flat=True)
    )
    emails = list(recipients)
    
    delivered = 0
    failed = []
    
    if emails:
        connection = get_connection()
        connection.open()
        try:
            for email in emails:
                try:
                    send_mail(
                        subject="Cult of the Tree -- Broadcast",
                        message=message,
                        from_email=None,
                        recipient_list=[email],
                        fail_silently=False,
                        connection=connection
                    )
                    delivered += 1
                except Exception:
                    failed.append(email)
        finally:
            connection.close()
    record = Broadcast.objects.create(
        message=message,
        statuses=statuses,
        recipients_count=len(emails),
        delivered_count=delivered,
        failed_emails=failed,
        sent_by=request.user if request.user.is_authenticated else None,
    )
    return JsonResponse(record.to_dict(), status=201)

@golden_required
@require_http_methods(["GET"])
def broadcasr_history(request):
    records = Broadcast.objects.all()[:50]
    return JsonResponse([r.to_dict() for r in records], safe=False)
# ---------- Invite ----------
@csrf_exempt
@require_http_methods(["POST"])
def invite(request):
    data = _body(request)
    email = (data.get("email") or "").strip()
    if not email:
        return JsonResponse({"error": "Email is required"}, status=400)
    Invite.objects.create(email=email)
    return JsonResponse({"sent": True, "email": email})
# ---------- Ban ----------
@csrf_exempt
@require_http_methods(["POST"])
def ban_ip(request):
    data = _body(request)
    ip = (data.get("ip") or "").strip()
    if not ip:
        return JsonResponse({"error": "IP is required"}, status=400)
    BannedIP.objects.get_or_create(ip_address=ip)
    return JsonResponse({"ip": ip, "banned": True})

@require_http_methods(["GET"])
def bans_list(request):
    bans = BannedIP.objects.all()
    return JsonResponse([b.to_dict() for b in bans], safe=False)
# ---------- Delete all ----------
@csrf_exempt
@require_http_methods(["POST"])
def delete_all(request):
    Post.objects.all().delete()
    return JsonResponse({"deleted": True})