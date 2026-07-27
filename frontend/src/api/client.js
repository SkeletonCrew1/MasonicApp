const API_BASE = "http://localhost:8000/api";
async function request(path, options = {}) {
    const res = await fetch(`${API_BASE}${path}`, {
        headers: { "Content-Type": "application/json" },
        ...options,
    });
    if (!res.ok) {
        const body = await res.json().catch(() => ({}));
        throw new Error(body.error || "Request failed");
    }
    return res.json();
}
export function searchUsers(query) {
    return request(`/users/search/?q=${encodeURIComponent(query)}`);
}
export function promoteUser(userId, status) {
    return request(`/users/${userId}/promote/`, {
        method: "POST",
        body: JSON.stringify({ status }),
    });
}
export function sendBroadcast(message, statuses) {
    return request("/broadcast/", {
        method: "POST",
        body: JSON.stringify({ message, statuses }),
    });
}
export function inviteUser(email) {
    return request("/invite/", { method: "POST", body: JSON.stringify({ email }) });
}

/**
 * @param { number } userId
 * @param { string } [ip]
*/
export function banIp(userId, ip = null) {
    if (!userId) throw new Error("userId is required");
    const payload = { user_id: userId };
    if (ip) {
        payload.ip = ip;
    }
    return request("/ban/", {
        method: "POST",
        body: JSON.stringify(payload)
    });
}
export function listBans() {
    return request("/bans/");
}
export function deleteAllData() {
    return request("/delete-all/", { method: "POST" });
}

export function getBroadcastHistory(status) {
    return request((`/broadcast/history/?status=${encodeURIComponent(status)}(`);
}