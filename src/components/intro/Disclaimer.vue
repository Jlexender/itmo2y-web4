<script setup>
defineEmits(['toSplash']);
</script>

<template>
  <div @click="$emit('toSplash')">
    <img src="@/assets/img/disclaimer.jpg" alt="Disclaimer" class="game-canvas" draggable="false"/>
  </div>
</template>

<style scoped>
</style>
