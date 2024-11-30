<script setup>
import {ref} from 'vue';

const logoPath = ref(null);

const time = new Date().getHours();

if (time >= 10 && time < 18) {
  logoPath.value = new URL('@/assets/img/logo_day.png', import.meta.url);
} else if (time >= 4 && time < 10) {
  logoPath.value = new URL('@/assets/img/logo_sunset.png', import.meta.url);
} else {
  logoPath.value = new URL('@/assets/img/logo_night.png', import.meta.url);
}

</script>

<template>
  <div>
    <img :src="logoPath" alt="Logo" draggable="false"/>
  </div>
</template>

<style scoped>
img {
  position: absolute;
  top: 30%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 65%;
}
</style>
