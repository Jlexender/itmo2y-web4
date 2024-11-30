<script setup>
import {onMounted, ref, watch} from 'vue';

defineProps({
  title: {
    type: String,
    required: true,
  },
});

const sliderValue = ref(100);

onMounted(() => {
  const savedValue = localStorage.getItem('sliderValue');
  if (savedValue !== null) {
    sliderValue.value = parseInt(savedValue, 10);
  }
});

watch(sliderValue, (newValue) => {
  localStorage.setItem('sliderValue', newValue);
});

defineEmits(['update']);
</script>

<template>
  <div class="slider-container">
    <label>{{ title }}</label>
    <div class="slider-track" @input="$emit('update', sliderValue)">
      <div
          :style="{ width: sliderValue + '%' }"
          class="slider-left"
      ></div>
      <div
          :style="{ width: 100 - sliderValue + '%', left: sliderValue + '%' }"
          class="slider-right"
      ></div>
      <input
          v-model="sliderValue"
          class="custom-slider"
          max="100"
          min="0"
          type="range"
      />
    </div>
  </div>
</template>

<style scoped>
.slider-container {
  display: flex;
  align-items: center;
  gap: 20px;
  width: 800px;
}

label {
  font-family: Roboto, sans-serif;
  font-size: 32px;
  width: 225px;
}

.slider-track {
  position: relative;
  width: 600px;
  height: 25px;
  border: 3px solid saddlebrown;
  background: transparent;
  overflow: hidden;
  border-radius: 0;
}

.slider-left {
  position: absolute;
  height: 100%;
  background: linear-gradient(180deg, #61af4c, #b7f388);
}

.slider-right {
  position: absolute;
  height: 100%;
  background: linear-gradient(180deg, #682600, #cf763f);
}

.custom-slider {
  appearance: none;
  width: 100%;
  height: 25px;
  background: transparent;
  position: absolute;
  top: 0;
  left: 0;
  z-index: 2;
  outline: none;
  opacity: 0;
  pointer-events: auto;
  cursor: none;
}

.custom-slider::-webkit-slider-thumb {
  appearance: none;
  width: 35px;
  height: 35px;
  background: linear-gradient(180deg, #61af4c, #9cda69);
  border: 3px inset #7a1600;
  border-radius: 0;
}

.custom-slider::-moz-range-thumb {
  width: 35px;
  height: 35px;
  background: linear-gradient(180deg, #61af4c, #9cda69);
  border: 3px inset #7a1600;
  border-radius: 0;
}
</style>
