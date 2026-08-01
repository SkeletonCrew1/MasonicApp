<template>
  <div class="page-wrapper">
    <div class="add-container">

      <div class="form-column">
        <button class="back-btn" @click="$router.push('/')">← Back to Map</button>
        <h1>Add Sighting</h1>

        <form @submit.prevent="submit">
          <div class="info-item">
            <label>Name</label>
            <input type="text" v-model="name" placeholder="Enter sighting name" required />
          </div>

          <div class="info-item">
            <label>Description</label>
            <textarea v-model="description" placeholder="Brief description"></textarea>
          </div>

          <div class="info-item">
            <label>Date of Discovery</label>
            <input type="date" v-model="date" required />
          </div>

          <div class="coordinates-row">
            <div class="info-item">
              <label>Latitude</label>
              <input type="text" step="any" v-model.number="latitude" @input="updateMarkerFromInputs" required />
            </div>
            <div class="info-item">
              <label>Longitude</label>
              <input type="text" step="any" v-model.number="longitude" @input="updateMarkerFromInputs" required />
            </div>
          </div>

          <div class="info-item">
            <label>Picture</label>
            <input type="file" ref="pictureInput" accept="image/*" />
          </div>

          <button type="submit" class="submit-btn">Submit Sighting</button>
        </form>
      </div>

      <div class="map-column">
        <div class="map-wrapper-box">
          <div id="picker-map"></div>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import "../styles/addSighting.css";
import { ref, onMounted, nextTick } from "vue";
import { useRouter } from "vue-router";
import L from "leaflet";
import "leaflet/dist/leaflet.css";

const customIcon = L.divIcon({
  className: 'custom-marker',
  iconSize: [16, 16],
  iconAnchor: [8, 8]
});

const router = useRouter();

const latitude = ref(Number(sessionStorage.getItem("latitude")) || 20);
const longitude = ref(Number(sessionStorage.getItem("longitude")) || 0);
const name = ref("");
const date = ref("");
const description = ref("");
const pictureInput = ref(null);
const POSTING_SERVICE_URL = import.meta.env.VITE_POSTING_SERVICE_URL

let map = null;
let marker = null;

onMounted(async () => {
  await nextTick();

  const bounds = [
    [-85, -Infinity],
    [85, Infinity]
  ];

  map = L.map("picker-map", {
    minZoom: 2,
    maxZoom: 20,
    maxBounds: bounds,
    maxBoundsViscosity: 1.0,
    worldCopyJump: true,
    zoomControl: false
  }).setView([latitude.value, longitude.value], latitude.value === 20 ? 2 : 7);

  L.tileLayer("https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}{r}.png", {
    attribution: '&copy; OpenStreetMap contributors &copy; CARTO',
    noWrap: false
  }).addTo(map);

  marker = L.marker([latitude.value, longitude.value], { draggable: true, icon: customIcon }).addTo(map);

  marker.on('dragend', (event) => {
    const pos = event.target.getLatLng();
    latitude.value = Number(pos.lat.toFixed(6));
    longitude.value = Number(pos.lng.toFixed(6));
  });

  map.on('click', (event) => {
    const pos = event.latlng;
    marker.setLatLng(pos);
    latitude.value = Number(pos.lat.toFixed(6));
    longitude.value = Number(pos.lng.toFixed(6));
  });
});

function updateMarkerFromInputs() {
  if (!isNaN(latitude.value) && !isNaN(longitude.value)) {
    const newPos = [latitude.value, longitude.value];
    if (marker) {
      marker.setLatLng(newPos);
      map.panTo(newPos);
    }
  }
}

async function submit() {
  let picture = "";
  const file = pictureInput.value?.files?.[0];

  if (file) {
    picture = await fileToBase64(file);
  }

  const sighting = {
    SightingLatitude: Number(latitude.value),
    SightingLongitude: Number(longitude.value),
    SightingName: name.value,
    SightingDiscoveryDate: date.value,
    SightingPicture: picture,
    SightingDescription: description.value
  };

  try {
    const response = await fetch(`${POSTING_SERVICE_URL}/sightings`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(sighting)
    });

    if (response.ok) {
      alert("Sighting submitted successfully!");
      router.push('/home');
    } else {
      alert("Failed to submit sighting.");
    }
  } catch (error) {
    console.error(error);
    alert("Cannot connect to the server.");
  }
}

function fileToBase64(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => resolve(reader.result);
    reader.onerror = reject;
    reader.readAsDataURL(file);
  });
}
</script>