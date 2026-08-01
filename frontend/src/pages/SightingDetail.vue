<template>
  <div class="page-wrapper">
    <div v-if="sighting" class="detail-container">

      <div class="map-column">
        <div id="sight-map"></div>
      </div>

      <div class="content-column">
        <button class="back-btn" @click="$router.push('/home')">← Back to Map</button>
        
        <h1 class="title">{{ sighting.SightingName }}</h1>
        
        <div class="image-wrapper">
          <img v-if="sighting.SightingPicture" 
               :src="`${POSTING_SERVICE_URL}` + sighting.SightingPicture" 
               :alt="sighting.SightingName" class="detail-image" />
        </div>

        <div class="info-grid">
          <div class="info-item">
            <span class="label">Description</span>
            <span class="value">{{ sighting.SightingDescription }}</span>
          </div>
          <div class="info-item">
            <span class="label">Discovery Date</span>
            <span class="value">{{ sighting.SightingDiscoveryDate }}</span>
          </div>
          <div class="info-item">
            <span class="label">Coordinates</span>
            <span class="value">
              {{ sighting.SightingLatitude.toFixed(6) }}, {{ sighting.SightingLongitude.toFixed(6) }}
            </span>
          </div>

          <div class="info-item approval-section">
            <button 
              class="submit-btn seen-too-btn" 
              :class="{ 'already-seen': sighting.HasSeenIt }"
              @click="markSeenToo"
            >
              {{ sighting.HasSeenIt ? 'You saw it too!' : 'I saw it too!' }}
            </button>
            <span class="label" style="margin-top: 8px;">Seen by: {{ sighting.SeenCount }} users</span>
          </div>
        </div>
      </div>

    </div>

    <div v-else class="loading">Loading sighting details...</div>
  </div>
</template>

<script setup>
import "../styles/sightingDetail.css";
import { ref, onMounted, nextTick } from 'vue';
import { useRoute } from 'vue-router';
import L from "leaflet";
import "leaflet/dist/leaflet.css";

const customIcon = L.divIcon({
  className: 'custom-marker',
  iconSize: [16, 16],
  iconAnchor: [8, 8]
});

const route = useRoute();
const sighting = ref(null);
const POSTING_SERVICE_URL = import.meta.env.VITE_POSTING_SERVICE_URL;
const AUTH_SERVICE_URL = import.meta.env.VITE_AUTH_SERVICE_URL;

onMounted(async() =>{
  try {
    const response = await fetch(`${AUTH_SERVICE_URL}/protected`, {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
    credentials : "include"
    });
    
    if (!response.ok) {
      router.push('/');
    } 
  } catch (error) {
    console.error(error);
    alert("Cannot connect to the server.");
  }
})

onMounted(async () => {
  const id = route.params.id;

  try {
    const response = await fetch(`${POSTING_SERVICE_URL}/sighting?id=${id}`, {
      credentials: "include"
    });
    if (!response.ok) throw new Error("Failed to fetch sighting details");
    sighting.value = await response.json();
  } catch (err) {
    console.error("Failed to load details.", err);
    return;
  }

  await nextTick();

  const bounds = [
    [-85, -Infinity],
    [85, Infinity]
  ];

  const map = L.map("sight-map", { 
    minZoom: 2, 
    maxZoom: 20,
    maxBounds: bounds,
    maxBoundsViscosity: 1.0,
    worldCopyJump: true,
    zoomControl: false
  }).setView([sighting.value.SightingLatitude, sighting.value.SightingLongitude], 7);

  L.tileLayer("https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}{r}.png", {
    attribution: '&copy; OpenStreetMap contributors &copy; CARTO',
    noWrap: false
  }).addTo(map);

  L.marker([sighting.value.SightingLatitude, sighting.value.SightingLongitude], { icon: customIcon })
    .addTo(map);
});

async function markSeenToo() {
  try {
    const response = await fetch(`${POSTING_SERVICE_URL}/seen-too`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ postId: String(sighting.value.SightingId) }),
      credentials: "include"
    });

    if (response.ok) {
      if (!sighting.value.HasSeenIt) {
        sighting.value.HasSeenIt = true;
        sighting.value.SeenCount += 1;
      }
    } else {
      alert("Some error has occured, please, try again later.");
    }
  } catch (err) {
    console.error("Error connecting to server", err);
  }
}
</script>