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
        <i class="fa-solid fa-circle-info text-xl hover:opacity-50 duration-150 cursor-pointer" @click="toggleModal"></i>
      </div>

      <BaseModal v-bind:modalActive="activeModal" @close-modal="toggleModal">
        <div class="text-black">
          <h1 class="text-2xl mb-1">About:</h1>
          <p class="mb-4">
            The email searcher allows you to search for emails with a particular content of your choice in the Enron data-set emails.
          </p>
          <h2 class="text-2xl">How it works:</h2>
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
