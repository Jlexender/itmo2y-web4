<script setup>
import { ref, onMounted } from "vue";
import Disclaimer from "@/components/intro/Disclaimer.vue";
import Splash from "@/components/intro/Splash.vue";
import Menu from "@/components/menu/Root.vue";
import Exit from "@/components/menu/Exit.vue";

const views = [Disclaimer, Splash, Menu];
const viewId = ref(0);

const handleClick = () => {
  if (viewId.value < views.length - 1) {
    viewId.value++;
  }
};

const scaleScene = () => {
  const scene = document.querySelector(".scene");
  if (!scene) return;

  const sceneWidth = 1920;
  const sceneHeight = 1080;

  const windowWidth = window.innerWidth;
  const windowHeight = window.innerHeight;

  const scale = Math.min(windowWidth / sceneWidth, windowHeight / sceneHeight);

  scene.style.transform = `scale(${scale})`;
  scene.style.left = `${(windowWidth - sceneWidth * scale) / 2}px`;
  scene.style.top = `${(windowHeight - sceneHeight * scale) / 2}px`;
};

onMounted(() => {
  scaleScene();
  window.addEventListener("resize", scaleScene);
});

</script>

<template>
  <div class="scene">
    <transition name="cross-fade">
      <component :is="views[viewId]" @click="handleClick" @toMain="handleClick"/>
    </transition>
  </div>
</template>

<style scoped>
.scene {
  position: fixed;
  width: 1920px;
  height: 1080px;
  transform-origin: top left;
  overflow: hidden;
}

.cross-fade-enter-active,
.cross-fade-leave-active {
  transition: opacity 0.5s ease-in-out;
}

.cross-fade-enter-from,
.cross-fade-leave-to {
  opacity: 0;
}

.cross-fade-enter-to,
.cross-fade-leave-from {
  opacity: 1;
}
</style>
