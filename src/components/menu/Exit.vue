<script setup>
import {inject, ref} from "vue";

const changeView = inject('changeView');

const showYes = ref(true)
const exit = ref('Вы действительно хотите выйти из моей лабы?')

const handleYes = () => {
  exit.value = 'Эй! Не нужно уходить из моей лабы!'
  showYes.value = false
}
</script>

<template>
  <div>
    <img src="@/assets/img/exit.jpg" alt="Game exit" class="game-element"/>

    <div class="text-block">
      {{ exit }}
    </div>

    <div class="choice">
      <ul>
        <li @click="handleYes" :style="{ visibility: showYes ? 'visible' : 'hidden' }">Да</li>
        <li @click="changeView('Main')">Нет</li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
@font-face {
  font-family: 'Century Gothic';
  src: url('@/assets/fonts/CenturyGothic.ttf') format("truetype");
  font-weight: normal;
  font-style: normal;
}

img {
  position: absolute;
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.text-block {
  position: absolute;
  width: 50%;
  top: 25%;
  left: 40%;
  font-size: 2vw;
  color: #2727b6;
  text-align: center;
  font-family: 'Century Gothic', sans-serif;
}

.choice {
  position: absolute;
  width: 50%;
  top: 50%;
  left: 40%;
  font-size: 1.5vw;
  font-family: 'Century Gothic', sans-serif;
}

ul {
  list-style-type: none;
  display: flex;
  justify-content: center;
  padding: 0;
  gap: 10vw;
  margin: 0;
}

li {
  display: inline;
  font-size: 1.5em;
  color: #2727b6;
}

li:hover {
  cursor: pointer;
  color: #ea7d00;
}
</style>
