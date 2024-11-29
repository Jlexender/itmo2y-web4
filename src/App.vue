<script setup>
import Splashscreen from "@/components/intro/Splashscreen.vue";
import { ref } from "vue";
import Disclaimer from "@/components/intro/Disclaimer.vue";
import Menu from "@/components/menu/Menu.vue";

const show = ref('Disclaimer');
</script>

<template>
  <div class="half debug">
    <transition name="fade" mode="in-out">
      <Disclaimer v-if="show === 'Disclaimer'" />
      <Splashscreen v-else-if="show === 'Splashscreen'" />
      <Menu v-else />
    </transition>
  </div>
  <button @click="show = 'Splashscreen'">Show Splashscreen</button>
  <button @click="show = 'Menu'">Show Menu</button>
  <button @click="show = 'Disclaimer'">Show Disclaimer</button>
</template>


<style scoped>
.half {
  width: 50vw;
  height: 50vh;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  position: relative;
}

.full {
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  position: relative;
}

button {
  margin-top: 10px;
}

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
