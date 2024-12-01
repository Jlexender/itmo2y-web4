<script setup>

</script>

<template>
  <div>
    <img src="@/assets/img/ext_camp_entrance_day.jpg" alt="UV ground" class="game-canvas" draggable="false"/>
    <img src="@/assets/img/base.png" alt="Unauthorized" class="game-canvas" draggable="false"/>
    <div class="unauthorized">
      <h1>Пожалуйста, авторизуйтесь</h1>
    </div>
  </div>
</template>

<style scoped>
.unauthorized {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

h1 {
  color: #a10000;
  font-size: 48px;
  font-family: "Century Gothic", sans-serif;
  font-weight: normal;
  transform: translateY(-50%);
}
</style>
