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
          <el-menu-item index="github" @click="goToGitHub">
            <el-icon><Link /></el-icon>
            GitHub
          </el-menu-item>
        </el-menu>
        <el-button 
          @click="toggleDarkMode"
          :type="isDarkMode ? 'warning' : 'primary'"
          :plain="false"
          circle
          class="dark-mode-toggle ml-4"
        >
          <el-icon class="dark-mode-icon">
            <Moon v-if="isDarkMode" />
            <Sunny v-else />
          </el-icon>
        </el-button>
      </div>
    </div>
  </el-header>
</template>

<script lang="ts">
import { ref, watch, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { Moon, Sunny, Link } from '@element-plus/icons-vue'

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

    // 切换暗色模式
    const toggleDarkMode = () => {
      isDarkMode.value = !isDarkMode.value
      if (isDarkMode.value) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
    }

    // 页面加载时检查系统偏好
    onMounted(() => {
      const prefersDark = window.matchMedia && 
        window.matchMedia('(prefers-color-scheme: dark)').matches
      isDarkMode.value = prefersDark
      if (prefersDark) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
    })

    // 跳转到GitHub仓库
    const goToGitHub = () => {
      window.open('https://github.com/Maicarons/azhot', '_blank')
    }

    return {
      activeIndex,
      isDarkMode,
      toggleDarkMode,
      goToGitHub
    }
  },
  components: {
    Moon,
    Sunny,
    Link
  }
}
</script>

<style scoped>
/* 暗色模式下的Header样式 */
html.dark .el-header {
  background-color: var(--el-bg-color);
  border-bottom: 1px solid var(--el-border-color);
}

/* 暗色模式下菜单项样式 */
html.dark .el-menu {
  background-color: transparent;
  border: none;
}

html.dark .el-menu-item {
  color: var(--el-text-color-primary);
}

html.dark .el-menu-item.is-active {
  color: var(--el-color-primary);
  background-color: var(--el-fill-color-lighter);
}

html.dark .el-menu-item:hover {
  color: var(--el-color-primary);
  background-color: var(--el-fill-color-light);
}

/* 菜单项样式 */
.el-menu-item {
  font-size: 16px;
}

/* 暗色模式切换按钮图标样式 */
.dark-mode-icon {
  color: #fff;
}

html.dark .dark-mode-icon {
  color: #000;
}

/* 暗色模式切换按钮样式 */
.dark-mode-toggle {
  opacity: 0.8;
}

.dark-mode-toggle:hover {
  opacity: 1;
}
</style>