<script setup>
import MainNav from "@/components/menu/MainNav.vue";
import {ref} from "vue";

const showGate = ref(false);

defineEmits(['toExit', 'toOptions']);
</script>

<template>
  <div class="container">
    <img src="@/assets/img/mainmenu_ground.jpg" alt="Main menu" class="game-canvas"/>
    <transition name="fade" appear>
      <img v-if="showGate" src="@/assets/img/mainmenu_hover.jpg" alt="Main menu" class="game-canvas"/>
    </transition>


    <MainNav @toExit="$emit('toExit')"
             @toOptions="$emit('toOptions')"
             @openGates="showGate = true"
             @closeGates="showGate = false"
    />
  </div>
</template>

<style scoped>
img {
  width: 1920px;
  height: 1080px;
  object-fit: contain;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s;
}
</style>
