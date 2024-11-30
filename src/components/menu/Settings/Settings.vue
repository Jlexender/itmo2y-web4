<script setup>
import Bar from "@/components/menu/Settings/Bar.vue";
import {inject} from "vue";

const volume = inject('volume');

const handle = (value) => {
  volume.value = value / 100;
};

defineEmits(['toMain']);
</script>

<template>
  <div>
    <img src="@/assets/img/history_bg.jpg" alt="Options" class="game-canvas" draggable="false"/>
    <img src="@/assets/img/preferences_bg.jpg" alt="Options 2" class="game-canvas" draggable="false"/>

    <header>
      <span class="star">☆</span> Настройки <span class="star">☆</span>
    </header>

    <div class="settings">
      Громкость
      <ul>
        <li>
          <Bar title="Общая" @update="handle"/>
        </li>
      </ul>
    </div>

    <div class="go-back" @click="$emit('toMain')">
      ← Назад
    </div>
  </div>
</template>

<style scoped>
header {
  position: absolute;
  font-size: 58px;
  text-transform: uppercase;
  font-family: "Century Gothic", sans-serif;
  top: 10%;
  left: 50%;
  transform: translateX(-50%);
}

.settings {
  text-align: center;
  position: absolute;
  top: 30%;
  left: 50%;
  transform: translateX(-50%);
  font-size: 48px;
  font-family: Roboto, sans-serif;
  color: #a35434;
}

.star {
  color: gray;
  font-weight: bold;
}

.go-back {
  position: absolute;
  bottom: 10%;
  left: 50px;
  font-size: 58px;
  font-family: "Century Gothic", sans-serif;
  cursor: none;
  color: gray;
}

.go-back:hover {
  color: white;
}

ul {
  list-style-type: none;
  padding: 0;
  margin: 0;
}

ul li {
  margin-top: 20px;
  display: flex;
  justify-content: space-between;
  font-size: 36px;
  font-family: Roboto, sans-serif;
  color: #a35434;
}

.slider {
  margin-left: 20px;
  width: 500px;
}
</style>
