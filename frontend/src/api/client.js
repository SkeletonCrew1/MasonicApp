
async function request(path, options = {}) {
    const res = await fetch(`/backend${path}`, {
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




async function votingRequest(path, options = {}) {
    const headers = { "Content-Type": "application/json" };

    const res = await fetch(`/voting${path}`, {
        headers: { ...headers, ...(options.headers || {}) },
        credentials: "include",
        ...options,
    });
    
    if (!res.ok) {
        const body = await res.json().catch(() => ({}));
        throw new Error(body.error || body.message || "Voting request failed");
    }
    return res.json();
}

export function getVotings(userId, status) {
    return votingRequest("/get_votings", {
        method: "POST",
        body: JSON.stringify({ user_id: userId, status })
    }); 
}

export function createVoting(votingSubject, votingCategory) {
    return votingRequest("/create_voting", {
        method: "POST",
        body: JSON.stringify({ voting_subject: votingSubject, voting_category: votingCategory })
    });
}

export function addVote(votingId, voterId) {
    return votingRequest("/vote", {
        method: "POST",
        body: JSON.stringify({ voting_id: votingId, voter_id: voterId })
    });
}