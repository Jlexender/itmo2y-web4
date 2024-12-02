<script setup>
defineEmits(['toSplash']);
</script>

<template>
  <div @click="$emit('toSplash')">
    <img alt="Disclaimer" class="game-canvas" draggable="false" src="@/assets/img/disclaimer.jpg"/>
  </div>
</template>

<style scoped>
</style>
