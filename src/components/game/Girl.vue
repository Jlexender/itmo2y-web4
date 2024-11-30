<script setup>

defineProps({
  notSad: Boolean,
});

</script>

<template>
  <transition name="fade">
    <div v-if="notSad">
      <img src="@/assets/img/un_1_body.png" alt="Girl1 body"/>
      <img src="@/assets/img/un_1_pioneer.png" alt="Girl1 clothes"/>
      <img src="@/assets/img/un_1_smile.png" alt="Girl1 smile"/>
    </div>
    <div v-else>
      <img src="@/assets/img/un_2_body.png" alt="Girl2 body"/>
      <img src="@/assets/img/un_2_pioneer.png" alt="Girl2 clothes"/>
      <img src="@/assets/img/un_2_cry.png" alt="Girl2 sad"/>
    </div>
  </transition>
</template>

<style scoped>
img {
  position: absolute;
  right: 50%;
  z-index: 0;
}
</style>
