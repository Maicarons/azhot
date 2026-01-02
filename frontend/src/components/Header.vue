<template>
  <el-header class="!p-0 !h-20 flex items-center">
    <div class="header-content w-full h-full flex items-center justify-between px-4">
      <h1 class="m-0 text-4xl text-blue-500 flex items-center">
        azhot
      </h1>
      <div class="header-right flex items-center">
        <el-menu
          :default-active="activeIndex"
          mode="horizontal"
          router
          class="nav-menu flex-1 max-w-600px mr-20"
        >
          <el-menu-item index="/">首页</el-menu-item>
          <el-menu-item index="/platforms">平台列表</el-menu-item>
          <el-menu-item index="/history">历史记录查询</el-menu-item>
        </el-menu>
        <el-button 
          :icon="isDarkMode ? Moon : Sunny" 
          @click="toggleDarkMode"
          type="info"
          :plain="!isDarkMode"
          circle
          class="dark-mode-toggle ml-4"
        />
      </div>
    </div>
  </el-header>
</template>

<script lang="ts">
import { ref, watch, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { Moon, Sunny } from '@element-plus/icons-vue'

export default {
  name: 'Header',
  setup() {
    const route = useRoute()
    const activeIndex = ref<string>(route.path)
    const isDarkMode = ref<boolean>(false)

    // 监听路由变化，更新激活的菜单项
    watch(() => route.path, (newPath: string) => {
      activeIndex.value = newPath
    })

    // 检查本地存储中的主题设置
    onMounted(() => {
      const savedTheme = localStorage.getItem('theme')
      if (savedTheme === 'dark') {
        isDarkMode.value = true
        document.documentElement.classList.add('dark')
      } else {
        isDarkMode.value = false
        document.documentElement.classList.remove('dark')
      }
    })

    const toggleDarkMode = () => {
      isDarkMode.value = !isDarkMode.value
      if (isDarkMode.value) {
        document.documentElement.classList.add('dark')
        localStorage.setItem('theme', 'dark')
      } else {
        document.documentElement.classList.remove('dark')
        localStorage.setItem('theme', 'light')
      }
    }

    return {
      activeIndex,
      isDarkMode,
      toggleDarkMode,
      Moon,
      Sunny
    }
  }
}
</script>

<style scoped>
/* 暗色模式下的样式 */
html.dark .el-header {
  background-color: var(--el-bg-color-overlay);
  color: var(--el-color-primary);
}

/* 暗色模式下标题颜色 */
html.dark .header-content h1 {
  color: var(--el-color-primary);
}

.logo-icon {
  margin-right: 8px;
}
</style>