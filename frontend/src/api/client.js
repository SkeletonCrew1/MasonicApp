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

export async function sendBroadcast(message, statuses) {
    for (const status of statuses) {
        await request("/broadcast/", {
            method: "POST",
            body: JSON.stringify({ message, status }),
        });
    }
}

export function inviteUser(email) {
    return request("/invite/", {
        method: "POST",
        body: JSON.stringify({ useremail: email }),
    });
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

export function promoteUser(userDisplayName) {
    return request("/promotion/", {
        method: "POST",
        body: JSON.stringify({ userdisplayname: userDisplayName }),
    });
}

export function deleteAllData() {
    return request("/delete-all/", { method: "POST" });
}