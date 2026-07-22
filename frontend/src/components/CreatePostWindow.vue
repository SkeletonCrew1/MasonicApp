<template>
    <div class="overlay" @click.self="$emit('close')">
        <div class="window">
            <div class="window-header">
                <h2>New sighting</h2>
                <button class="close-btn" @click="$emit('close')">×</button>
            </div>
            <label>Title</label>
            <input v-model="title" type="text" placeholder="UFO over the city" />
            <label>Image URL</label>
            <input v-model="imageUrl" type="text" placeholder="https://..." />
            <label>Description</label>
            <textarea v-model="description" rows="3" placeholder="What did you see?" />
            <div class="coords-row">
                <div>
                    <label>Latitude</label>
                    <input v-model.number="latitude" type="number" step="0.0001" placeholder="50.45" />
                </div>
                <div>
                    <label>Longitude</label>
                    <input v-model.number="longitude" type="number" step="0.0001" placeholder="30.52" />
                </div>
            </div>
            <p v-if="error" class="error">{{ error }}</p>
            <button class="submit-btn" :disabled="submitting" @click="submit">
                {{ submitting ? "Posting..." : "Post" }}
            </button>
        </div>
    </div>
</template>
<script setup>
import { ref } from "vue";
import { createPost } from "../api/client";
const emit = defineEmits(["close", "created"]);
const title = ref("");
const imageUrl = ref("");
const description = ref("");
const latitude = ref(null);
const longitude = ref(null);
const error = ref("");
const submitting = ref(false);
async function submit() {
    if (!title.value.trim()) {
        error.value = "Title is required";
        return;
    }
    if (latitude.value === null || longitude.value === null) {
        error.value = "Latitude and longitude are required";
        return;
    }
    submitting.value = true;
    error.value = "";
    try {
        const post = await createPost({
            title: title.value,
            image_url: imageUrl.value,
            description: description.value,
            latitude: latitude.value,
            longitude: longitude.value,
        });
        emit("created", post);
        emit("close");
    } catch (e) {
        error.value = "Could not create post";
    } finally {
        submitting.value = false;
    }
}
</script>
<style scoped>
.overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.35);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
}

.window {
    background: #ffffff;
    border-radius: 10px;
    padding: 20px 24px 24px;
    width: 320px;
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.18);
    font-family: -apple-system, "Segoe UI", Roboto, sans-serif;
}

.window-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 4px;
}

.window-header h2 {
    font-size: 18px;
    margin: 0;
    color: #1a1a1a;
}

.close-btn {
    border: none;
    background: none;
    font-size: 20px;
    line-height: 1;
    cursor: pointer;
    color: #666;
}

label {
    display: block;
    font-size: 12px;
    color: #555;
    margin: 10px 0 4px;
}

input,
textarea {
    width: 100%;
    box-sizing: border-box;
    padding: 8px 10px;
    border: 1px solid #d7d7d7;
    border-radius: 6px;
    font-size: 14px;
    font-family: inherit;
}

.coords-row {
    display: flex;
    gap: 10px;
}

.coords-row>div {
    flex: 1;
}

.error {
    color: #d33;
    font-size: 12px;
    margin-top: 8px;
}

.submit-btn {
    margin-top: 16px;
    width: 100%;
    padding: 10px;
    border: none;
    border-radius: 6px;
    background: #4caf50;
    color: white;
    font-size: 14px;
    font-weight: 600;
    cursor: pointer;
}

.submit-btn:disabled {
    opacity: 0.6;
    cursor: default;
}
</style>