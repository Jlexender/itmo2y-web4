<script setup>
import Main from "@/components/menu/Main.vue";
import {onMounted, ref} from "vue";
import Exit from "@/components/menu/Exit.vue";
import Options from "@/components/menu/Options.vue";

const views = {
  main: Main,
  exit: Exit,
  options: Options,
};

const currentView = ref(null);

const setView = (view) => {
  currentView.value = view;
};

onMounted(() => {
  setView('main');
});
</script>

<template>
  <div>
    <audio autoplay loop>
      <source src="@/assets/audio/blow_with_the_fires.ogg" type="audio/ogg"/>
    </audio>

    <transition name="fade">
      <component :is="views[currentView]"
                 @toExit="setView('exit')"
                 @dontExit="setView('main')"
                 @toOptions="setView('options')"
      />
    </transition>
  </div>
</template>

<style scoped>
</style>
