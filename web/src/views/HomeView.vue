<template>
  <main class="container">
    <div class="text-center flex flex-col items-center justify-center mt-7 py-4">
      <h1 class="text-3xl md:text-5xl font-bold leading-relaxed md:leading-relaxed text-transparent bg-clip-text bg-gradient-to-r from-color-primary to-color-secondary">
        Welcome to the Email Searcher Platform!
      </h1>
      <p class="text-xl font-light text-color-text-header my-5">This website search for emails at the Enron data set. Search for any term and see the results.</p>
      <button @click="redirectToSearch" class="text-white bg-gradient-to-r from-color-primary to-color-secondary rounded py-2 px-6 hover:opacity-80 duration-150">Start searching
        <i class="fa-solid fa-arrow-right ml-2"></i>
      </button>
    </div>
  </main>
</template>

<script setup>
  import { useRouter } from 'vue-router';

  const router = useRouter();

  const redirectToSearch = () => {
  router.push('/search');
  };
</script>

<style lang="scss" scoped>

</style>
