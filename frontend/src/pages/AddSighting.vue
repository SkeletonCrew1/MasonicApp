<template>
  <div class="container">
    <h1>Add Sighting</h1>

    <form @submit.prevent="submit">
      <label>Latitude</label>
      <input type="text" v-model="latitude"/>

      <label>Longitude</label>
      <input type="text" v-model="longitude"/>

      <label>Picture</label>
      <input type="file" ref="pictureInput" accept="image/*"/>

      <label>Name</label>
      <input type="text" v-model="name"/>

      <label>Date of discovery</label>
      <input type="date" v-model="date"/>

      <button type="submit">
        Submit
      </button>
    </form>
  </div>
</template>

<script setup>
import "../styles/addSighting.css";
import { ref } from "vue";

const latitude = ref(sessionStorage.getItem("latitude") || "");
const longitude = ref(sessionStorage.getItem("longitude") || "");
const name = ref("");
const date = ref("");
const pictureInput = ref(null);

async function submit() {
    let picture = "";
    const file = pictureInput.value?.files?.[0];

    if (file) {
        picture = await fileToBase64(file);
    }

    const sighting = {
        latitude: Number(latitude.value),
        longitude: Number(longitude.value),
        name: name.value,
        date: date.value,
        picture: picture
    };

    try {
        const response = await fetch("http://localhost:8080/sightings", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(sighting)
        });

        if (response.ok) {
            alert("Sighting submitted successfully!");
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