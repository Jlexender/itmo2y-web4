<script setup>
import {inject, ref} from 'vue';

const radius = inject('cRadius');

const data = ref([1, 2, 3, 4]);
const selectRadius = (event) => {

  const items = document.querySelectorAll('li');
  items.forEach(item => item.style.color = '');
  event.target.style.color = 'blue';
  radius.value = event.target.textContent;
}


</script>

<template>
  <div class="container">
    <ul>
      <li v-for="item in data" :key="item" @click="selectRadius">{{ item }}</li>
    </ul>
  </div>
</template>

<style scoped>
ul {
  position: absolute;
  list-style-type: none;
  right: 3%;
}

li {
  font-family: "Century Gothic", sans-serif;
  font-size: 72px;
  transition: color 0.5s ease;
}
</style>
