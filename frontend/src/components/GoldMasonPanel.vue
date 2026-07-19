<template>
    <div class="backdrop">
        <div class="panel-grid">
            <div class="title-bar">
                <span>Gold mason control panel</span>
                <button class="close-btn" @click="$emit('close')">×</button>
            </div>
            <div class="box broadcast-box">
                <h3>Broadcast</h3>
                <textarea v-model="broadcastMessage" rows="3" placeholder="Message" />
                <label v-for="s in statuses" :key="s" class="checkbox-row">
                    <input type="checkbox" :value="s" v-model="selectedStatuses" />
                    {{ s[0].toUpperCase() + s.slice(1) }}
                </label>
                <button class="action-btn" :disabled="sendingBroadcast" @click="doBroadcast">
                    {{ sendingBroadcast ? "Sending..." : "Send" }}
                </button>
                <p v-if="broadcastMsg" class="feedback">{{ broadcastMsg }}</p>
            </div>
            <div class="box promotion-box">
                <h3>Promotion</h3>
                <input v-model="userQuery" type="text" placeholder="username search" @input="doSearch" />
                <ul v-if="searchResults.length" class="results">
                    <li v-for="u in searchResults" :key="u.id"
                        :class="{ selected: selectedUser && selectedUser.id === u.id }" @click="selectedUser = u">
                        {{ u.username }}
                    </li>
                </ul>
                <select v-model="chosenStatus">
                    <option v-for="s in statuses" :key="s" :value="s">{{ s }}</option>
                </select>
                <button class="action-btn" :disabled="!selectedUser" @click="doPromote">
                    Promote
                </button>
                <p v-if="promotionMsg" class="feedback">{{ promotionMsg }}</p>
            </div>
            <div class="box control-box">
                <button class="link-btn ban-link" @click="showBan = !showBan">BAN</button>
                <button class="link-btn" @click="showInvite = !showInvite">invite</button>
                <button class="delete-btn" @click="doDeleteAll">Delete all data</button>
                <BanWindow v-if="showBan" @close="showBan = false" />
            </div>
            <div v-if="showInvite" class="box invite-box">
                <h3>Invite</h3>
                <input v-model="inviteEmail" type="text" placeholder="Email" />
                <button class="action-btn" :disabled="!inviteEmail" @click="doInvite">Send</button>
                <p v-if="inviteMsg" class="feedback">{{ inviteMsg }}</p>
            </div>
        </div>
    </div>
</template>
<script setup>
import { ref } from "vue";
import {
    deleteAllData,
    inviteUser,
    promoteUser,
    searchUsers,
    sendBroadcast,
} from "../api/client";
import BanWindow from "./BanWindow.vue";
const statuses = ["bronze", "silver", "golden"];
const broadcastMessage = ref("");
const selectedStatuses = ref([...statuses]);
const sendingBroadcast = ref(false);
const broadcastMsg = ref("");
async function doBroadcast() {
    if (!broadcastMessage.value.trim()) return;
    sendingBroadcast.value = true;
    broadcastMsg.value = "";
    try {
        await sendBroadcast(broadcastMessage.value, selectedStatuses.value);
        broadcastMsg.value = "Broadcast sent";
        broadcastMessage.value = "";
    } catch (e) {
        broadcastMsg.value = "Could not send broadcast";
    } finally {
        sendingBroadcast.value = false;
    }
}
const userQuery = ref("");
const searchResults = ref([]);
const selectedUser = ref(null);
const chosenStatus = ref("bronze");
const promotionMsg = ref("");
async function doSearch() {
    if (!userQuery.value) {
        searchResults.value = [];
        return;
    }
    try {
        searchResults.value = await searchUsers(userQuery.value);
    } catch (e) {
        searchResults.value = [];
    }
}
async function doPromote() {
    if (!selectedUser.value) return;
    try {
        await promoteUser(selectedUser.value.id, chosenStatus.value);
        promotionMsg.value = `${selectedUser.value.username} is now ${chosenStatus.value}`;
    } catch (e) {
        promotionMsg.value = "Could not promote user";
    }
}
const showBan = ref(false);
const showInvite = ref(false);
const inviteEmail = ref("");
const inviteMsg = ref("");
async function doInvite() {
    try {
        await inviteUser(inviteEmail.value);
        inviteMsg.value = `Invite sent to ${inviteEmail.value}`;
        inviteEmail.value = "";
    } catch (e) {
        inviteMsg.value = "Could not send invite";
    }
}
async function doDeleteAll() {
    if (!confirm("This will delete all posts. Are you sure?")) return;
    try {
        await deleteAllData();
        alert("All data deleted");
    } catch (e) {
        alert("Could not delete data");
    }
}
</script>
<style scoped>
.backdrop {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.35);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    font-family: -apple-system, "Segoe UI", Roboto, sans-serif;
}

.panel-grid {
    background: #f2f2f2;
    border-radius: 14px;
    padding: 20px;
    width: min(90vw, 820px);
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 16px;
    position: relative;
}

.title-bar {
    grid-column: 1 / -1;
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-weight: 600;
    color: #1a1a1a;
    margin-bottom: 4px;
}

.close-btn {
    border: none;
    background: none;
    font-size: 20px;
    cursor: pointer;
    color: #666;
}

.box {
    background: #ffffff;
    border: 1px solid #e2e2e2;
    border-radius: 10px;
    padding: 14px 16px;
}

.box h3 {
    margin: 0 0 10px;
    font-size: 14px;
    color: #1a1a1a;
}

textarea,
input,
select {
    width: 100%;
    box-sizing: border-box;
    padding: 8px 10px;
    border: 1px solid #d7d7d7;
    border-radius: 6px;
    font-size: 13px;
    font-family: inherit;
    margin-bottom: 8px;
}

.checkbox-row {
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 13px;
    color: #444;
    margin-bottom: 4px;
}

.checkbox-row input {
    width: auto;
    margin: 0;
}

.action-btn {
    padding: 8px 14px;
    border: none;
    border-radius: 6px;
    background: #4caf50;
    color: white;
    font-weight: 600;
    cursor: pointer;
    font-size: 13px;
}

.action-btn:disabled {
    opacity: 0.5;
    cursor: default;
}

.feedback {
    font-size: 12px;
    color: #555;
    margin-top: 8px;
}

.results {
    list-style: none;
    margin: 0 0 8px;
    padding: 0;
    max-height: 90px;
    overflow-y: auto;
    border: 1px solid #eee;
    border-radius: 6px;
}

.results li {
    padding: 6px 8px;
    font-size: 13px;
    cursor: pointer;
}

.results li:hover,
.results li.selected {
    background: #eef6ee;
}

.control-box {
    display: flex;
    flex-direction: column;
    gap: 10px;
    align-items: flex-start;
    position: relative;
}

.link-btn {
    border: 1px solid #d7d7d7;
    background: #fafafa;
    border-radius: 6px;
    padding: 8px 14px;
    cursor: pointer;
    font-size: 13px;
}

.ban-link {
    font-weight: 700;
    font-style: italic;
}

.delete-btn {
    border: none;
    background: #fbdada;
    color: #b02a2a;
    border-radius: 8px;
    padding: 10px 14px;
    cursor: pointer;
    font-weight: 600;
    width: 100%;
}
</style>
