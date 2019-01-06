<template>
  <div id="app">
    <Notes/>
  </div>
</template>

<script>
import Notes from './components/Notes'

export default {
  name: 'App',
  components: {
    Notes
  }
}
</script>

<style>
#app {
  font-family: 'Noto Sans', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin-top: 60px;

  @import url('https://fonts.googleapis.com/css?family=Noto+Sans');
}
</style>
