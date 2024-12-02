<script setup>
import {onMounted, provide, ref} from "vue";
import Disclaimer from "@/components/intro/Disclaimer.vue";
import Splash from "@/components/intro/Splash.vue";
import Menu from "@/components/menu/Root.vue";
import Cursor from "@/components/Cursor.vue";
import Game from "@/components/game/Game.vue";
import Unauthorized from "@/components/menu/Unauthorized.vue";
import Author from "@/components/intro/Author.vue";

const mode = ref('default');
const name = ref('cross-fade');
const views = {
  'author': Author,
  'disclaimer': Disclaimer,
  'splash': Splash,
  'menu': Menu,
  'game': Game,
}
const viewId = ref(0);

const volume = localStorage.getItem('sliderValue') ?
    ref(parseInt(localStorage.getItem('sliderValue'), 10) / 100) : ref(1);
provide('volume', volume);

const setScene = (name) => {
  viewId.value = Object.keys(views).findIndex((key) => key === name);
};

const hasAuthenticated = ref(false);
provide('auth', hasAuthenticated);
const showWarn = ref(false);
const handleStart = () => {
  if (hasAuthenticated.value === true) {
    startGame();
  } else {
    showWarn.value = true;
  }
}

const startGame = () => {
  mode.value = 'out-in';
  name.value = 'long-fade';
  setScene('game');
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
  <Cursor/>
  <div class="scene">
    <transition :mode="mode" :name="name">
      <component :is="views[Object.keys(views)[viewId]]"
                 @toDisclaimer="setScene('disclaimer')"
                 @toMain="setScene('menu')"
                 @toSplash="setScene('splash')"
                 @toStart="handleStart"
      />
    </transition>

    <transition name="short-fade">
      <Unauthorized v-if="showWarn" @click="showWarn=false"/>
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

.long-fade-enter-active,
.long-fade-leave-active {
  transition: opacity 2s ease-in-out;
}

.long-fade-enter-from,
.long-fade-leave-to {
  opacity: 0;
}

.long-fade-enter-to,
.long-fade-leave-from {
  opacity: 1;
}

.short-fade-enter-active,
.short-fade-leave-active {
  transition: opacity 0.5s ease-in-out;
}

.short-fade-enter-from,
.short-fade-leave-to {
  opacity: 0;
}

.short-fade-enter-to,
.short-fade-leave-from {
  opacity: 1;
}
</style>
