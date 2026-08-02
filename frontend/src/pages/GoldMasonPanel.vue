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
                        <input v-model="banUserId" type="number" placeholder="User ID (required)" style="margin-bottom: 5px;" />
                        <input v-model="banIpAddress" type="text" placeholder="IP Address (optional)" />
                        <button class="action-btn right-align" :disabled="!banUserId" @click="doBan">BAN</button>
                        <p v-if="banMsg" class="feedback">{{ banMsg }}</p>
                    </div>
                </div>
                <div class="col center-col">
                    <button class="delete-btn" @click="doDeleteAll">Delete all data</button>
                </div>
                <div class="col right-col">
                    <div class="box">
                        <div class="box-header">Promotion <span>G</span></div>
                        <input v-model="promoteUsername" type="text" placeholder="Username" />
                        <button class="action-btn right-align mt-auto" :disabled="!promoteUsername"
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
import { deleteAllData, inviteUser, promoteUser, sendBroadcast, banIp } from "../api/client";
const emit = defineEmits(["close"]);
const statuses = ["bronze", "silver", "gold"];


const broadcastMessage = ref("");
const selectedStatuses = ref([...statuses]);
const sendingBroadcast = ref(false);
const broadcastMsg = ref("");

async function doBroadcast() {
    if (!broadcastMessage.value.trim() || !selectedStatuses.value.length) return;
    try {
        await sendBroadcast(broadcastMessage.value, selectedStatuses.value);
        broadcastMsg.value = "Broadcast message sucessfully sent!";
    } catch (e) {
        broadcastMsg.value = e.message || "Error sending broadcast!";
    } finally {
        sendingBroadcast.value = false;
    }
}


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


const promoteUsername = ref("");
const promotionMsg = ref("");

async function doPromote() {
    if (!promoteUsername.value) return;
    try {
        const data = await promoteUser(promoteUsername.value);
        promotionMsg.value = data.message || "User sucessfully promoted!";
    } catch (e) {
        promotionMsg.value = e.message || "Error!";
    }
}


const inviteEmail = ref("");
const inviteMsg = ref("");

async function doInvite() {
    try {
        await inviteUser(inviteEmail.value);
        inviteMsg.value = "Invite message sucessfully sent!";
    } catch (e) {
        inviteMsg.value = e.message || "Error!";
    }
}


async function doDeleteAll() {
    if (!confirm("Delete all data?")) return;
    try { await deleteAllData(); alert("Deleted"); }
    catch (e) { alert("Error"); }
}

</script>

<style scoped>
@import "../styles/GoldMasonPanel.css";
</style>