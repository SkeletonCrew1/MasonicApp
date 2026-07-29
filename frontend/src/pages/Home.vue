<template>
  <div class="page">
    <nav class="navbar">
      <div class="logo">Cult of the Tree</div>
      <ul class="nav-links">
        <li><a href="#">Register</a></li>
        <li><a href="#">Login</a></li>
      </ul>
    </nav>

    <div class="coordinates">
      <strong>Latitude:</strong> {{ lat }}<br />
      <strong>Longitude:</strong> {{ lng }}
    </div>

    <button v-if="showAddButton" class="add-sighting-button" @click="goToAddSighting">
      Add Sighting
    </button>

    <div id="map"></div>
  </div>
</template>

<script setup>
import "../styles/home.css";
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import L from "leaflet";
import "leaflet/dist/leaflet.css";

const customIcon = L.divIcon({
  className: 'custom-marker',
  iconSize: [16, 16],
  iconAnchor: [8, 8]
});

const router = useRouter();
const lat = ref("Click map");
const lng = ref("Click map");
const showAddButton = ref(false);
const markersLayer = L.featureGroup();
const POSTING_SERVICE_URL = import.meta.env.VITE_POSTING_SERVICE_URL

async function loadSightings(map) {
  try {
    const response = await fetch(`${POSTING_SERVICE_URL}/sightings`);
    if (!response.ok) throw new Error("Failed to fetch");
    const sightings = await response.json();

    sightings.forEach(sighting => {
      const popupContent = document.createElement('div');
      popupContent.className = 'sighting-popup';
      popupContent.innerHTML = `
        <h3>${sighting.SightingName}</h3>
        ${sighting.SightingPicture ? `<img src="${POSTING_SERVICE_URL}${sighting.SightingPicture}" class="popup-image" />` : ''}
        <button class="open-page-btn" data-id="${sighting.SightingId}">Open Page</button>
      `;

      popupContent.querySelector('.open-page-btn').addEventListener('click', (e) => {
        const id = e.target.getAttribute('data-id');
        router.push(`/sighting/${id}`);
      });

      L.marker([sighting.SightingLatitude, sighting.SightingLongitude], { icon: customIcon })
        .addTo(map) 
        .bindPopup(popupContent);
    });
  } catch (err) {
    console.error("Error loading sightings:", err);
  }
}

onMounted(() => {
  const bounds = [
    [-85, -Infinity],
    [85, Infinity]
  ];

  const map = L.map("map", { 
    minZoom: 2, 
    maxZoom: 20,
    maxBounds: bounds,
    maxBoundsViscosity: 1.0,
    worldCopyJump: true,
    zoomControl: false
  }).setView([20, 0], 2);

  L.tileLayer("https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}{r}.png", {
    attribution: '&copy; OpenStreetMap contributors &copy; CARTO',
    noWrap: false
  }).addTo(map);

  loadSightings(map);

  let marker = null;
  map.on("click", (e) => {
    lat.value = e.latlng.lat.toFixed(6);
    lng.value = e.latlng.lng.toFixed(6);
    sessionStorage.setItem("latitude", lat.value);
    sessionStorage.setItem("longitude", lng.value);
    showAddButton.value = true;
    if (marker) {
      marker.setLatLng(e.latlng);
    } else {
      marker = L.marker(e.latlng, { icon: customIcon } ).addTo(map);
    }
  });
});

function goToAddSighting() {
  router.push("/add-sighting");
}
</script>