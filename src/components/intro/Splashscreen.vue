<script setup>
import {ref} from "vue";
import Logo from "@/components/intro/Logo.vue";

const showLogo = ref(false);

setTimeout(() => {
  showLogo.value = true;
}, 4000);
</script>

<template>
  <div class="container">
    <img src="../../assets/img/splashscreen_night.png" alt="Game splashscreen" class="game-element move-up"/>

    <transition name="fade">
      <Logo v-if="showLogo"/>
    </transition>
  </div>

</template>

<style scoped>
@keyframes moveUp {
  0% {
    transform: translateY(-5%);
  }
  100% {
    transform: translateY(-45%);
  }
}

.move-up {
  animation: moveUp 6s forwards;
}

img {
  position: absolute;
  width: 100%;
  height: 200%;
  object-fit: contain;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 1s;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

.fade-enter-to, .fade-leave-from {
  opacity: 1;
}
</style>


