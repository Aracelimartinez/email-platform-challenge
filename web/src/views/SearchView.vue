<template>
  <main class="container text-white">
    <div class="pt-4 mb-8 relative">
      <input type="text" v-model="searchQuery" @input="getSearchResults" placeholder="Search for a term" class="py-2 px-1 w-full bg-transparent border-b focus:rounded focus:outline-none focus:text-gray-color focus:bg-white duration-150">
    </div>

    <div class=" w-full h-96 mt-4 bg-color-primary-bg rounded ">
something
    </div>
  </main>
</template>

<script setup>
import { ref } from "vue";
import api from "../services/api";

const searchQuery = ref("");
const queryTimeout = ref(null);
const emailsSearchResult = ref(null);

const getSearchResults = () => {
  clearTimeout(queryTimeout.value);
  queryTimeout.value = setTimeout(async () => {
    if (searchQuery.value !== "") {
      const results = await api.get("/search", {
        params: {
          query: searchQuery.value
        }
      });
      emailsSearchResult.value = results.data;
      console.log(emailsSearchResult.value)
      return;
    }
    emailsSearchResult.value = null;
  }, 300)
}


// import { onMounted } from "vue";
// import api from "../services/api.js"

// const emails = ref([]);
// const searchEmails = async (query) => {
//   try {
//     const response = await api.get('/search', { params: { query } });
//     console.log(response.data);
//     return response.data;
//   } catch (error) {
//     console.error(error);
//     throw error;
//   }
// };

// onMounted(searchEmails);

</script>
