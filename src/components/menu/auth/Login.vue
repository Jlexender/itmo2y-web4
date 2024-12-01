<script setup>
import {inject, ref} from "vue";

defineEmits(['toMain', 'toOptions', 'toRegister']);

const hasAuthenticated = inject('auth');

const loginInput = ref('');
const passwordInput = ref('');


const login = async () => {
  try {
    await fetch('http://localhost:3080/api/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        name: loginInput.value,
        pass: passwordInput.value
      })
    }).then(response => {
      if (response.ok) {
        console.log("Успешная авторизация!")
        hasAuthenticated.value = true;
      } else {}
    });
  } catch (e) {
    console.log(e);
  }

}
</script>

<template>
  <div>
    <img alt="Login" class="game-canvas" draggable="false" src="@/assets/img/load_bg.jpg"/>

    <header>
      <span class="star">☆</span> Логин <span class="star">☆</span>
    </header>

    <nav>
      <ul>
        <li @click="$emit('toOptions')">
          Настройки
        </li>
        <li @click="$emit('toRegister')">
          Регистрация
        </li>
      </ul>
    </nav>

    <form autocomplete="off">
      <label for="login">Логин</label>
      <input type="text" id="login" name="login" v-model="loginInput" required> <br>
      <label for="password">Пароль</label>
      <input type="password" id="password" name="password" v-model="passwordInput" required> <br>
      <button @click.prevent="login">Войти</button>
    </form>


    <div class="go-back" @click="$emit('toMain')">
      ← Назад
    </div>
  </div>
</template>

<style scoped>
header {
  position: absolute;
  font-size: 58px;
  text-transform: uppercase;
  font-family: "Century Gothic", sans-serif;
  top: 10%;
  left: 50%;
  transform: translateX(-50%);
}

nav {
  display: flex;
  justify-content: center;
  position: absolute;
  width: 100%;
  top: 5%;
  font-size: 48px;
  font-family: Roboto, sans-serif;
  color: gray;
}

form {
  position: absolute;
  top: 20%;
  left: 50%;
  transform: translateX(-50%);
  font-size: 38px;
  font-family: "Century Gothic", sans-serif;
  color: white;
  display: flex;
  flex-direction: column;
}

label {
  color: #a2ca5a;
}

input {
  font-size: 48px;
  font-family: "Century Gothic", sans-serif;
  color: white;
  background-color: transparent;
  border: none;
  border-bottom: 2px solid gray;
  padding: 5px;
  cursor: none;
}

input:focus {
  outline: none;
  cursor: none;
}

button {
  font-size: 48px;
  font-family: "Century Gothic", sans-serif;
  color: gray;
  background-color: transparent;
  padding: 5px;
  border: none;
}

button:hover {
  color: white;
  cursor: none;
}

ul {
  list-style-type: none;
  padding: 0;
  margin: 0;
  display: flex;
  gap: 900px;
}

ul li {
  display: flex;
  justify-content: center;
  font-size: 58px;
  font-family: "Century Gothic", sans-serif;
  text-transform: uppercase;
}

.star {
  color: gray;
  font-weight: bold;
}

.go-back {
  position: absolute;
  bottom: 10%;
  left: 50px;
  font-size: 58px;
  font-family: "Century Gothic", sans-serif;
  cursor: none;
  color: gray;
}

.go-back:hover, ul li:hover {
  color: white;
}
</style>
