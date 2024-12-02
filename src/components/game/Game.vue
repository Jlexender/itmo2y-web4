<script setup>
import {inject, onMounted, onUnmounted, provide, ref} from "vue";
import Plot from "@/components/game/Plot.vue";
import Girl from "@/components/game/Girl.vue";
import Blank from "@/components/game/Blank.vue";

const bgPath = ref(null);
const time = new Date().getHours();
if (time >= 10 && time < 18) {
  bgPath.value = new URL('@/assets/img/ext_beach_day.jpg', import.meta.url);
} else if (time >= 4 && time < 10) {
  bgPath.value = new URL('@/assets/img/ext_beach_sunset.jpg', import.meta.url);
} else {
  bgPath.value = new URL('@/assets/img/ext_beach_night.jpg', import.meta.url);
}

const cRadius = ref(0);
provide('cRadius', cRadius);

const volume = inject('volume');
const bgMusic = new Audio(new URL('@/assets/audio/lets_be_friends.ogg', import.meta.url));
const ambient = new Audio(new URL('@/assets/audio/boat_station_night.ogg', import.meta.url));

onMounted(() => {
  setTimeout(() => {
    bgMusic.volume = volume.value;
    bgMusic.loop = true;
    bgMusic.play();
    ambient.volume = volume.value;
    ambient.loop = true;
    ambient.play();
  }, 1000);
});

onUnmounted(() => {
  setInterval(() => {
    if (bgMusic.volume > 0.01) {
      bgMusic.volume -= 0.01;
      ambient.volume -= 0.01;
    } else {
      bgMusic.pause();
      ambient.pause();
    }
  }, 30);
});

defineEmits(['toMain']);

const notSad = ref(true);

const makeSad = () => {
  notSad.value = false;
}

const makeHappy = () => {
  notSad.value = true;
}

</script>

<template>
  <div>
    <img :src="bgPath" alt="Game background" class="game-canvas" draggable="false"/>

    <div class="to-menu" @click="$emit('toMain')">
      В меню
    </div>

    <Blank/>
    <Plot
        @makeHappy="makeHappy"
        @makeSad="makeSad"
    />
    <Girl :notSad="notSad"/>
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
  z-index: 1;
}

.to-menu:hover {
  color: blue;
  transition: 0.5s;
}

</style>
