<template>
  <header class="sticky top-0 bg-weather-primary shadow-lg">
    <nav class="container flex flex-col items-center gap-4 text-white py-6">
      <RouterLink :to="{ name: 'home' }">
        <div class="flex items-center gap-3">
          <i class="fa-solid fa-envelope text-2xl"></i>
          <p class="text-2xl">Email Imbox</p>
        </div>
      </RouterLink>

    </nav>
  </header>
</template>
  
<script setup>
import { RouterLink } from "vue-router";
</script>