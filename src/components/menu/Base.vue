<script setup>
import Main from "@/components/menu/Main.vue";
import {provide, ref} from "vue";
import Exit from "@/components/menu/Exit.vue";

const currentView = ref('Main');

const changeView = (value) => {
  currentView.value = value;
};

provide('changeView', changeView);
</script>

<template>
  <div class="container">
    <audio autoplay loop>
      <source src="@/assets/audio/blow_with_the_fires.ogg" type="audio/ogg">
    </audio>

    <transition name="fade">
      <Main v-if="currentView === 'Main'" />
      <Exit v-else />
    </transition>

  </div>
</template>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

.fade-enter-to, .fade-leave-from {
  opacity: 1;
}
</style>
