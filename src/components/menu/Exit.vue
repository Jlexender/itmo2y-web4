<script setup>

const exitMessage = "Вы действительно хотите выйти из моей лабы?";
</script>

<template>
  <div>
    <img src="@/assets/img/exit.jpg" alt="Exit" />

    <div>
      {{ exitMessage }}
    </div>

    <ul>
      <li>
        <button @click="$emit('exit')">Да</button>
      </li>
      <li>
        <button @click="$emit('back')">Нет</button>
      </li>
    </ul>
  </div>
</template>

<style scoped>

</style>
