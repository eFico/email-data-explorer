<template>
  <div class="flex flex-col  font-Roboto bg-weather-primary min-h-screen">
    <SiteNavigation/>
    <RouterView/>
  </div>
</template>

<script setup>
import {RouterView} from "vue-router"
import SiteNavigation from "./components/SiteNavigation.vue";
</script>

<style lang="scss" scoped>

</style>