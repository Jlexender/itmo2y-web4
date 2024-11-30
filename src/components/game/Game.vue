<script setup>
import {ref} from "vue";

const bgPath = ref(null);

const time = new Date().getHours();

if (time >= 10 && time < 18) {
  bgPath.value = new URL('@/assets/img/ext_beach_day.jpg', import.meta.url);
} else if (time >= 4 && time < 10) {
  bgPath.value = new URL('@/assets/img/ext_beach_sunset.jpg', import.meta.url);
} else {
  bgPath.value = new URL('@/assets/img/ext_beach_night.jpg', import.meta.url);
}
</script>

<template>
  <div>
    <img :src="bgPath" alt="Game background" class="game-canvas" draggable="false"/>
  </div>
</template>

<style scoped>

</style>
