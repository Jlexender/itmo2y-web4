<script setup>
import { ref, onMounted } from "vue";

const show = ref(false);
const event = ref({ clientX: 0, clientY: 0 });

onMounted(() => {
  window.addEventListener("mousemove", (e) => {
    show.value = true;
    event.value = { clientX: e.clientX, clientY: e.clientY };
  });
});
</script>

<template>
  <img src="@/assets/img/cursor.png" alt="Cursor" class="cursor"
       :style="{ top: `${event.clientY}px`, left: `${event.clientX}px` }"
       v-if="show"
  />
</template>

<style scoped>
.cursor {
  position: fixed;
  width: 1.5vw;
  height: auto;
  pointer-events: none;
  z-index: 1000;
}
</style>
