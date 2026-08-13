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
    const response = await fetch(`/posting/user-status`, {
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
    alert(error.message || 'Voting conflict');
  }
};

const goToAddVoting = () => router.push('/add-voting');

const Logout = async () => {
  try {
    const response = await fetch(`/auth/logout`, {
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

<style scoped>
.page-background {
    min-height: 100vh;
    background: #e0e5ec;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}

/* Navbar */
.navbar {
    height: 60px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: #fcfbf5;
    padding: 0 20px;
    box-shadow: 0 2px 5px rgba(0,0,0,0.05);
}
.logo { font-size: 24px; font-weight: bold; color: #2D2E2F; text-decoration: none; }
.nav-links { display: flex; gap: 30px; list-style: none; margin: 0; padding: 0; }
.nav-links a { color: #2D2E2F; text-decoration: none; font-weight: 500; font-size: 16px; }
.nav-links a:hover { color: #6c7293; }

/* Neumorphic Container */
.container {
    padding: 40px 20px;
    display: flex;
    justify-content: center;
}
.neu-card {
    background: #e0e5ec;
    border-radius: 30px;
    padding: 40px;
    box-shadow: 20px 20px 60px #bec3cf, -20px -20px 60px #ffffff;
    width: 100%;
    max-width: 600px;
}
.title {
    color: #3d4468;
    font-size: 2rem;
    font-weight: 600;
    margin-bottom: 30px;
    text-align: center;
}

/* List Items */
.votings-list {
    display: flex;
    flex-direction: column;
    gap: 20px;
}
.neu-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 16px 24px;
    border-radius: 15px;
    background: #e0e5ec;
    box-shadow: inset 8px 8px 16px #bec3cf, inset -8px -8px 16px #ffffff;
}
.voting-text {
    color: #3d4468;
    font-weight: 500;
    font-size: 16px;
}

/* Small Button */
.neu-button-small {
    background: #e0e5ec;
    border: none;
    border-radius: 10px;
    padding: 10px 24px;
    color: #3d4468;
    font-size: 14px;
    font-weight: 600;
    cursor: pointer;
    box-shadow: 6px 6px 12px #bec3cf, -6px -6px 12px #ffffff;
    transition: all 0.3s ease;
}
.neu-button-small:hover:not(:disabled) {
    box-shadow: 8px 8px 16px #bec3cf, -8px -8px 16px #ffffff;
    transform: translateY(-2px);
}
.neu-button-small:active:not(:disabled) {
    box-shadow: inset 4px 4px 8px #bec3cf, inset -4px -4px 8px #ffffff;
    transform: translateY(0);
}
.neu-button-small.voted, .neu-button-small:disabled {
    color: #9499b7;
    box-shadow: inset 4px 4px 8px #bec3cf, inset -4px -4px 8px #ffffff;
    cursor: not-allowed;
}
</style>