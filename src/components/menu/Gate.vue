<script setup>
import {ref, defineEmits, inject} from 'vue';

const isExitHovered = ref(false);
const gateSound = new Audio(new URL('@/assets/audio/menu_gate.ogg', import.meta.url));
const emit = defineEmits(['update:hovered']);

const changeView = inject('changeView');

const handleMouseOver = () => {
  gateSound.play();
  isExitHovered.value = true;
  emit('update:hovered', true);
};

const handleMouseLeave = () => {
  isExitHovered.value = false;
  emit('update:hovered', false);
};
</script>

<template>
  <div
      class="clickable-rectangle game-element"
      @mouseover="handleMouseOver"
      @mouseleave="handleMouseLeave"
      @click="changeView('Exit')"
  >
    Gate
  </div>
</template>

<style scoped>
.clickable-rectangle {
  background-color: #b9ffae;
  position: absolute;
  width: 100px;
  height: 100px;
  z-index: 2;
}
</style>
