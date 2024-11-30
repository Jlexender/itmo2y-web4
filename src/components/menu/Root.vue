<script setup>
import Main from "@/components/menu/Main.vue";
import {inject, onMounted, onUnmounted, provide, ref, watch} from "vue";
import Exit from "@/components/menu/Exit.vue";
import Options from "@/components/menu/Settings/Settings.vue";

const views = {
  main: Main,
  exit: Exit,
  options: Options,
};

const currentView = ref(null);
const setView = (view) => {
  currentView.value = view;
};

const volume = inject('volume');
const menuMusic = new Audio(new URL('@/assets/audio/blow_with_the_fires.ogg', import.meta.url));
watch(volume, (newValue) => {
  menuMusic.volume = newValue;
});


onMounted(() => {
  setView('main');
  menuMusic.loop = true;
  menuMusic.volume = volume.value;
  menuMusic.play();
});

onUnmounted(() => {
  setInterval(() => {
    if (menuMusic.volume > 0.01) {
      menuMusic.volume -= 0.01;
    } else {
      menuMusic.pause();
    }
  }, 30);
});
defineEmits(['toStart']);
</script>

<template>
  <div>

    <transition name="fade">
      <component :is="views[currentView]"
                 @toExit="setView('exit')"
                 @dontExit="setView('main')"
                 @toOptions="setView('options')"
                 @toMain="setView('main')"
                 @toStart="$emit('toStart')"
      />
    </transition>
  </div>
</template>

<style scoped>
</style>
