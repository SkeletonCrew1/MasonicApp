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
                        <div class="box-header">Ban agent <span>G</span></div>
                        <input v-model="banIpAddress" type="text" placeholder="Find somebody by IP" />
                        <button class="action-btn right-align" :disabled="!banIpAddress" @click="doBan">BAN</button>
                        <p v-if="banMsg" class="feedback">{{ banMsg }}</p>
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
import { deleteAllData, inviteUser, promoteUser, searchUsers, sendBroadcast, banIp } from "../api/client";
const emit = defineEmits(["close"]);
const statuses = ["bronze", "silver", "golden"];
const broadcastMessage = ref("");
const selectedStatuses = ref([...statuses]);
const sendingBroadcast = ref(false);
const broadcastMsg = ref("");
async function doBroadcast() {
    if (!broadcastMessage.value.trim()) return;
    sendingBroadcast.value = true;
    try {
        await sendBroadcast(broadcastMessage.value, selectedStatuses.value);
        broadcastMsg.value = "Broadcast sent"; broadcastMessage.value = "";
    } catch (e) { broadcastMsg.value = "Error sending"; }
    finally { sendingBroadcast.value = false; }
}
const banIpAddress = ref("");
const banMsg = ref("");
async function doBan() {
    try { await banIp(banIpAddress.value); banMsg.value = "IP Banned"; banIpAddress.value = ""; }
    catch (e) { banMsg.value = "Error banning"; }
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
.backdrop {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.3);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
}

.panel-wrapper {
    background: #ffffff;
    border-radius: 12px;
    padding: 24px;
    width: 900px;
    max-width: 95vw;
    border: 1px solid #777;
    font-family: -apple-system, "Segoe UI", Roboto, sans-serif;
}

.title-bar {
    display: flex;
    justify-content: space-between;
    margin-bottom: 20px;
}

.title-bar h2 {
    margin: 0;
    font-size: 20px;
    font-weight: normal;
    color: #1a1a1a;
}

.close-btn {
    background: none;
    border: none;
    font-size: 24px;
    cursor: pointer;
    color: #666;
}

.panel-layout {
    display: grid;
    grid-template-columns: 1fr auto 1fr;
    gap: 40px;
}

.col {
    display: flex;
    flex-direction: column;
    gap: 24px;
}

.center-col {
    justify-content: center;
    align-items: center;
}

.box {
    background: #ffffff;
    border: 1px solid #777;
    border-radius: 8px;
    padding: 16px;
    display: flex;
    flex-direction: column;
    min-height: 140px;
}

.box-header {
    display: flex;
    justify-content: space-between;
    margin-bottom: 12px;
    font-size: 16px;
    color: #1a1a1a;
}

.box-header span {
    font-weight: bold;
}

.broadcast-body {
    display: flex;
    gap: 16px;
    margin-bottom: 12px;
}

.broadcast-body textarea {
    flex: 1;
    resize: none;
    border-radius: 6px;
    border: 1px solid #777;
    padding: 8px;
    background: #ffffff;
    font-family: inherit;
    font-size: 14px;
}

.checkboxes {
    display: flex;
    flex-direction: column;
    gap: 8px;
    justify-content: center;
    font-size: 14px;
    color: #333;
}

.checkboxes label {
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 80px;
}

input,
select {
    width: 100%;
    padding: 8px;
    border-radius: 6px;
    border: 1px solid #777;
    background: #ffffff;
    margin-bottom: 12px;
    box-sizing: border-box;
    font-family: inherit;
    font-size: 14px;
}

.action-btn {
    padding: 6px 16px;
    border-radius: 6px;
    border: 1px solid #777;
    background: #f0f0f0;
    cursor: pointer;
    align-self: flex-start;
    font-size: 14px;
    font-weight: 600;
}

.right-align {
    align-self: flex-end;
}

.mt-auto {
    margin-top: auto;
}

.delete-btn {
    padding: 12px 24px;
    border-radius: 6px;
    border: 1px solid #ff4d4f;
    background: #ffd6d6;
    color: #a80000;
    cursor: pointer;
    font-weight: 600;
}

.results {
    list-style: none;
    padding: 0;
    margin: 0 0 10px;
    border: 1px solid #777;
    border-radius: 6px;
    max-height: 80px;
    overflow-y: auto;
    background: #ffffff;
}

.results li {
    padding: 4px 8px;
    cursor: pointer;
    font-size: 13px;
}

.results li.selected {
    background: #e0e0e0;
}

.feedback {
    font-size: 12px;
    color: #555;
    margin: 4px 0 0;
}
</style>