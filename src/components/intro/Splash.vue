<script setup>

import Logo from "@/components/intro/Logo.vue";
</script>

<template>
  <div>
    <img src="@/assets/img/splashscreen_day.png" alt="Splashscreen" class="move-up-half" />
    <Logo />
  </div>

</template>

<style scoped>
@keyframes moveUpHalf {
  0% {
    transform: translateY(0);
  }
  100% {
    transform: translateY(-50%);
  }
}

.move-up-half {
  animation: moveUpHalf 5s ease-out forwards;
}

img {
  position: absolute;
  width: 100%;
  height: 200%;
  object-fit: contain;
}

</style>
