<script setup>
import MainNav from "@/components/menu/MainNav.vue";
import {ref} from "vue";

const showGate = ref(false);

defineEmits(['toExit', 'toOptions', 'toStart', 'toLogin', 'toTable']);
</script>

<template>
  <div class="container">
    <img alt="Main menu" class="game-canvas" draggable="false" src="@/assets/img/mainmenu_ground.jpg"/>
    <transition appear name="fade">
      <img v-if="showGate" alt="Main menu" class="game-canvas" draggable="false" src="@/assets/img/mainmenu_hover.jpg"/>
    </transition>


    <MainNav @closeGates="showGate = false"
             @openGates="showGate = true"
             @toLogin="$emit('toLogin')"
             @toExit="$emit('toExit')"
             @toOptions="$emit('toOptions')"
             @toStart="$emit('toStart')"
             @toTable="$emit('toTable')"
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
