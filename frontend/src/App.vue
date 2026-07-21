<template>
    <div class="page">
        <header class="topbar">
            <h1>Cult of the Tree</h1>
            <div class="topbar-actions">
                <button class="pill-btn create-btn" @click="showCreate = true">Create post</button>
                <button v-if="isGolden" class="pill-btn mason-btn" @click="showPanel = true">Golden mason menu</button>
                <button class="pill-btn logout-btn" @click="doLogout">Logout</button>
            </div>
        </header>
        <div class="map-frame">
            <span class="map-label">map</span>
            <p v-if="!posts.length" class="empty">No sightings yet.</p>
            <button v-for="post in posts" :key="post.id" class="pin" :style="pinStyle(post)" :title="post.title"
                @click="selectedPost = post" />
        </div>
        <CreatePostWindow v-if="showCreate" @close="showCreate = false" @created="onPostCreated" />
        <GoldMasonPanel v-if="showPanel" @close="showPanel = false" />
        <PostDetailWindow v-if="selectedPost" :post="selectedPost" @close="selectedPost = null"
            @approved="onApproved" />
    </div>
</template>
<script setup>
import { computed, onMounted, ref } from "vue";
import { fetchPosts } from "./api/client";
import CreatePostWindow from "./components/CreatePostWindow.vue";
import GoldMasonPanel from "./components/GoldMasonPanel.vue";
import PostDetailWindow from "./components/PostDetailWindow.vue";
const posts = ref([]);
const showCreate = ref(false);
const showPanel = ref(false);
const selectedPost = ref(null);
const isGolden = ref(true);
function onPostCreated(post) {
    posts.value.unshift(post);
}
function onApproved(updatedPost) {
    const idx = posts.value.findIndex((p) => p.id === updatedPost.id);
    if (idx !== -1) posts.value[idx] = updatedPost;
    selectedPost.value = updatedPost;
}
function doLogout() {
    console.log("logout");
}
const bounds = computed(() => {
    if (!posts.value.length) return null;
    const lats = posts.value.map((p) => p.latitude);
    const lngs = posts.value.map((p) => p.longitude);
    return {
        minLat: Math.min(...lats),
        maxLat: Math.max(...lats),
        minLng: Math.min(...lngs),
        maxLng: Math.max(...lngs),
    };
});
function pinStyle(post) {
    const b = bounds.value;
    if (!b) return { left: "50%", top: "50%" };
    const latRange = b.maxLat - b.minLat || 1;
    const lngRange = b.maxLng - b.minLng || 1;
    const left = ((post.longitude - b.minLng) / lngRange) * 80 + 10;
    const top = (1 - (post.latitude - b.minLat) / latRange) * 80 + 10;
    return { left: `${left}%`, top: `${top}%` };
}
onMounted(async () => {
    posts.value = await fetchPosts();
});
</script>
<style>
html,
body,
#app {
    margin: 0;
    min-height: 100%;
    font-family: -apple-system, "Segoe UI", Roboto, sans-serif;
    background: #fafafa;
}
</style>
<style scoped>
.page {
    min-height: 100vh;
}

.topbar {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 14px 24px;
    background: #fafafa;
    border-bottom: 1px solid #eaeaea;
}

.topbar h1 {
    font-size: 18px;
    margin: 0;
    color: #1a1a1a;
}

.topbar-actions {
    display: flex;
    gap: 10px;
}

.pill-btn {
    border: 1px solid #d7d7d7;
    border-radius: 20px;
    padding: 8px 16px;
    font-size: 13px;
    font-weight: 600;
    cursor: pointer;
    background: #ffffff;
}

.create-btn {
    background: #4caf50;
    border-color: #4caf50;
    color: white;
}

.mason-btn {
    background: #ffd75e;
    border-color: #ffd75e;
    color: #4a3b00;
}

.logout-btn {
    background: #ffffff;
    color: #555;
}

.map-frame {
    position: relative;
    margin: 20px 24px;
    height: 60vh;
    min-height: 380px;
    background: #eef3ee;
    border: 1px solid #dfe6df;
    border-radius: 14px;
    overflow: hidden;
}

.map-label {
    position: absolute;
    top: 12px;
    left: 16px;
    font-size: 13px;
    color: #8a978a;
    font-weight: 600;
}

.empty {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    color: #999;
    font-size: 13px;
    margin: 0;
}

.pin {
    position: absolute;
    width: 14px;
    height: 14px;
    border-radius: 50%;
    background: #4caf50;
    border: 2px solid #ffffff;
    box-shadow: 0 1px 4px rgba(0, 0, 0, 0.25);
    transform: translate(-50%, -50%);
    cursor: pointer;
    padding: 0;
}

.pin:hover {
    background: #3d8b40;
}
</style>