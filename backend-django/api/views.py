import json
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from django.views.decorators.http import require_http_methods
from .models import BannedIP, Invite, Profile

def _body(request):
    if not request.body:
        return {}
    try:
        return json.loads(request.body)
    except json.JSONDecodeError:
        return {}

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
@require_http_methods(["POST"])
def broadcast(request):
    data = _body(request)
    message = (data.get("message") or "").strip()
    statuses = data.get("statuses") or []
    if not message:
        return JsonResponse({"error": "Message is required"}, status=400)
    recipients_count = Profile.objects.filter(status__in=statuses).count()
    return JsonResponse({"sent": True, "recipients": recipients_count})
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