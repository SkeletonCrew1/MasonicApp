<template>
  <div class="page-background">
    <!-- Навбар з home.css -->
    <nav class="navbar">
      <div class="logo">Cult of the Tree</div>
      <ul class="nav-links">
        <li>
          <router-link to="/voting-page">Back to List</router-link>
        </li>
      </ul>
    </nav>

    <div class="container">
      <div class="neu-card">
        <h2 class="title">Create Voting</h2>

        <div class="form-group">
          <p class="label-text">Voting Category</p>
          <div class="radio-group">
            <label class="neu-radio">
              <input type="radio" value="promote" v-model="votingCategory" />
              <div class="indicator"></div>
              <span>Promote</span>
            </label>
            <label class="neu-radio">
              <input type="radio" value="exclude" v-model="votingCategory" />
              <div class="indicator"></div>
              <span>Exclude</span>
            </label>
          </div>
        </div>

        <div class="form-group">
          <div class="neu-input">
            <input type="text" v-model="votingSubject" placeholder=" " />
            <label>Username</label>
          </div>
        </div>

        <button class="neu-button" @click="submitNewVoting">Submit</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import '../styles/addVoting.css';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { createVoting } from '../api/client.js'; 

const router = useRouter();
const votingCategory = ref('');
const votingSubject = ref('');

const submitNewVoting = async () => {
  if (!votingCategory.value || !votingSubject.value.trim()) {
    alert('Error: empty fields');
    return;
  }

  try {
    await createVoting(votingSubject.value.trim(), votingCategory.value);
    router.push('/voting-page');
  } catch (error) {
    alert(`Backend error: ${error.message}`);
  }
};
</script>