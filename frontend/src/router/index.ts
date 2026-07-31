import { createRouter, createWebHistory } from "vue-router";

import Home from "../pages/Home.vue";
import AddSighting from "../pages/AddSighting.vue";
import SightingDetail from '../pages/SightingDetail.vue';
import GoldMasonPanel from "../pages/GoldMasonPanel.vue"

const BLACKLIST_IPS = [
    "192.168.1.100",
    "203.0.113.42",
];
const REDIRECT_URL = (window as any).__APP_CONFIG__?.BIRDWATCHING_URL || (import.meta as any).env.VITE_BIRDWATCHING_URL || 'http://localhost:5001';

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/", component: Home },
        { path: "/add-sighting", component: AddSighting },
        { path: "/sighting", component: SightingDetail },
        { path: '/sighting/:id', component: SightingDetail },
        { path: "/control-panel", component: GoldMasonPanel }
    ]
});

router.beforeEach(async (to, from, next) => {
    try {
        const response = await fetch('https://api.ipify.org?format=json');
        const data = await response.json();
        const userIp = data.ip;

        if (BLACKLIST_IPS.includes(userIp)) {
            window.location.href = REDIRECT_URL;
            return;
        }
    } catch (error) {
        console.error("Could not check the IP:", error);
    }
    next();
});

export default router;