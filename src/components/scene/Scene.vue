<script setup>

import {ref} from "vue";
import Disclaimer from "@/components/intro/Disclaimer.vue";
import Splash from "@/components/intro/Splash.vue";
import Menu from "@/components/menu/Menu.vue";

const views = [
  Disclaimer,
  Splash,
  Menu,
];

const viewId = ref(0);

const handleClick = () => {
  if (viewId.value < views.length - 1) {
    viewId.value++;
  }
}
</script>

<template>
  <div class="scene">
    <transition name="cross-fade">
      <component :is="views[viewId]" @click="handleClick"/>
    </transition>
  </div>
</template>

<style scoped>
.scene {
  position: fixed;
  top: 30px;
  left: 30px;
  width: 1920px;
  height: 1080px;
  overflow: hidden;
  outline: 3px solid red;
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

