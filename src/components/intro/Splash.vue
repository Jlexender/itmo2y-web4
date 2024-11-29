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
</script>

<template>
  <div>
    <img src="@/assets/img/splashscreen_day.png" alt="Splashscreen" class="game-canvas"/>
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
  animation: moveUp 6s ease-out forwards;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

.fade-enter-to, .fade-leave-from {
  opacity: 1;
}
</style>
