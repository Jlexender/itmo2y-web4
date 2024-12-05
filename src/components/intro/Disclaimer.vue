<script setup>
import { ref, onMounted } from 'vue';
import disclaimerImage from '@/assets/img/disclaimer.jpg';

const emit = defineEmits(['toSplash']);
const imagesLoaded = ref(false);
const audiosLoaded = ref(false);
const userClicked = ref(false);

const loadImagesAndAudios = async () => {
  const images = import.meta.glob('@/assets/img/*.{png,jpg,jpeg,svg}');
  const loadImagePromises = Object.values(images).map(async (importImage) => {
    const image = new Image();
    image.src = (await importImage()).default;
    return new Promise((resolve) => {
      image.onload = resolve;
    });
  });

  const audios = import.meta.glob('@/assets/audio/*.{mp3,wav,ogg}');
  const loadAudioPromises = Object.values(audios).map(async (importAudio) => {
    const audio = new Audio();
    audio.src = (await importAudio()).default;
    return new Promise((resolve) => {
      audio.onloadeddata = resolve;
    });
  });

  await Promise.all([...loadImagePromises, ...loadAudioPromises]);

  imagesLoaded.value = true;
  audiosLoaded.value = true;

  if (userClicked.value) {
    emit('toSplash');
  }
};

const handleClick = () => {
  if (imagesLoaded.value && audiosLoaded.value) {
    emit('toSplash');
  } else {
    userClicked.value = true;
  }
};

onMounted(() => {
  loadImagesAndAudios();
});
</script>

<template>
  <div @click="handleClick">
    <img alt="Disclaimer" class="game-canvas" draggable="false" src="@/assets/img/disclaimer.jpg"/>
  </div>
</template>

<style scoped>
</style>
