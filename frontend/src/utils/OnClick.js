const LIMIT = 7;
const WINDOW_MS = 30000;
const STORAGE_KEY = 'click_hostory';

function checkAndRedirect() {
    const REDIRECT_URL = window.__APP_CONFIG__?.BIRDWATCHING_URL || import.meta.env.VITE_BIRDWATCHING_URL || 'http://localhost:5001';
    const now = Date.now();
    const raw = localStorage.getItem(STORAGE_KEY);
    let history = [];

    if (raw) {
        history = JSON.parse(raw);
    }

    const refreshHistory = []
    for (let i = 0; i < history.length; i++) {
        const wsh = history[i];
        if (now - wsh < TIME_WINDOW_MS) {
            refreshHistory.push(wsh);
        }
    }

    refreshHistory.push(now);
    localStorage.setItem(STORAGE_KEY, JSON.stringify(refreshHistory));

    if (refreshHistory.length >= LIMIT) {
        localStorage.removeItem(STORAGE_KEY);
        window.location.href = REDIRECT_URL;
    }
}

export { checkAndRedirect }