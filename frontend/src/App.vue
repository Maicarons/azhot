<template>
  <div class="app">
    <el-container direction="vertical">
      <Header />
      <el-main class="bg-$el-bg-color-2">
        <el-alert
          title="内容免责声明"
          type="info"
          :closable="false"
          show-icon
          class="disclaimer-alert"
        >
          <p>本网站仅提供内容聚合服务，展示各大平台的热搜信息，所有数据均来源于相应平台的公开接口。</p>
          <p>本网站不对所展示内容的真实性、准确性、合法性负责，亦不认可其中任何观点或立场。</p>
        </el-alert>
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

/* 免责声明样式 */
.disclaimer-alert {
  margin-bottom: 20px;
  border-radius: 4px;
}
</style>