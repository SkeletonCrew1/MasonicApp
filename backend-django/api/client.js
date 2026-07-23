export function sendBroadcast(message, statuses) {
    return request("/broadcast/", {
        method: "POST",
        body: JSON.stringify({ message, statuses }),
    });
}
export function fetchBroadcastHistory() {
    return request("/broadcasts/");
}