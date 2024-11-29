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
    transform: translateY(-50%);
  }
}

img {
  width: 1920px;
  height: 2160px;
  position: absolute;
  animation: moveUp 6s ease-out forwards;
}
</style>
