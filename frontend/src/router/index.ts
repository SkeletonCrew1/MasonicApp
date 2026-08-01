import { createRouter, createWebHistory } from "vue-router";

import Home from "../pages/Home.vue";
import AddSighting from "../pages/AddSighting.vue";
import SightingDetail from '../pages/SightingDetail.vue';
import GoldMasonPanel from "../pages/GoldMasonPanel.vue"
import Register from "../pages/Registration.vue";
import Login from "../pages/Login.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/home",
            component: Home
        },
        {
            path: "/add-sighting",
            component: AddSighting
        },
        {
            path: "/sighting",
            component: SightingDetail
        },
        {
            path: '/sighting/:id',
            component: SightingDetail
        },
        {
            path: "/control-panel",
            component: GoldMasonPanel
        },
        {
            path: "/login",
            component: Login
        },
        {
            path: "/",
            component: Register
        }
    ]
});

export default router;