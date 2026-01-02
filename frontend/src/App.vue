<template>
  <div class="app">
    <el-container direction="vertical">
      <Header />
      <el-main class="bg-$el-bg-color-2">
        <router-view />
      </el-main>
      <Footer />
    </el-container>
  </div>
</template>

<script lang="ts">
import Header from './components/Header.vue'
import Footer from './components/Footer.vue'

export default {
  name: 'App',
  components: {
    Header,
    Footer
  }
}
</script>

<style>
/* 移除原来的CSS，保留必要的全局样式 */
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

body,
html {
  margin: 0 !important;
  padding: 0 !important;
}

/* 暗色模式支持 */
html.dark .el-main {
  background-color: var(--el-bg-color);
}
</style>