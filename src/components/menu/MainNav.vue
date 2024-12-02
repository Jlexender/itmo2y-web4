<script setup>

import {inject} from "vue";

const volume = inject('volume');

const gateSound = new Audio(new URL('@/assets/audio/menu_gate.ogg', import.meta.url));
gateSound.volume = volume.value;

const playGateSound = () => {
  gateSound.play();
};

const emit = defineEmits([
  'toExit',
  'toOptions',
  'toStart',
  'toLogin',
  'openGates',
  'closeGates',
  'toTable']
);

const handleExitEnter = () => {
  playGateSound();
  emit('openGates');
};

</script>

<template>
  <div
      class="transparent exit-area"
      @click="$emit('toExit')"
      @mouseenter="handleExitEnter"
      @mouseleave="$emit('closeGates')">
    Exit
  </div>

  <div class="transparent start-area" @click="$emit('toStart')">
    Start
  </div>

  <div class="transparent options-area" @click="$emit('toOptions')">
    Options
  </div>

  <div class="transparent authorize-area" @click="$emit('toLogin')">
    Authorization
  </div>

  <div class="transparent table-area" @click="$emit('toTable')">
    Table
  </div>

</template>

<style scoped>
.transparent {
  opacity: 0;
}

.exit-area {
  position: absolute;
  bottom: 24%;
  right: 16%;
  width: 150px;
  height: 280px;
  background-color: white;
  transform: skewY(-8deg);
}

.start-area {
  position: absolute;
  bottom: 22%;
  left: 23.5%;
  width: 320px;
  height: 580px;
  background-color: white;
}

.options-area {
  position: absolute;
  bottom: 3%;
  right: 30%;
  width: 250px;
  height: 250px;
  transform: skewY(-8deg);
  background-color: white;
}

.authorize-area {
  position: absolute;
  bottom: 25%;
  left: 41.5%;
  width: 270px;
  height: 550px;
  transform: skewY(-8deg);
  background-color: white;
}

.table-area {
  position: absolute;
  bottom: 29%;
  left: 56.5%;
  width: 230px;
  height: 505px;
  transform: skewY(-8deg);
  background-color: white;
}

.start-area:hover, .authorize-area:hover, .table-area:hover {
  opacity: 0.2;
  background-color: #ffffff;
}

</style>
