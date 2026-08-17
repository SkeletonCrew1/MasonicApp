<template>
  <div class="page-background">
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
            <label v-if="is_inquisitor" class="neu-radio">
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
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { createVoting } from '../api/client.js'; 
import Cookies from 'js-cookie';
// import VueJwtDecode from 'vue-jwt-decode'
// import VueCookies from 'vue-cookies'
const router = useRouter();
const votingCategory = ref('');
const votingSubject = ref('');
const is_inquisitor =ref(false) ;
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

onMounted(async () => {
  try {
    const response = await fetch('/api/auth/check_inquisitor', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      credentials: 'include'
    });

    if (!response.ok) {
      
      throw new Error(`HTTP error: ${response.status}`);
    }

    
    const data = await response.json();
    is_inquisitor.value = String(data["text"]).trim().toLowerCase() === 'true';
    
  } catch (error) {
    console.error(error);
    
    alert("Cannot connect to the server or process response.");
  }
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
.logo { font-size: 24px; font-weight: bold; color: #2D2E2F; }
.nav-links { display: flex; gap: 30px; list-style: none; margin: 0; padding: 0; }
.nav-links a { color: #2D2E2F; text-decoration: none; font-weight: 500; font-size: 16px; }
.nav-links a:hover { color: #6c7293; }

/* Neumorphic Container */
.container {
    padding: 40px 20px;
    display: flex;
    justify-content: center;
    align-items: flex-start;
}
.neu-card {
    background: #e0e5ec;
    border-radius: 30px;
    padding: 50px 40px;
    box-shadow: 20px 20px 60px #bec3cf, -20px -20px 60px #ffffff;
    width: 100%;
    max-width: 420px;
}
.title {
    color: #3d4468;
    font-size: 2rem;
    font-weight: 600;
    margin-bottom: 40px;
    text-align: center;
}

/* Custom Radio Buttons */
.form-group {
    margin-bottom: 30px;
}
.label-text {
    color: #9499b7;
    font-size: 16px;
    margin-bottom: 15px;
    font-weight: 500;
}
.radio-group {
    display: flex;
    flex-direction: column;
    gap: 15px;
}
.neu-radio {
    display: flex;
    align-items: center;
    gap: 15px;
    cursor: pointer;
    color: #6c7293;
    font-weight: 500;
    font-size: 16px;
}
.neu-radio input {
    display: none;
}
.neu-radio .indicator {
    width: 24px;
    height: 24px;
    border-radius: 50%;
    background: #e0e5ec;
    box-shadow: 4px 4px 10px #bec3cf, -4px -4px 10px #ffffff;
    transition: all 0.3s ease;
    position: relative;
}
.neu-radio input:checked + .indicator {
    box-shadow: inset 4px 4px 8px #bec3cf, inset -4px -4px 8px #ffffff;
}
.neu-radio input:checked + .indicator::after {
    content: '';
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 10px;
    height: 10px;
    background: #3d4468;
    border-radius: 50%;
}

/* Floating Label Input */
.neu-input {
    position: relative;
    background: #e0e5ec;
    border-radius: 15px;
    box-shadow: inset 8px 8px 16px #bec3cf, inset -8px -8px 16px #ffffff;
    transition: all 0.3s ease;
}
.neu-input input {
    width: 100%;
    background: transparent;
    border: none;
    padding: 20px 24px;
    color: #3d4468;
    font-size: 16px;
    font-weight: 500;
    outline: none;
}
.neu-input input::placeholder {
    color: transparent;
}
.neu-input label {
    position: absolute;
    left: 24px;
    top: 50%;
    transform: translateY(-50%);
    color: #9499b7;
    font-size: 16px;
    pointer-events: none;
    transition: all 0.3s ease;
}
.neu-input input:focus + label,
.neu-input input:not(:placeholder-shown) + label {
    top: 8px;
    font-size: 12px;
    color: #6c7293;
}

/* Main Button */
.neu-button {
    width: 100%;
    background: #e0e5ec;
    border: none;
    border-radius: 15px;
    padding: 18px 32px;
    color: #3d4468;
    font-size: 16px;
    font-weight: 600;
    cursor: pointer;
    box-shadow: 8px 8px 20px #bec3cf, -8px -8px 20px #ffffff;
    transition: all 0.3s ease;
    margin-top: 10px;
}
.neu-button:hover {
    transform: translateY(-2px);
    box-shadow: 12px 12px 30px #bec3cf, -12px -12px 30px #ffffff;
}
.neu-button:active {
    transform: translateY(0);
    box-shadow: inset 4px 4px 10px #bec3cf, inset -4px -4px 10px #ffffff;
}
</style>