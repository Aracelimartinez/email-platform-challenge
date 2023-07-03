<template>
  <header class="sticky top-0 bg-gradient-to-r from-color-primary to-color-secondary">
    <nav class="container flex flex-col sm:flex-row items-center gap-4 text-white py-6">
      <RouterLink :to="{name: 'home' }">
        <div class="flex items-center gap-3">
          <i class="fa-solid fa-paper-plane text-2xl"></i>
          <p class="text-2xl">Email searcher</p>
        </div>
      </RouterLink>
      <div class="flex gap-3 flex-1 justify-end">
        <i class="fa-solid fa-circle-info text-xl hover:opacity-50 duration-150 cursor-pointer"></i>
      </div>
    </nav>
  </header>
</template>

<script setup>
import { RouterLink } from "vue-router"
</script>

<style lang="scss" scoped>

</style>
