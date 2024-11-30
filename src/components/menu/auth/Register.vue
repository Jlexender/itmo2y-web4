<script setup>
defineEmits(['toMain', 'toOptions', 'toLogin']);
</script>

<template>
  <div>
    <img alt="Login" class="game-canvas" draggable="false" src="@/assets/img/save_bg.jpg"/>

    <header>
      <span class="star">☆</span> Регистрация <span class="star">☆</span>
    </header>

    <form autocomplete="off">
      <label for="login">Логин</label>
      <input type="text" id="login" name="login" required> <br>
      <label for="password">Пароль</label>
      <input type="password" id="password" name="password" required> <br>
      <label for="password-repeat">Повторите пароль</label>
      <input type="password" id="password-repeat" name="password-repeat" required> <br>
      <button @click.prevent>Зарегистрироваться</button>
    </form>

    <nav>
      <ul>
        <li @click="$emit('toOptions')">
          Настройки
        </li>
        <li @click="$emit('toLogin')">
          Логин
        </li>
      </ul>
    </nav>


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

ul {
  list-style-type: none;
  padding: 0;
  margin: 0;
  display: flex;
  gap: 1100px;
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
