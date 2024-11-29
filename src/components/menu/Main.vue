<script setup>
import MainNav from "@/components/menu/MainNav.vue";
</script>

<template>
  <div>
    <img src="@/assets/img/mainmenu_ground.jpg" alt="Main menu" />
    <MainNav />
  </div>
</template>

<style scoped>
img {
  position: absolute;
  width: 100vw;
  height: 100vh;
  object-fit: contain;
}
</style>
