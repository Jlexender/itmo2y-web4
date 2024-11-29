<script setup>
import { reactive, onMounted, onUnmounted } from 'vue';

const cursorPosition = reactive({ x: 0, y: 0 });

const updateCursorPosition = (event) => {
  cursorPosition.x = event.clientX;
  cursorPosition.y = event.clientY;
};

onMounted(() => {
  window.addEventListener('mousemove', updateCursorPosition);
});

onUnmounted(() => {
  window.removeEventListener('mousemove', updateCursorPosition);
});
</script>

<template>
  <div>
    <img
        src="@/assets/img/cursor.png"
        alt="Game cursor"
        class="game-element cursor"
        :style="{ left: `${cursorPosition.x}px`, top: `${cursorPosition.y}px` }"
    />
  </div>
</template>

<style scoped>
.cursor {
  position: absolute;
  width: 2rem;
  height: auto;
  z-index: 100;
  pointer-events: none;
}
</style>
