<script setup>

import {ref} from "vue";

defineProps({
  notSad: Boolean,
});

const hovered = ref(false);

</script>

<template>
  <div>
    <transition name="fade">
      <div v-if="notSad">
        <img alt="Girl1 body" src="@/assets/img/un_1_body.png"/>
        <img alt="Girl1 clothes" src="@/assets/img/un_1_pioneer.png"/>

        <transition name="fade">
          <img v-if="hovered" alt="Girl1 shy" src="@/assets/img/un_1_shy.png"/>
          <img v-else alt="Girl1 smile" src="@/assets/img/un_1_smile.png"/>
        </transition>
      </div>
      <div v-else>
        <img alt="Girl2 body" src="@/assets/img/un_2_body.png"/>
        <img alt="Girl2 clothes" src="@/assets/img/un_2_pioneer.png"/>
        <img alt="Girl2 sad" src="@/assets/img/un_2_cry.png"/>
      </div>
    </transition>
    <div class="hover-area" @mouseenter="hovered = true" @mouseleave="hovered = false"/>
  </div>

</template>

<style scoped>
.hover-area {
  opacity: 0;
  background-color: red;
  position: absolute;
  width: 430px;
  height: 900px;
  bottom: 0;
  left: 170px;
  border-top-left-radius: 50%;
  border-top-right-radius: 50%;
}

img {
  position: absolute;
  right: 50%;
  z-index: 0;
}
</style>
