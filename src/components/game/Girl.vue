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
        <img src="@/assets/img/un_1_body.png" alt="Girl1 body"/>
        <img src="@/assets/img/un_1_pioneer.png" alt="Girl1 clothes"/>

        <transition name="fade">
          <img src="@/assets/img/un_1_shy.png" alt="Girl1 shy" v-if="hovered"/>
          <img src="@/assets/img/un_1_smile.png" alt="Girl1 smile" v-else/>
        </transition>
      </div>
      <div v-else>
        <img src="@/assets/img/un_2_body.png" alt="Girl2 body"/>
        <img src="@/assets/img/un_2_pioneer.png" alt="Girl2 clothes"/>
        <img src="@/assets/img/un_2_cry.png" alt="Girl2 sad"/>
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
