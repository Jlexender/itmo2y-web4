<script setup>
import {inject, onMounted, onUnmounted, ref} from "vue";

const bgPath = ref(null);
const time = new Date().getHours();
if (time >= 10 && time < 18) {
  bgPath.value = new URL('@/assets/img/ext_beach_day.jpg', import.meta.url);
} else if (time >= 4 && time < 10) {
  bgPath.value = new URL('@/assets/img/ext_beach_sunset.jpg', import.meta.url);
} else {
  bgPath.value = new URL('@/assets/img/ext_beach_night.jpg', import.meta.url);
}

const volume = inject('volume');
const bgMusic = new Audio(new URL('@/assets/audio/lets_be_friends.ogg', import.meta.url));

onMounted(() => {
  setTimeout(() => {
    bgMusic.volume = volume.value;
    bgMusic.loop = true;
    bgMusic.play();
  }, 1000);
});

onUnmounted(() => {
  setInterval(() => {
    if (bgMusic.volume > 0.01) {
      bgMusic.volume -= 0.01;
    } else {
      bgMusic.pause();
    }
  }, 30);
});

defineEmits(['toMain']);
</script>

<template>
  <div>
    <img :src="bgPath" alt="Game background" class="game-canvas" draggable="false"/>

    <div class="to-menu" @click="$emit('toMain')">
      В меню
    </div>
  </div>
</template>

<style scoped>

.to-menu {
  position: absolute;
  left: 10px;
  top: 10px;
  font-size: 48px;
  font-family: "Century Gothic", sans-serif;
  color: white;
  transition: 0.5s;
}

.to-menu:hover {
  text-shadow: 0 0 20px white;
  transition: 0.5s;
}

</style>
