<script setup>

import Logo from "@/components/intro/Logo.vue";
</script>

<!-- FIXME: FIX SPLASHSCREEN ANIMATION -->
<template>
  <div>
    <img src="@/assets/img/splashscreen_day.png" alt="Splashscreen"/>
    <Logo/>
  </div>

</template>

<style scoped>
@keyframes moveUp {
  0% {
    transform: translateY(0);
  }
  100% {
    transform: translateY(-100vh);
  }
}

img {
  position: absolute;
  width: 100vw;
  height: 200vh;
  object-fit: contain;
  animation: moveUp 6s ease-out forwards;
}
</style>
