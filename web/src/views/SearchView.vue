<template>
  <main class="container text-color-text-header">

    <!-- Search Bar -->
    <div class="pt-4 mb-8 relative">
      <input type="text" v-model="searchQuery" @input="getSearchResults" placeholder="Search for a term..." class="my-6 py-2 px-3 w-full bg-white shadow-md rounded focus:outline-none focus:shadow-color-primary/20 duration-150">
    </div>

    <!-- Email box -->
    <div class="flex flex-col justify-center text-color-text-details w-full h-96 bg-white rounded shadow-md p-4 overflow-y-scroll font-light">
      <!-- Default box view -->
      <div v-if="!searchQuery" class="text-center">
        <i class="fa-solid fa-magnifying-glass text-8xl"></i>
        <p class="mt-3 italic">Start searching for emails</p>
      </div>

      <!-- Display error message -->
      <div v-else-if="searchError" class="text-center">
        <i class="fa-solid fa-circle-xmark text-8xl"></i>
        <p class="mt-3 italic">Sorry, something went wrong, please try again.</p>
      </div>

      <!-- Display no results message -->
      <div v-else-if="!searchError  && emailsSearchResult && emailsSearchResult.length === 0" class="text-center">
        <i class="fa-solid fa-lightbulb text-8xl"></i>
        <p class="mt-3 italic">No results match your term, try a different word</p>
      </div>

    <!-- Display emails list result -->
      <div v-else v-for="emailResult in emailsSearchResult" :key="emailResult.id" class="border-y border-gray-100 text-base p-4">
        <div class="text-color-text-header font-bold">
          {{ emailResult.subject }}
        </div>
        <div class="flex flex-col sm:flex-row text-color-text-header my-2">
          <div class="flex">
            <i class="fa-solid fa-user"></i>
            <p class="text-sm mx-2 font-normal">{{ emailResult.from }}</p>
          </div>
          <div class="flex  ml-4">
            <i class="fa-solid fa-calendar"></i>
            <p class="text-sm mx-2 font-normal">{{ emailResult.date }}</p>
          </div>
        </div>
        <div>
          <p class="text-sm">{{ emailResult.body }}</p>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup>
import { ref } from "vue";

import api from "@/services/api";

const searchQuery = ref("");
const queryTimeout = ref(null);
const emailsSearchResult = ref(null);
const searchError = ref(null);

const getSearchResults = () => {
  clearTimeout(queryTimeout.value);
  queryTimeout.value = setTimeout(async () => {
    if (searchQuery.value !== "") {
      try {
        const results = await api.get("/search", {
          params: {
            query: searchQuery.value
          }
        });
        emailsSearchResult.value = results.data;

      } catch {
        searchError.value = true
      }

      return;
    }
    emailsSearchResult.value = null;
  }, 300)
}
</script>

<style scoped>
.highlight {
  background-color: yellow;
}

</style>
