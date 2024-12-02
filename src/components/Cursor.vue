<script setup>
import {onMounted, ref} from "vue";

const show = ref(false);
const event = ref({clientX: 0, clientY: 0});

onMounted(() => {
  window.addEventListener("mousemove", (e) => {
    show.value = true;
    event.value = {clientX: e.clientX, clientY: e.clientY};
  });
});
</script>

<template>
  <img v-if="show" :style="{ top: `${event.clientY}px`, left: `${event.clientX}px` }" alt="Cursor"
       class="cursor"
       src="@/assets/img/cursor.png"
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
