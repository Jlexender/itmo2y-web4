<script setup>
import { inject, ref, onMounted } from 'vue';

const radius = inject('cRadius');

const value = ref(1);

const selectRadius = (event) => {
  radius.value = event.target.value;
};

onMounted(() => {
  radius.value = value.value;
});
</script>

<template>
  <div class="container">
    <input 
      type="range" 
      min="1" 
      max="4" 
      step="0.01"
      v-model="value" 
      @input="selectRadius" 
      class="slider" 
    />
    <div class="value-display">{{ value }}</div>
  </div>
</template>

<style scoped>
.container {
  position: relative;
}

.slider {
  position: absolute;
  top: 200px;
  right: -100px;
  writing-mode: bt-lr;
  width: 400px; 
  height: 100px;
  transform: rotate(-90deg);
  border-radius: 5px;
  cursor: none;
}

.value-display {
  position: absolute;
  top: 470px; 
  right: 78px;
  width: 50px;
  font-weight: bold;
  font-family: "Century Gothic", sans-serif;
  font-size: 24px;
  text-align: center;
  color: white;
}
</style>

