<template>
  <img alt="Vue logo" src="./assets/logo.png"/>
  <!--  <HelloWorld msg="Hello Vue 3 + Vite" />-->
  <button @click="ClickFunc">Click Me</button>
</template>

<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
// import HelloWorld from './components/HelloWorld.vue'
import axios from 'axios'

function ClickFunc() {
  axios.post('http://localhost:9000/api/list_items').then((res) => {
    console.log(res)
  }).catch((err) => {
    console.log(err)
  })
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
