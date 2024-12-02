<script setup>
const emit = defineEmits(['toDisclaimer'])

setTimeout(() => {
  emit('toDisclaimer')
}, 3000)
</script>

<template>
  <div>
    <div class="author">
      <ul>
        <li>ФИО: <strong>Калиев Александр Дмитриевич</strong></li>
        <li>Группа: <strong>P3211</strong></li>
        <li>Вариант: <strong>11501</strong></li>
      </ul>

      <transition appear name="fade">
        <img alt="Developers" class="game-canvas" src="@/assets/img/soviet_games.jpg">
      </transition>
    </div>
  </div>
</template>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity 2s ease-in-out;
}

ul {
  list-style-type: none;
  padding: 0;
}

.author {
  font-family: "Century Gothic", sans-serif;
  left: 30px;
  top: 30px;
  width: 100%;
  height: 100%;
  position: absolute;
  font-size: 48px;
  color: white;
}

.author-avatar {
  height: 200px;
  width: 200px;
  position: absolute;
  top: 30px;
  right: 100px;
  border-radius: 50%;
}

.game-canvas {
  z-index: -1;
}

</style>
