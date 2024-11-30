<script setup>
import Logo from "@/components/intro/Logo.vue";
import {ref} from "vue";

const logo = ref(false);

const showLogo = () => {
  logo.value = true;
}

setTimeout(() => {
  showLogo();
}, 4000);

const imgPath = ref(null);

const time = new Date().getHours();

if (time >= 10 && time < 18) {
  imgPath.value = new URL('@/assets/img/splashscreen_day.png', import.meta.url);
} else if (time >= 4 && time < 10) {
  imgPath.value = new URL('@/assets/img/splashscreen_sunset.png', import.meta.url);
} else {
  imgPath.value = new URL('@/assets/img/splashscreen_night.png', import.meta.url);
}

const emit = defineEmits(['toMain']);

setTimeout(() => {
  emit('toMain');
}, 7000);
</script>

<template>
  <div>
    <img :src="imgPath" alt="Splashscreen" class="game-canvas"/>
    <transition name="fade" appear>
      <Logo v-if="logo"/>
    </transition>
  </div>

</template>

<style scoped>
@keyframes moveUp {
  0% {
    transform: translateY(0);
  }
  100% {
    transform: translateY(-50%);
  }
}

img {
  width: 1920px;
  height: 2160px;
  animation: moveUp 4s linear forwards;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 2s;
}
</style>
