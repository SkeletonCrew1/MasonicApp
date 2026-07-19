<template>
    <div class="page">
        <header class="topbar">
            <h1>Cult of the Tree</h1>
            <button class="mason-btn" @click="showPanel = true">Gold mason panel</button>
        </header>
        <button class="report-btn" @click="showCreate = true">+ Report sighting</button>
        <div class="post-list">
            <p v-if="!posts.length" class="empty">No sightings yet.</p>
            <div v-for="post in posts" :key="post.id" class="post-card">
                <img v-if="post.image_url" :src="post.image_url" />
                <div class="post-body">
                    <strong>{{ post.title }}</strong>
                    <p v-if="post.description">{{ post.description }}</p>
                    <span class="coords">{{ post.latitude }}, {{ post.longitude }}</span>
                </div>
            </div>
        </div>
        <CreatePostWindow v-if="showCreate" @close="showCreate = false" @created="onPostCreated" />
        <GoldMasonPanel v-if="showPanel" @close="showPanel = false" />
    </div>
</template>
<script setup>
import { onMounted, ref } from "vue";
import { fetchPosts } from "./api/client";
import CreatePostWindow from "./components/CreatePostWindow.vue";
import GoldMasonPanel from "./components/GoldMasonPanel.vue";
const posts = ref([]);
const showCreate = ref(false);
const showPanel = ref(false);
function onPostCreated(post) {
    posts.value.unshift(post);
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

.mason-btn {
    background: #ffd75e;
    border: none;
    border-radius: 6px;
    padding: 6px 12px;
    font-size: 13px;
    font-weight: 600;
    cursor: pointer;
}

.report-btn {
    margin: 20px 24px 0;
    background: #4caf50;
    color: white;
    border: none;
    padding: 10px 16px;
    border-radius: 8px;
    font-size: 13px;
    font-weight: 600;
    cursor: pointer;
}

.post-list {
    padding: 20px 24px;
    display: flex;
    flex-direction: column;
    gap: 12px;
    max-width: 480px;
}

.empty {
    color: #999;
    font-size: 13px;
}

.post-card {
    background: white;
    border: 1px solid #eaeaea;
    border-radius: 10px;
    overflow: hidden;
}

.post-card img {
    width: 100%;
    display: block;
    max-height: 180px;
    object-fit: cover;
}

.post-body {
    padding: 10px 14px;
}

.post-body p {
    margin: 4px 0;
    font-size: 13px;
    color: #555;
}

.coords {
    font-size: 11px;
    color: #999;
}
</style>
