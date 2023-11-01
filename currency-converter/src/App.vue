<template>
  <div id="app">
    <!-- Logo -->
    <img alt="Vue logo" src="@/assets/logo.png" class="app-logo">

    <!-- Currency Converter Component -->
    <ConvertCurrency msg="Welcome to Currency Converter App"/>
  </div>
</template>

<script>
import ConvertCurrency from '@/components/ConvertCurrency.vue'

export default {
  name: 'App',
  components: {
    ConvertCurrency
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin: 60px auto;
  padding: 20px;
  max-width: 500px; /* Adjust based on your preference */
  background-color: #f4f4f2; /* Off-white background */
  border-radius: 10px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.app-logo {
  max-width: 120px; /* Adjust as needed */
  margin-bottom: 20px;
}
</style>
