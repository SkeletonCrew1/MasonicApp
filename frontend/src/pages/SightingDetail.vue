<template>
  <div class="page-wrapper">
    <div v-if="sighting" class="detail-container">

      <div class="map-column">
        <div id="sight-map"></div>
      </div>

      <div class="content-column">
        <button class="back-btn" @click="$router.push('/')">← Back to Map</button>
        
        <h1 class="title">{{ sighting.SightingName }}</h1>
        <h1 class="title">{{ sighting.SightingName }}</h1>
        
        <div class="image-wrapper">
          <img v-if="sighting.SightingPicture" :src="`${POSTING_SERVICE_URL}` + sighting.SightingPicture" :alt="sighting.SightingName" class="detail-image" />
        </div>

        <div class="info-grid">
          <div class="info-item">
            <span class="label">Description</span>
            <span class="value">{{ sighting.SightingDescription }}</span>
            <span class="value">{{ sighting.SightingDescription }}</span>
          </div>
          <div class="info-item">
            <span class="label">Discovery Date</span>
            <span class="value">{{ sighting.SightingDiscoveryDate }}</span>
            <span class="value">{{ sighting.SightingDiscoveryDate }}</span>
          </div>
          <div class="info-item">
            <span class="label">Coordinates</span>
            <span class="value">{{ sighting.SightingLatitude.toFixed(6) }}, {{ sighting.SightingLongitude.toFixed(6) }}</span>
            <span class="value">{{ sighting.SightingLatitude.toFixed(6) }}, {{ sighting.SightingLongitude.toFixed(6) }}</span>
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
const POSTING_SERVICE_URL = import.meta.env.VITE_POSTING_SERVICE_URL

onMounted(async () => {
  const id = route.params.id;

  try {
    const response = await fetch(`${POSTING_SERVICE_URL}/sighting?id=${id}`);
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
  }).setView([sighting.value.SightingLatitude, sighting.value.SightingLongitude], 7);

  L.tileLayer("https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}{r}.png", {
    attribution: '&copy; OpenStreetMap contributors &copy; CARTO',
    noWrap: false
  }).addTo(map);

  L.marker([sighting.value.SightingLatitude, sighting.value.SightingLongitude], { icon: customIcon })
  L.marker([sighting.value.SightingLatitude, sighting.value.SightingLongitude], { icon: customIcon })
    .addTo(map);
});
</script>