<script setup>

</script>

<!-- FIXME BAD CODE !!!! -->
<template>
  <div class="transparent exit-area">
    Exit
  </div>

  <div class="transparent start-area">
    Start
  </div>

  <div class="transparent options-area">
    Options
  </div>
</template>

<style scoped>
.transparent {
  opacity: 0.5;
}

.exit-area {
  position: absolute;
  top: 50%;
  left: 75%;
  width: 10%;
  height: 25vh;
  background-color: red;
  color: white;
  font-size: 40px;
  transform: skewY(-10deg);
}

.start-area {
  position: absolute;
  top: 24%;
  left: 24%;
  width: 17vw;
  height: 57vh;
  background-color: blue;
  color: white;
  font-size: 40px;
  transform: skewY(-6deg);
}

.options-area {
  position: absolute;
  top: 76%;
  left: 55%;
  width: 13vw;
  height: 21vh;
  background-color: green;
  font-size: 40px;
}
</style>
