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
    
def get_client_ip(request):
    x_forwarded_for = request.META.get("HTTP_X_FORWARDED_FOR")
    if x_forwarded_for:
        return x_forwarded_for.split(",")[0].strip()
    return request.META.get("REMOTE_ADDR", "")

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
        return JsonResponse({"error": "Please select at least one status"}, status=400)
    if not statuses:
            return JsonResponse({"error": "Message is required"}, status=400)
    recipients_count = Profile.objects.filter(status__in=statuses).count()
    if recipients_count == 0:
        return JsonResponse({"error": "No users found for selected statuses"}, status=404)
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
    user_id = data.get("user_id")
    ip = (data.get("ip") or get_client_ip(request)).strip()
    if not user_id:
            return JsonResponse({"error": "user_id is required"}, status=400)
    if not ip:
        return JsonResponse({"error": "IP is required"}, status=400)
    try:
        profile = Profile.objects.get(id=user_id)
    except Profile.DoesNotExist:
        return JsonResponse({"error": "User not found"}, status=404)
    banned_ip_obj, created = BannedIP.objects.update_or_create(
        user=profile,
        defaults={"banned_ip": ip}
    )
    
    return JsonResponse({
        "user_id": profile.id,
        "ip": banned_ip_obj.banned_ip,
        "banned": True,
        "is_new_ban": created
    })

@require_http_methods(["GET"])
def bans_list(request):
    bans = BannedIP.objects.select_related('user').all()
    return JsonResponse([b.to_dict() for b in bans], safe=False)
# ---------- Delete all ----------
@csrf_exempt
@require_http_methods(["POST"])
def delete_all(request):
    BannedIP.objects.all().delete()
    return JsonResponse({"deleted": True})