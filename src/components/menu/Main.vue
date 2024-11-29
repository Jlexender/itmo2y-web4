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

    <transition name="fade" mode="in-out">
      <img v-if="!isExitHovered" src="@/assets/img/mainmenu_ground.jpg" alt="Game menu" class="game-element"/>
      <img v-else src="@/assets/img/mainmenu_hover.jpg" alt="Game exit" class="game-element"/>
    </transition>
  </div>
</template>

<style scoped>
img {
  position: absolute;
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

.fade-enter-to, .fade-leave-from {
  opacity: 1;
}
</style>
