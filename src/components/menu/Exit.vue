<script setup>
import {ref} from "vue";

const exit = ref('Вы действительно хотите выйти из моей лабы?')
const showYes = ref(true)

const handleYes = () => {
  exit.value = 'Эй! Не нужно выходить из моей лабы!';
  showYes.value = false;
}

defineEmits(['dontExit'])
</script>

<template>
  <div>
    <img src="@/assets/img/exit.jpg" alt="Game exit" class="game-canvas" draggable="false"/>

    <div class="text-block">
      {{ exit }}
    </div>

    <div class="choice">
      <ul>
        <li @click="handleYes" :style="{ visibility: showYes ? 'visible' : 'hidden' }">Да</li>
        <li @click="$emit('dontExit')">Нет</li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
.text-block {
  position: absolute;
  top: 28%;
  left: 40%;
  width: 900px;
  color: #2727b6;
  text-align: center;
  font-size: 60px;
  font-family: 'Century Gothic', sans-serif;
}

.choice {
  position: absolute;
  width: 50%;
  top: 50%;
  left: 40%;
  font-family: 'Century Gothic', sans-serif;
}

ul {
  list-style-type: none;
  display: flex;
  justify-content: center;
  padding: 0;
  margin: 0;
  gap: 40%;
}

li {
  display: inline;
  font-size: 64px;
  color: #397bf1;
}

li:hover {
  color: #ea7d00;
}
</style>
