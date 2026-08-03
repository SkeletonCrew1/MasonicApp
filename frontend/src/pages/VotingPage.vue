<template>
  <div class="page-background">
    <nav class="navbar">
      <router-link to="/home" class="logo">Cult of the Tree</router-link>
      <ul class="nav-links">
        <li v-if="currentUserStatus?.trim().toLowerCase() === 'gold'">
          <router-link to="/control-panel">Gold User page</router-link>
        </li>
        <li v-if="currentUserStatus?.trim().toLowerCase() === 'gold'">
          <router-link to="/add-voting">Add voting</router-link>
        </li>
        <li class="user-greeting" v-if="currentUsername">
          {{ currentUsername }}
        </li>
        <li>
          <a href="#" @click.prevent="Logout">Logout</a>
        </li>
      </ul>
    </nav>

    <div class="container">
      <div class="neu-card">
        <h2 class="title">Active Votings</h2>
        
        <main class="votings-list">
          <div v-for="voting in votings" :key="voting.voting_id" class="neu-item">
            <span class="voting-text">
              {{ capitalize(voting.category) }} {{ voting.username }}
            </span>
            <button 
              class="neu-button-small"
              :class="{ 'voted': voting.is_approved }"
              :disabled="voting.is_approved" 
              @click="submitVote(voting.voting_id)"
            >
              {{ voting.is_approved ? 'Voted' : 'Vote' }}
            </button>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script setup>
import '../styles/votingPage.css';
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getVotings, addVote } from '../api/client.js';

const router = useRouter();
const votings = ref([]);

const currentUserId = ref(null);
const currentUserStatus = ref(null);
const POSTING_SERVICE_URL = import.meta.env.VITE_POSTING_SERVICE_URL;
const AUTH_SERVICE_URL = import.meta.env.VITE_AUTH_SERVICE_URL;

const checkUserStatus = async () => {
  try {
    const response = await fetch(`${POSTING_SERVICE_URL}/user-status`, {
      credentials: "include"
    });
    if (response.ok) {
      const data = await response.json();
      currentUserStatus.value = data.status;
      currentUserId.value = data.userid || data.id || 1; 
      return true;
    }
    return false;
  } catch (err) {
    console.error("Failed to fetch user status", err);
    return false;
  }
};

const fetchVotingsData = async () => {
  try {
    const data = await getVotings(currentUserId.value, currentUserStatus.value);
    votings.value = data.votings;
  } catch (error) {
    console.error('API Error:', error.message);
  }
};

const submitVote = async (votingId) => {
  try {
    await addVote(votingId, currentUserId.value);
    await fetchVotingsData(); 
  } catch (error) {
    alert(error.message || 'Конфлікт голосування');
  }
};

const goToAddVoting = () => router.push('/add-voting');

const Logout = async () => {
  try {
    const response = await fetch(`${AUTH_SERVICE_URL}/logout`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include"
    });
    if (response.ok) {
      router.push('/');
    }
  } catch (error) {
    console.error(error);
  }
};

const capitalize = (str) => str.charAt(0).toUpperCase() + str.slice(1);

onMounted(async () => {
  const isAuth = await checkUserStatus();
  if (!isAuth) {
    router.push('/');
    return;
  }
  await fetchVotingsData();
});
</script>
