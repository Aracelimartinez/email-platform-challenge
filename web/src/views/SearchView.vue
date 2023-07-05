<template>
  <main class="container text-color-text-header">

    <!-- Search Bar -->
    <div class="sticky top-[100px] mb-8 bg-color-primary-bg">
      <input type="text" v-model="searchQuery" @input="getSearchResults" placeholder="Search for a term..." class="py-2 px-3 w-full bg-white shadow-md rounded focus:outline-none focus:shadow-color-primary/20 duration-150">
      <h2 v-if="!searchError  && emailsSearchResult" class="mt-6 text-xl font-bold">You have {{ emailsSearchResult.length }} result(s)...</h2>
    </div>

    <!-- Email box -->
    <div class="w-full font-light">
      <!-- Default box view -->
      <div v-if="!searchQuery" class="flex flex-col justify-center h-96 text-center  text-color-text-details/80">
        <i class="fa-solid fa-magnifying-glass text-8xl sm:text5xl"></i>
        <p class="mt-3 italic">Start searching for emails</p>
      </div>

      <!-- Display error message -->
      <div v-else-if="searchError" class="flex flex-col justify-center h-96 text-center  text-color-text-details/80">
        <i class="fa-solid fa-circle-xmark text-8xl sm:text5xl"></i>
        <p class="mt-3 italic">Sorry, something went wrong, please try again.</p>
      </div>

      <!-- Display no results message -->
      <div v-else-if="!searchError  && emailsSearchResult && emailsSearchResult.length === 0" class="flex flex-col justify-center h-96 text-center text-color-text-details/80">
        <i class="fa-solid fa-lightbulb text-8xl sm:text5xl"></i>
        <p class="mt-3 italic">No results match your term, try a different word</p>
      </div>

    <!-- Display emails list result -->
      <div v-else v-for="emailResult in emailsSearchResult" :key="emailResult.id" class=" bg-white rounded shadow-md border-gray-100 text-base p-5 my-3 text-color-text-details/80">
        <div class="text-transparent bg-clip-text bg-gradient-to-r from-color-primary to-color-secondary font-bold mb-2">
          {{ emailResult.subject }}
        </div>
        <div class="flex flex-col sm:flex-row text-color-text-header my-2 mr-4">
          <div class="flex mb-2">
            <i class="fa-solid fa-user"></i>
            <p class="text-sm mx-2 font-normal">{{ emailResult.from }}</p>
          </div>
          <div class="flex">
            <i class="fa-solid fa-calendar"></i>
            <p class="text-sm mx-2 font-normal">{{ emailResult.date }}</p>
          </div>
        </div>
        <div>

          <p class="text-sm"  v-html="highlight(emailResult.body)"></p>
          <!-- <p class="text-sm">{{ emailResult.body }}</p> -->
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

const highlight = (paragraph) => {
  if(searchQuery.value !== "") {
    return paragraph.replace(new RegExp(searchQuery.value, "gi"), match => {
      return '<span class="highlight"><strong>' + match + '</strong></span>';
    });
  }
}
</script>

<style>
.highlight {
  /* background: rgba(213, 42, 71);
  color: #ffffff */
  text-decoration: underline;
  text-decoration-color: #F5D751;
  text-decoration-thickness: 3px;
  text-underline-offset: 3px;
}
</style>
