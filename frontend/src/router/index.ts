import { createRouter, createWebHistory } from "vue-router";

import Home from "../pages/Home.vue";
import AddSighting from "../pages/AddSighting.vue";
import SightingDetail from '../pages/SightingDetail.vue';

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/",
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
        }
    ]
});

export default router;