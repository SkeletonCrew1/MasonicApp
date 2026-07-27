import json
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from django.views.decorators.http import require_http_methods
from .models import BannedIP, Invite, User

def _body(request):
    if not request.body:
        return {}
    try:
        return json.loads(request.body)
    except json.JSONDecodeError:
        return {}


# ---------- Broadcast ----------
@csrf_exempt
@require_http_methods(["POST"])
def broadcast(request):
    data = _body(request)
    message = (data.get("message") or "").strip()
    statuses = data.get("statuses") or []
    if not message :
        return JsonResponse({"error": "Message is required"}, status=400)
    recipients_count = User.objects.filter(status__in=statuses).count()
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
    User.objects.all().delete()
    return JsonResponse({"deleted": True})
@csrf_exempt


# promotion feature                                                                      
@csrf_exempt
def user_promotion(request):

    data = json.loads(request.body)
    uid = data.get("id")

    try:
        user = User.objects.get(userid=uid)
    except User.DoesNotExist:
        return JsonResponse({"error": "User not found"}, status=404)

    if user.userstatus == "bronze":
        user.userstatus = "silver"
    elif user.userstatus == "silver":
        user.userstatus = "gold"
    else:
        return JsonResponse({"message": f"{user.userdisplayname} already with 'gold' status"}, status=200)

    user.save(update_fields=["userstatus"])
    return JsonResponse({"message": f"{user.userdisplayname} promoted to {user.userstatus} status"}, status=200)