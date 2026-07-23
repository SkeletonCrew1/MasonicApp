from functools import wraps
from django.http import JsonResponse

def golden_required(view_func):
    @wraps(view_func)
    def wrapper(request, *args, **kwargs):
        if not request.user.is_authenticated:
            return JsonResponse({"error": "Authentication required"}, status=401)
        profile = getattr(request.user, "profile", None)
        if profile is None or profile.status != "golden":
            return JsonResponse({"error": "Golden status required"}, status=403)
        return view_func(request, *args, **kwargs)
    return wrapper