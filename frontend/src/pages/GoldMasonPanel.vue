<template>
    <div class="backdrop" @click.self="$emit('close')">
        <div class="panel-wrapper">
            <div class="title-bar">
                <h2>Gold mason control panel: G</h2>
                <button class="close-btn" @click="$emit('close')">×</button>
            </div>
            <div class="panel-layout">
                <div class="col left-col">
                    <div class="box">
                        <div class="box-header">Broadcast <span>G</span></div>
                        <div class="broadcast-body">
                            <textarea v-model="broadcastMessage" placeholder="Message"></textarea>
                            <div class="checkboxes">
                                <label v-for="s in statuses" :key="s">
                                    {{ s[0].toUpperCase() + s.slice(1) }}
                                    <input type="checkbox" :value="s" v-model="selectedStatuses" />
                                </label>
                            </div>
                        </div>
                        <button class="action-btn" :disabled="sendingBroadcast" @click="doBroadcast">Send</button>
                        <p v-if="broadcastMsg" class="feedback">{{ broadcastMsg }}</p>
                    </div>
                    <div class="box">
                        <div class="box-header">Broadcast History</div>
                        <div class="status-tabs" style="display: flex; gap: 5px; margin-bottom: 8px;">
                            <button v-for="s in statuses" :key="s" 
                                    :class="{ 'active-tab': activeHistoryTab === s }" 
                                    @click="loadHistory(s)"
                                    style="flex: 1; padding: 4px; cursor: pointer;">
                                {{ s }}
                            </button>
                        </div>
                        <ul class="history-list" v-if="historyItems.length" style="max-height: 120px; overflow-y: auto; padding-left: 15px; font-size: 12px;">
                            <li v-for="item in historyItems" :key="item.id" style="margin-bottom: 4px;">
                                <span style="color: #888;">{{ new Date(item.created_at).toLocaleTimeString() }}</span>: 
                                {{ item.message }}
                            </li>
                        </ul>
                        <p v-else class="empty-history" style="font-size: 12px; color: #888;">No history for {{ activeHistoryTab }}</p>
                    </div>
                </div>
                <div class="col center-col">
                    <button class="delete-btn" @click="doDeleteAll">Delete all data</button>
                </div>
                <div class="col right-col">
                    <div class="box">
                        <div class="box-header">Promotion <span>G</span></div>
                        <input v-model="userQuery" type="text" placeholder="username search" @input="doSearch" />
                        <ul v-if="searchResults.length" class="results">
                            <li v-for="u in searchResults" :key="u.id"
                                :class="{ selected: selectedUser && selectedUser.id === u.id }"
                                @click="selectedUser = u">
                                {{ u.username }}
                            </li>
                        </ul>
                        <select v-model="chosenStatus">
                            <option v-for="s in statuses" :key="s" :value="s">{{ s }}</option>
                        </select>
                        <button class="action-btn right-align mt-auto" :disabled="!selectedUser"
                            @click="doPromote">Promote</button>
                        <p v-if="promotionMsg" class="feedback">{{ promotionMsg }}</p>
                    </div>
                    <div class="box">
                        <div class="box-header">Invite <span>G</span></div>
                        <input v-model="inviteEmail" type="text" placeholder="Email" />
                        <button class="action-btn right-align" :disabled="!inviteEmail" @click="doInvite">Send</button>
                        <p v-if="inviteMsg" class="feedback">{{ inviteMsg }}</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from "vue";
import { deleteAllData, inviteUser, promoteUser, searchUsers, sendBroadcast, banIp, getBroadcastHistory } from "../api/client";
const emit = defineEmits(["close"]);
const statuses = ["bronze", "silver", "golden"];
const broadcastMessage = ref("");
const selectedStatuses = ref([...statuses]);
const sendingBroadcast = ref(false);
const broadcastMsg = ref("");
const activeHistoryTab = ref("bronze");
const historyItems = ref([]);
async function loadHistory(status) {
    activeHistoryTab.value = status;
    try {
        historyItems.value = await getBroadcastHistory(status);
    } catch (e) {
        historyItems.value = [];
    }
}
async function doBroadcast() {
    if (!broadcastMessage.value.trim()) return;
    sendingBroadcast.value = true;
    try {
        const data = await sendBroadcast(broadcastMessage.value, selectedStatuses.value);
        broadcastMsg.value = `Broadcast sent to ${data.recipients} users`; 
        broadcastMessage.value = "";
        loadHistory(activeHistoryTab.value);
    } catch (e) { 
        broadcastMsg.value = e.message || "Error sending broadcast"; 
    }
    finally { 
        sendingBroadcast.value = false; 
    }
}
loadHistory("bronze");
const banUserId = ref("");
const banIpAddress = ref("");
const banMsg = ref("");
async function doBan() {
    try { 
        await banIp(banUserId.value, banIpAddress.value); 
        banMsg.value = "User/IP Banned successfully"; 
        banUserId.value = ""; 
        banIpAddress.value = ""; 
    }
    catch (e) { 
        banMsg.value = e.message || "Error banning user"; 
    }
}
const userQuery = ref("");
const searchResults = ref([]);
const selectedUser = ref(null);
const chosenStatus = ref("bronze");
const promotionMsg = ref("");
async function doSearch() {
    if (!userQuery.value) { searchResults.value = []; return; }
    try { searchResults.value = await searchUsers(userQuery.value); } catch (e) { searchResults.value = []; }
}
async function doPromote() {
    if (!selectedUser.value) return;
    try { await promoteUser(selectedUser.value.id, chosenStatus.value); promotionMsg.value = "Promoted"; }
    catch (e) { promotionMsg.value = "Error"; }
}
const inviteEmail = ref("");
const inviteMsg = ref("");
async function doInvite() {
    try { await inviteUser(inviteEmail.value); inviteMsg.value = "Sent"; inviteEmail.value = ""; }
    catch (e) { inviteMsg.value = "Error"; }
}
async function doDeleteAll() {
    if (!confirm("Delete all data?")) return;
    try { await deleteAllData(); alert("Deleted"); }
    catch (e) { alert("Error"); }
}
</script>

<style scoped>
@import "../styles/GoldMasonPanel.css";
.active-tab {
    background-color: #333;
    color: #fff;
    font-weight: bold;
}
</style>