<template>
  <header class="sticky top-0 bg-color-primary-bg">
    <nav class="container flex  items-center gap-4 py-6 mb-6 text-transparent bg-clip-text bg-gradient-to-r from-color-primary to-color-secondary border-b-2 border-b-color-primary">
      <RouterLink :to="{name: 'home' }">
        <div class="flex items-center gap-3">
          <i class="fa-solid fa-paper-plane text-2xl"></i>
          <p class="text-2xl font-bold">Email searcher</p>
        </div>
      </RouterLink>
      <div class="flex gap-3 flex-1 justify-end">
        <i class="fa-solid fa-circle-info text-xl hover:text-color-secondary hover:opacity-50 cursor-pointer" @click="toggleModal"></i>
      </div>

      <BaseModal v-bind:modalActive="activeModal" @close-modal="toggleModal">
        <div class="text-color-text-header">
          <h1 class="text-xl font-bold mb-2 text-transparent bg-clip-text bg-gradient-to-r from-color-primary to-color-secondary">About:</h1>
          <p class="mb-4">
            The email searcher allows you to search for emails with a particular content of your choice in the Enron data-set emails.
          </p>
          <h2 class="text-xl mb-2 font-bold text-transparent bg-clip-text bg-gradient-to-r from-color-primary to-color-secondary">How it works:</h2>
          <ol class="list-decimal list-inside mb-4">
            <li>
              Search a word by entering it into the search bar.
            </li>
            <li>
              Select an email within the results to display the full content.
            </li>
            <li>
              And just like that you get all the information you need in a second!
            </li>
          </ol>
        </div>
      </BaseModal>
    </nav>
  </header>
</template>

<script setup>
import { RouterLink } from "vue-router";
import BaseModal from "./BaseModal.vue";
import { ref } from "vue";

const activeModal = ref(null);
const toggleModal = () => {
  activeModal.value = !activeModal.value;
};
</script>

<!-- <style lang="scss" scoped>

</style> -->
