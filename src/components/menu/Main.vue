<script setup>
import MainNav from "@/components/menu/MainNav.vue";

defineEmits(['toExit']);
</script>

<template>
  <div class="container">
    <img src="@/assets/img/mainmenu_ground.jpg" alt="Main menu" class="game-canvas"/>

    <MainNav @toExit="$emit('toExit')"/>
  </div>
</template>

<style scoped>
img {
  width: 1920px;
  height: 1080px;
  object-fit: contain;
}
</style>
