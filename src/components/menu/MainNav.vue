<script setup>

const gateSound = new Audio(new URL('@/assets/audio/menu_gate.ogg', import.meta.url));

const playGateSound = () => {
  gateSound.play();
};

const emit = defineEmits(['toExit', 'openGates', 'closeGates', 'toOptions', 'toStart']);

const handleExitEnter = () => {
  playGateSound();
  emit('openGates');
};

</script>

<template>
  <div
      class="transparent exit-area"
      @mouseenter="handleExitEnter"
      @mouseleave="$emit('closeGates')"
      @click="$emit('toExit')">
    Exit
  </div>

  <div class="transparent start-area" @click="$emit('toStart')">
    Start
  </div>

  <div class="transparent options-area" @click="$emit('toOptions')">
    Options
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
  bottom: 20%;
  left: 23%;
  width: 330px;
  height: 615px;
  transform: skewY(-4deg);
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

.start-area:hover{
  opacity: 0.2;
  background-color: #ffffff;
}

</style>
