<script setup>
import Main from "@/components/menu/Main.vue";
import {ref, shallowRef} from "vue";
import Exit from "@/components/menu/Exit.vue";

const views = [
  Main,
  Exit,
];

const viewId = ref(0);

const setViewId = (id) => {
  viewId.value = id;
}

</script>

<template>
  <div>
    <audio autoplay loop>
      <source src="@/assets/audio/blow_with_the_fires.ogg" type="audio/ogg"/>
    </audio>

    <transition name="fade">
      <component :is="views[viewId]" @toExit="setViewId(1)" @dontExit="setViewId(0)"/>
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

</style>
