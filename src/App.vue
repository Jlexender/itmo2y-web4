<script setup>
import Splashscreen from "@/components/intro/Splashscreen.vue";
import {onMounted, ref} from "vue";
import Disclaimer from "@/components/intro/Disclaimer.vue";
import Menu from "@/components/menu/Base.vue";

const view = ref('Disclaimer');
</script>

<template>
  <div class="game">
    <transition name="fade" mode="in-out">
      <Disclaimer v-if="view === 'Disclaimer'" />
      <Splashscreen v-else-if="view === 'Splashscreen'" />
      <Menu v-else />
    </transition>
  </div>

  <button @click="view = 'Splashscreen'">Splashscreen</button>
  <button @click="view = 'Menu'">Menu</button>
  <button @click="view = 'Disclaimer'">Disclaimer</button>
</template>


<style scoped>
.game {
  aspect-ratio: 16 / 9;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  position: relative;
  background-color: #000;
  outline: red solid 2px;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 1s;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

.fade-enter-to, .fade-leave-from {
  opacity: 1;
}

</style>
