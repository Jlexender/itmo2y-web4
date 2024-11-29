<script setup>
import { ref } from 'vue';

const isHovered = ref(false);
const gateSound = new Audio(new URL('@/assets/audio/menu_gate.ogg', import.meta.url));

const handleMouseOver = () => {
  gateSound.play();
  isHovered.value = true;
};

const handleMouseLeave = () => {
  isHovered.value = false;
};
</script>

<template>
  <div
      class="clickable-rectangle"
      @mouseover="handleMouseOver"
      @mouseleave="handleMouseLeave">
  </div>
</template>

<style scoped>
.clickable-rectangle {
  position: sticky;
  width: 8%;
  height: 25%;
  top: 55%;
  left: 75%;
}
</style>
