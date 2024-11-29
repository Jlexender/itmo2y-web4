<script setup>
import { ref } from 'vue'
import Main from "@/components/menu/Main.vue";

const page = Main
</script>

<template>
  <component :is="page" />
</template>

<style scoped>

</style>
