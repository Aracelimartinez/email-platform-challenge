<template>
  <div class="flex flex-col min-h-screen font-Roboto bg-gradient-to-r from-color-primary to-color-secondary">
    <SiteNavigation />
    <RouterView />
  </div>
</template>

<script setup>
import { RouterView } from 'vue-router';
import SiteNavigation from './components/SiteNavigation.vue';


</script>

<style lang="scss" scoped>

</style>
