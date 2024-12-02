<script setup>
import {computed, ref} from 'vue';

const itemsPerPage = 8;
const currentPage = ref(1);
const items = ref([]);

const totalPages = computed(() => Math.ceil(items.value.length / itemsPerPage));

const paginatedItems = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage;
  const end = start + itemsPerPage;
  return items.value.slice(start, end);
});

const goToPage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page;
  }
};
</script>

<template>
  <div>
    <img alt="Data table background" class="game-canvas" draggable="false" src="@/assets/img/int_library_day.jpg"/>
    <table v-if="items.length !== 0">
      <thead>
      <tr>
        <th>X</th>
        <th>Y</th>
        <th>R</th>
        <th>Результат</th>
        <th>Дата отправки</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="item in paginatedItems" :key="item.date">
        <td>{{ item.x }}</td>
        <td>{{ item.y }}</td>
        <td>{{ item.r }}</td>
        <td>{{ item.result }}</td>
        <td>{{ item.date }}</td>
      </tr>
      </tbody>
    </table>
    <div v-if="items.length !== 0" class="pagination">
      <button :disabled="currentPage === 1" @click="goToPage(currentPage - 1)">Назад</button>
      <span>Страница {{ currentPage }} из {{ totalPages }}</span>
      <button :disabled="currentPage === totalPages.value" @click="goToPage(currentPage + 1)">Вперёд</button>
    </div>
    <div class="to-menu" @click="$emit('toMain')">
      В меню
    </div>
  </div>
</template>

<style scoped>

table {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 70%;
  border-collapse: collapse;
  font-size: 36px;
  font-family: Roboto, sans-serif;
  border-radius: 15px;
  overflow: hidden;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

th, td {
  padding: 16px;
  text-align: left;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

th {
  background: rgba(156, 218, 105, 0.8);
  color: white;
  text-align: center;
  font-weight: bold;
  font-size: 40px;
}

td {
  color: #ffffff;
  background-color: rgba(0, 0, 0, 0);
  transition: background-color 0.3s ease, color 0.3s ease;
}

tbody tr:nth-child(odd) {
  background-color: rgba(50, 50, 50, 0.7);
}

tbody tr:nth-child(even) {
  background-color: rgba(30, 30, 30, 0.7);
}

tbody tr:hover {
  background: rgba(255, 255, 255, 0.51);
  transition: background 0.3s ease, color 0.3s ease;
}

tr:hover td {
  color: #000;
  transition: color 0.3s ease;
}

.pagination {
  position: absolute;
  left: 50%;
  top: 5%;
  transform: translateX(-50%);
}

.pagination {
  font-size: 24px;
}

button {
  margin: 0 10px;
  background-color: transparent;
  border: 1px solid white;
  color: white;
  padding: 10px 20px;
  font-size: 36px;
  font-family: "Century Gothic", sans-serif;
  cursor: none;
  transition: background-color 0.5s ease, color 0.5s ease;
}

button:hover {
  background-color: white;
  color: black;
  transition: background-color 0.5s ease, color 0.5s ease;
}

.to-menu {
  position: absolute;
  left: 10px;
  top: 10px;
  font-size: 48px;
  font-family: "Century Gothic", sans-serif;
  color: white;
  transition: color 0.5s ease;
  z-index: 1;
}

.to-menu:hover {
  color: #69c97e;
  transition: color 0.5s ease;
}
</style>
