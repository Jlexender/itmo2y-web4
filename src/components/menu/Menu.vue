<script setup>
import Gate from "@/components/menu/Gate.vue";
</script>

<template>
  <div class="container">

    <img src="@/assets/img/mainmenu_ground.jpg" alt="Game menu" class="game-element"/>
    <Gate />

    <audio autoplay loop>
      <source src="@/assets/audio/blow_with_the_fires.ogg" type="audio/ogg">
    </audio>
  </div>
</template>

<style scoped>
img {
  position: absolute;
  width: 100%;
  height: 100%;
  object-fit: contain;
}
</style>
