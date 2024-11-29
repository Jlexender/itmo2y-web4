<script setup>

const gateSound = new Audio(new URL('@/assets/audio/menu_gate.ogg', import.meta.url));

const playGateSound = () => {
  gateSound.play();
};

defineEmits(['toExit']);
</script>

<template>
  <div class="transparent exit-area" @mouseenter="playGateSound" @click="$emit('toExit')">
    Exit
  </div>

  <div class="transparent start-area">
    Start
  </div>

</template>

<style scoped>
.transparent {
  opacity: 0.5;
}

.exit-area {
  position: absolute;
  bottom: 24%;
  right: 16%;
  width: 150px;
  height: 280px;
  background-color: red;
  transform: skewY(-8deg);
}

.start-area {
  position: absolute;
  bottom: 20%;
  left: 23%;
  width: 330px;
  height: 615px;
  transform: skewY(-9deg);
  background-color: green;
}

</style>
