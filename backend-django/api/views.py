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
    

# ---------- Broadcast ----------
@csrf_exempt
@require_http_methods(["POST"])
def broadcast(request):
    data = _body(request)
    message = (data.get("message") or "").strip() # do we need .strip() ?
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
    Profile.objects.all().delete()
    return JsonResponse({"deleted": True})


# promotion feature 
@csrf_exempt
def user_promotion(request):

  #  data = json.loads(request.body)
   # uid = data.get("id")
    data = _body(request)
    uid = data.get("id")

    try:
        user = Profile.objects.get(uid)
    except Profile.DoesNotExist:
        return JsonResponse({"error": "User not found"}, status=404)

    if user.status == "bronze":
        user.status = "silver"
    elif user.status == "silver":
        user.status = "gold"
    else:
        return JsonResponse({"message": "User already with gold status"}, status=200)
    
    user.save()
    return JsonResponse({"message": f"User {user.username} promoted to {user.status}"}, status=200)