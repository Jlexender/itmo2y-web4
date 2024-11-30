<script setup>
import Main from "@/components/menu/Main.vue";
import {inject, onMounted, onUnmounted, ref, watch} from "vue";
import Exit from "@/components/menu/Exit.vue";
import Options from "@/components/menu/settings/Settings.vue";
import Login from "@/components/menu/auth/Login.vue";
import Register from "@/components/menu/auth/Register.vue";

const views = {
  main: Main,
  exit: Exit,
  options: Options,
  login: Login,
  register: Register
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
  setTimeout(() => {
    menuMusic.loop = true;
    menuMusic.volume = volume.value;
    menuMusic.play();
  }, 500);
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
                 @dontExit="setView('main')"
                 @toExit="setView('exit')"
                 @toLogin="setView('login')"
                 @toMain="setView('main')"
                 @toOptions="setView('options')"
                 @toRegister="setView('register')"
                 @toStart="$emit('toStart')"
      />
    </transition>
  </div>
</template>

<style scoped>
</style>
