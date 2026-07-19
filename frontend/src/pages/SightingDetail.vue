<template>
  <div class="page-wrapper">
    <div v-if="sighting" class="sighting-card">
      <button class="back-btn" @click="$router.push('/')">← Back to Map</button>
      
      <h1 class="title">{{ sighting.name }}</h1>
      
      <div class="image-wrapper">
        <img v-if="sighting.picture" :src="'http://localhost:8080' + sighting.picture" :alt="sighting.name" class="detail-image" />
      </div>

      <div class="info-grid">
        <div class="info-item">
          <span class="label">Discovery Date</span>
          <span class="value">{{ sighting.date }}</span>
        </div>
        <div class="info-item">
          <span class="label">Coordinates</span>
          <span class="value">{{ sighting.latitude.toFixed(6) }}, {{ sighting.longitude.toFixed(6) }}</span>
        </div>
      </div>
    </div>
    
    <div v-else class="loading">Loading sighting details...</div>
  </div>
</template>

<script setup>
import "../styles/sightingDetail.css";
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const sighting = ref(null);

onMounted(async () => {
  const id = route.params.id;
  try {
    const response = await fetch(`http://localhost:8080/sighting?id=${id}`);
    sighting.value = await response.json();
  } catch (err) {
    console.error("Failed to load details", err);
  }
});
</script>