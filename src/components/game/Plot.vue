<script setup>
import {inject, onMounted, ref, watch} from "vue";

const radius = inject('cRadius');
const limit = 5;
watch(radius, () => draw());

const emit = defineEmits(['makeSad', 'makeHappy']);

const canvasRef = ref(null);
const userDots = ref([]);

const drawPolyline = (ctx, polyline) => {
  ctx.lineWidth = 3;
  ctx.fillStyle = 'rgba(178,53,255,0.3)';
  ctx.beginPath();
  ctx.moveTo(polyline[0].x, polyline[0].y);
  for (let i = 1; i < polyline.length; i++) {
    ctx.lineTo(polyline[i].x, polyline[i].y);
  }
  ctx.fill();
};

const drawArc = (ctx, x, y, radius, startAngle, endAngle) => {
  ctx.fillStyle = 'rgba(178,53,255,0.3)';
  ctx.beginPath();
  ctx.arc(x, y, radius, startAngle, endAngle)
  ctx.lineTo(x, y);
  ctx.fill();
};

const drawAxis = (ctx) => {
  ctx.lineWidth = 3;
  ctx.strokeStyle = 'white';

  // Y-axis
  ctx.beginPath();
  ctx.moveTo(400, 0);
  ctx.lineTo(400, 800);
  ctx.stroke();

  // X-axis
  ctx.beginPath();
  ctx.moveTo(0, 400);
  ctx.lineTo(800, 400);
  ctx.stroke();
};

const labelAxis = (ctx) => {
  ctx.fillStyle = 'blue';
  ctx.font = '20px Century Gothic';
  ctx.fillText('Y', 400 + 10, 20);
  ctx.fillText('X', 780, 400 - 10);

  // x-axis labels
  for (let i = 1; i < limit; i++) {
    ctx.fillText(i, 400 + i*100 - 10, 400 + 20);
  }

  // negative x-axis labels
  for (let i = 1; i < limit; i++) {
    ctx.fillText(-i, 400 - i*100 - 10, 400 + 20);
  }

  // y-axis labels
  for (let i = 1; i < limit; i++) {
    ctx.fillText(i, 400 - 20, 400 - i*100 + 10);
  }

  // negative y-axis labels
  for (let i = 1; i < limit; i++) {
    ctx.fillText(-i, 400 - 20, 400 + i*100 + 10);
  }
};

const drawFigure = (ctx, radius) => {
  const x = 400;
  const y = 400;

  const unit = 100*radius;

  const polyline = [
    { x: x - unit, y: y - unit },
    { x: x - unit, y: y },
    { x: x, y: y + unit/2 },
    { x: x, y: y - unit },
  ];

  drawPolyline(ctx, polyline);

  const startAngle = 0;
  const endAngle = Math.PI/2;
  drawArc(ctx, x, y, unit/2, startAngle, endAngle);
};

const checkIfInsideFigure = (x, y, radius) => {
  const unit = 100*radius;

  if (x > 400 && y < 400) {
    return false;
  } else if (x <= 400 && y >= 400) {
    return 2*(y - 400) <= (x - 400) + unit;
  } else if (x < 400 && y < 400) {
    return x > 400 - unit && y > 400 - unit;
  } else {
    return (400 - x)**2 + (400 - y)**2 < (unit/2)**2;
  }
};

const backendIfInsideFigure = async (x, y, radius) => {
  try {
    const response = await fetch('http://localhost:3080/api/check', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ x, y, radius }),
    });
    const data = await response.json();
    return data.result;
  } catch (error) {
    console.error('Error:', error);
    return false;
  }
};

const handleClick = (event) => {
  const canvas = canvasRef.value;
  const ctx = canvas.getContext('2d');
  const x = event.offsetX;
  const y = event.offsetY;

  const isInside = checkIfInsideFigure(x, y, radius.value);
  const dot = { x, y, isInside };
  userDots.value.push(dot);

  if (isInside) {
    emit('makeHappy');
  } else {
    emit('makeSad');
  }

  drawDot(ctx, x, y, isInside);
  drawPointCoords(ctx, x, y);

  // drawCanvasCoords(ctx, x, y);
};

const drawCanvasCoords = (ctx, x, y) => {
  ctx.fillStyle = 'blue';
  ctx.font = '16px Century Gothic';
  ctx.fillText(`(${x}, ${y})`, x + 10, y - 10);
};

const drawDot = (ctx, x, y, isInside) => {
  ctx.fillStyle = isInside ? 'green' : 'red';
  ctx.beginPath();
  ctx.arc(x, y, 8, 0, 2 * Math.PI);
  ctx.fill();
};

const drawUserDots = (ctx) => {
  ctx.fillStyle = 'white';
  userDots.value.forEach(dot => {
    drawDot(ctx, dot.x, dot.y, dot.isInside);
    drawPointCoords(ctx, dot.x, dot.y);
  });
};

const drawPointCoords = (ctx, x, y) => {
  ctx.fillStyle = 'blue';
  ctx.font = '16px Century Gothic';

  const unit = 100;
  const realX = (x - 400) / unit;
  const realY = (400 - y) / unit;

  ctx.fillText(`(${realX}, ${realY})`, x + 10, y - 10);
};

const draw = () => {
  const canvas = canvasRef.value;
  const ctx = canvas.getContext('2d');

  ctx.clearRect(0, 0, canvas.width, canvas.height);
  drawAxis(ctx);
  labelAxis(ctx);
  drawFigure(ctx, radius.value);
  drawUserDots(ctx);
};

onMounted(() => {
  draw();
});
</script>

<template>
  <canvas ref="canvasRef" width="800" height="800" @click="handleClick"></canvas>
</template>

<style scoped>
canvas {
  position: absolute;
  right: 10%;
  top: 5%;
  border-radius: 20%;
  transform: skewY(10deg);
}
</style>
