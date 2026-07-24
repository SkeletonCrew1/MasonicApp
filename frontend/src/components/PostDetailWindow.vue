<template>
    <div class="overlay" @click.self="$emit('close')">
        <div class="detail-window">
            <div class="header">
                <h2>Page with post</h2>
                <span class="rank-badge">B/S/G</span>
            </div>
            <div class="content-split">
                <div class="left-pane">
                    <div class="map-placeholder">
                        <span class="label">map</span>
                        <div class="pin" />
                    </div>
                </div>
                <div class="right-pane">
                    <div class="field title-field">{{ post.title || 'Name' }}</div>
                    <div class="field desc-field">{{ post.description || 'Description' }}</div>
                    <div class="field loc-field">{{ post.latitude }}, {{ post.longitude }}</div>
                    <div class="field photo-field">
                        <img v-if="post.image_url" :src="post.image_url" />
                        <span v-else>Photo</span>
                    </div>
                </div>
            </div>
            <div class="footer">
                <p v-if="error" class="error">{{ error }}</p>
                <button class="approve-btn" :disabled="approving" @click="doApprove">
                    Leave approve
                </button>
            </div>
        </div>
    </div>
</template>
<script setup>
import { ref } from "vue";
import { approvePost } from "../api/client";
const props = defineProps({ post: { type: Object, required: true } });
const emit = defineEmits(["close", "approved"]);
const approving = ref(false);
const error = ref("");
async function doApprove() {
    approving.value = true;
    try {
        const updated = await approvePost(props.post.id);
        emit("approved", updated);
    } catch (e) { error.value = "Error"; }
    finally { approving.value = false; }
}
</script>
<style scoped>
.overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.3);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
}

.detail-window {
    background: #ffffff;
    border: 1px solid #777;
    border-radius: 12px;
    padding: 24px;
    width: 700px;
    display: flex;
    flex-direction: column;
    font-family: -apple-system, "Segoe UI", Roboto, sans-serif;
}

.header {
    display: flex;
    justify-content: space-between;
    margin-bottom: 20px;
}

.header h2 {
    margin: 0;
    font-size: 20px;
    font-weight: normal;
    color: #1a1a1a;
}

.rank-badge {
    font-size: 20px;
    color: #333;
}

.content-split {
    display: flex;
    gap: 24px;
}

.left-pane {
    flex: 1;
    display: flex;
}

.map-placeholder {
    flex: 1;
    border: 1px solid #777;
    border-radius: 8px;
    position: relative;
    background: #f0f0f0;
    min-height: 250px;
}

.map-placeholder .label {
    position: absolute;
    top: 10px;
    left: 12px;
    color: #888;
    font-size: 13px;
}

.pin {
    position: absolute;
    top: 30%;
    right: 20%;
    width: 12px;
    height: 12px;
    border-radius: 50%;
    border: 2px solid #555;
    background: #fff;
}

.right-pane {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 16px;
    justify-content: flex-start;
}

.field {
    border: 1px solid #777;
    border-radius: 8px;
    padding: 10px;
    text-align: center;
    background: #f0f0f0;
    font-size: 14px;
    color: #333;
}

.photo-field {
    min-height: 80px;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0;
    overflow: hidden;
    color: #888;
}

.photo-field img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.footer {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    margin-top: 24px;
}

.approve-btn {
    border: 1px solid #777;
    border-radius: 8px;
    padding: 10px 20px;
    background: #f0f0f0;
    cursor: pointer;
    font-size: 14px;
    font-weight: 600;
}

.error {
    color: #d33;
    margin-right: 16px;
    font-size: 13px;
}
</style>