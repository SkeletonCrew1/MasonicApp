<template>
  <div class="panel">
    <div class="panel-header">
      <h3>Ban agent</h3>
      <button class="close-btn" @click="$emit('close')">×</button>
    </div>
    <input v-model="ip" type="text" placeholder="Find somebody by IP" />
    <button class="ban-btn" :disabled="banning || !ip" @click="doBan">
      {{ banning ? "Banning..." : "BAN" }}
    </button>
    <p v-if="message" class="message">{{ message }}</p>
    <div v-if="bans.length" class="ban-list">
      <p class="list-title">Banned IPs</p>
      <ul>
        <li v-for="b in bans" :key="b.ip">{{ b.ip }}</li>
      </ul>
    </div>
  </div>
</template>
<script setup>
import { onMounted, ref } from "vue";
import { banIp, listBans } from "../api/client";
const ip = ref("");
const banning = ref(false);
const message = ref("");
const bans = ref([]);
async function refreshBans() {
  try {
    bans.value = await listBans();
  } catch (e) {
  }
}
async function doBan() {
  banning.value = true;
  message.value = "";
  try {
    await banIp(ip.value);
    message.value = `${ip.value} has been banned`;
    ip.value = "";
    await refreshBans();
  } catch (e) {
    message.value = "Could not ban this IP";
  } finally {
    banning.value = false;
  }
}
onMounted(refreshBans);
</script>
<style scoped>
.panel {
  background: #ffffff;
  border: 1px solid #e2e2e2;
  border-radius: 10px;
  padding: 16px 18px;
  width: 260px;
  font-family: -apple-system, "Segoe UI", Roboto, sans-serif;
}

.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.panel-header h3 {
  margin: 0;
  font-size: 15px;
  color: #1a1a1a;
}

.close-btn {
  border: none;
  background: none;
  font-size: 18px;
  cursor: pointer;
  color: #666;
}

input {
  width: 100%;
  box-sizing: border-box;
  padding: 8px 10px;
  border: 1px solid #d7d7d7;
  border-radius: 6px;
  font-size: 13px;
  margin-bottom: 10px;
}

.ban-btn {
  width: 100%;
  padding: 9px;
  border: none;
  border-radius: 6px;
  background: #e5484d;
  color: white;
  font-weight: 700;
  font-style: italic;
  cursor: pointer;
}

.ban-btn:disabled {
  opacity: 0.5;
  cursor: default;
}

.message {
  font-size: 12px;
  color: #555;
  margin-top: 8px;
}

.ban-list {
  margin-top: 12px;
  border-top: 1px solid #eee;
  padding-top: 8px;
}

.list-title {
  font-size: 11px;
  color: #999;
  margin: 0 0 4px;
}

.ban-list ul {
  list-style: none;
  margin: 0;
  padding: 0;
  max-height: 100px;
  overflow-y: auto;
  font-size: 12px;
  color: #444;
}

.ban-list li {
  padding: 2px 0;
}
</style>
