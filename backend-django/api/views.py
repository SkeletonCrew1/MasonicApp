import json
import os
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from django.views.decorators.http import require_http_methods
from .models import BannedIP, User

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

    subject = "Broadcast"
    data = _body(request)
    message = data.get("message")
    status = data.get("status")
    if not message:
        return JsonResponse({"error": "Please define a message"}, status=400)
    if not status:
        return JsonResponse({"error": "Status is required"}, status=400)

    emails = list(User.objects.values_list('useremail', flat=True))

    data_to_send = {
        "dest": emails,
        "subject": subject,
        "body": message
    }

    response = requests.post(settings.MAIL_SERVICE_URL, json = data_to_send)
    return JsonResponse({ "data": data_to_send, "response": response.status_code }, status=200)
     



# ---------- Ban ----------
@csrf_exempt
@require_http_methods(["POST"])
def ban_ip(request):
    data = json.loads(request.body)
    ip = (data.get("ip"))
    
    if not ip:
        return JsonResponse({"error": "IP is required"}, status=400)
    
    try:
        BannedIP.objects.get(bannedip=ip)
        return JsonResponse({"message": f"{ip} already banned"}, status=200)
    except BannedIP.DoesNotExist:
        ban = BannedIP(bannedip=ip)
        ban.save()
        return JsonResponse({"message": f"{ip} banned successfully"}, status=200)

@require_http_methods(["GET"])
def bans_list(request):
    bans = list(BannedIP.objects.values("ipid", "bannedip"))
    return JsonResponse(bans, safe=False)


# ---------- Delete all ----------
@csrf_exempt
@require_http_methods(["POST"])
def delete_all(request):
    User.objects.all().delete()
    return JsonResponse({"deleted": True})
@csrf_exempt


# promotion feature                                                                      
@csrf_exempt
@require_http_methods(["POST"])
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