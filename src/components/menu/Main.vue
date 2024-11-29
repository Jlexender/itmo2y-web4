<script setup>
import Gate from "@/components/menu/Gate.vue";
import { ref } from "vue";

const isExitHovered = ref(false);

const handleHoverUpdate = (hovered) => {
  isExitHovered.value = hovered;
};
</script>

<template>
  <div>
    <Gate @update:hovered="handleHoverUpdate" />

    <img v-if="!isExitHovered" src="@/assets/img/mainmenu_ground.jpg" alt="Game menu" class="game-element"/>
    <img v-else src="@/assets/img/mainmenu_hover.jpg" alt="Game exit" class="game-element"/>
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
