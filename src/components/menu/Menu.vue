<script setup>
import Main from "@/components/menu/Main.vue";

const page = Main
</script>

<template>
  <div>
    <audio autoplay loop>
      <source src="@/assets/audio/blow_with_the_fires.ogg" type="audio/ogg"/>
    </audio>

    <component :is="page"/>
  </div>
</template>

<style scoped>

</style>
