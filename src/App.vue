<script setup>
import Splashscreen from "@/components/intro/Splashscreen.vue";
import { ref } from "vue";
import Disclaimer from "@/components/intro/Disclaimer.vue";
import Menu from "@/components/menu/Base.vue";
import Gate from "@/components/menu/Gate.vue";
import Cursor from "@/components/Cursor.vue";

const view = ref('Disclaimer');
</script>

<template>
  <div class="game-screen">
    <div class="game">
      <Cursor/>
      <transition name="fade" mode="in-out">
        <Disclaimer v-if="view === 'Disclaimer'" />
        <Splashscreen v-else-if="view === 'Splashscreen'" />
        <Menu v-else />
      </transition>
    </div>
  </div>

  <button @click="view = 'Splashscreen'">Show Splashscreen</button>
  <button @click="view = 'Menu'">Show Menu</button>
  <button @click="view = 'Disclaimer'">Show Disclaimer</button>
</template>


<style scoped>
.game {
  width: 1920px;
  height: 1080px;
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  position: relative;
  padding: 0;
}

.user-screen {
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #000;
}

button {
  margin-top: 10px;
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
