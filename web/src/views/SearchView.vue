<template>
  <main class="container text-color-text-header">
    <div class="pt-4 mb-8 relative">
      <input type="text" v-model="searchQuery" @input="getSearchResults" placeholder="Search for a term..." class="my-6 py-2 px-3 w-full bg-white shadow-md rounded focus:outline-none focus:shadow-color-primary/20 duration-150">
    </div>

    <div class="flex flex-col items-center justify-center text-color-text-details w-full h-96 bg-white rounded shadow-md p-4">
      <div v-if="!searchQuery" class="text-center">
        <i class="fa-solid fa-magnifying-glass text-8xl"></i>
        <p class="font-light mt-3 italic">Start searching for emails</p>
      </div>

      <div v-else-if="searchError" class="text-center">
        <i class="fa-solid fa-circle-xmark text-8xl"></i>
        <p class="font-light mt-3 italic">Sorry, something went wrong, please try again.</p>
      </div>

      <div v-else-if="!searchError  && emailsSearchResult && emailsSearchResult.length === 0" class="text-center">
        <i class="fa-solid fa-lightbulb text-8xl"></i>
        <p class="font-light mt-3 italic">No results match your term, try a different word</p>
      </div>

      <div v-else>
        <div v-for="emailResult in emailsSearchResult" :key="emailResult.id">
          {{ emailResult.body }}
        </div>
      </div>
    </div>
  </main>
</template>

<script setup>
import { ref } from "vue";
import api from "../services/api";

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
