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
  />
</template>

<style scoped>
.clickable-rectangle {
  position: absolute;
  left: 76.5%;
  top: 50%;

  width: 8rem;
  height: 17rem;
  z-index: 2;

  transform: skewY(-10deg);

}

</style>
