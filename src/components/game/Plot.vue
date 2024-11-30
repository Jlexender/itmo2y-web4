<script setup>
import {inject, ref, watch} from 'vue';

const canvasRef = ref(null);
const cRadius = inject('cRadius');

const emit = defineEmits(['makeSad', 'makeHappy']);
watch(cRadius, drawCanvas);

function drawDot(ctx, x, y, color) {
  ctx.fillStyle = color;
  ctx.beginPath();
  ctx.arc(x, y, 5, 0, Math.PI * 2);
  ctx.fill();
}

function drawRect(ctx, x, y, width, height, color) {
  ctx.fillStyle = color;
  ctx.fillRect(x, y, width, height);
}

function drawArc(ctx, x, y, radius, startAngle, endAngle, color) {
  ctx.fillStyle = color;
  ctx.beginPath();
  ctx.moveTo(x, y);
  ctx.arc(x, y, radius, startAngle, endAngle);
  ctx.closePath();
  ctx.fill();
}

function drawAxes(ctx) {
  const width = ctx.canvas.width;
  const height = ctx.canvas.height;

  drawLine(ctx, 0, height / 2, width, height / 2, 'white');
  drawLine(ctx, width / 2, 0, width / 2, height, 'white');

  for (let i = 0; i < width; i += 50) {
    drawLine(ctx, i, height / 2 - 5, i, height / 2 + 5, 'white');
  }

  for (let i = 0; i < height; i += 50) {
    drawLine(ctx, width / 2 - 5, i, width / 2 + 5, i, 'white');
  }

  ctx.font = '20px Arial';
  ctx.fillStyle = 'blue';

  ctx.fillText('x', width - 20, height / 2 - 20);
  ctx.fillText('y', width / 2 + 20, 20);
  ctx.fillText('0', width / 2 + 20, height / 2 - 20);
}

function drawLine(ctx, x1, y1, x2, y2, color) {
  ctx.strokeStyle = color;
  ctx.lineWidth = 2;
  ctx.beginPath();
  ctx.moveTo(x1, y1);
  ctx.lineTo(x2, y2);
  ctx.stroke();
}

function drawTriangle(ctx, x1, y1, x2, y2, x3, y3, color) {
  ctx.fillStyle = color;
  ctx.beginPath();
  ctx.moveTo(x1, y1);
  ctx.lineTo(x2, y2);
  ctx.lineTo(x3, y3);
  ctx.closePath();
  ctx.fill();
}

function drawCanvas() {
  const canvas = canvasRef.value;
  const ctx = canvas.getContext('2d');
  const drawColor = 'rgba(67,109,189,0.7)';
  const centerX = canvas.width / 2;
  const centerY = canvas.height / 2;

  ctx.clearRect(0, 0, canvas.width, canvas.height);

  drawAxes(ctx);
  drawArc(ctx, centerX, centerY, 50 * cRadius.value, 0, Math.PI / 2, drawColor);
  drawRect(ctx, centerX - 100 * cRadius.value, centerY - 100 * cRadius.value, 100 * cRadius.value, 100 * cRadius.value, drawColor);
  drawTriangle(ctx, centerX, centerY, centerX - 100 * cRadius.value, centerY, centerX, centerY + 50 * cRadius.value, drawColor);
}

const handleClick = (event) => {
  const canvas = canvasRef.value;
  const ctx = canvas.getContext('2d');
  const rect = canvas.getBoundingClientRect();

  const x = (event.clientX - rect.left) * canvas.width / rect.width;
  const y = (event.clientY - rect.top) * canvas.height / rect.height;

  drawDot(ctx, x, y, 'white');
  const result = checkIfInArea(x, y);
  if (result) {
    emit('makeHappy');
  } else {
    emit('makeSad');
  }
};

const checkIfInArea = (x, y) => {
  const centerX = canvasRef.value.width / 2;
  const centerY = canvasRef.value.height / 2;
  const radius = cRadius.value * 50;

  if (x > centerX && y < centerY) {
    return false;
  } else if (x > centerX && y > centerY) {
    return (x - centerX) ** 2 + (y - centerY) ** 2 <= radius**2;
  } else if (x < centerX && y > centerY) {
    return -0.5 * (x - centerX) - radius <= centerY - y;
  } else {
    return 2*x >= centerX - radius && 2*y >= centerY - radius;
  }
};
</script>

<template>
  <canvas ref="canvasRef" width="800" height="800" @click="handleClick"></canvas>
</template>

<style scoped>
canvas {
  position: absolute;
  right: 5%;
  top: 5%;
  border-radius: 20%;
  transform: skewY(10deg);
}
</style>
