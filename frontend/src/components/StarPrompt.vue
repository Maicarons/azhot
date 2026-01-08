<template>
  <div class="star-prompt">
    <el-alert
      title="喜欢这个项目吗？"
      type="success"
      :closable="false"
      show-icon
      class="star-alert"
    >
      <template #default>
        <div class="alert-content">
          <p>如果您觉得这个项目对您有帮助，请在 GitHub 上给我们点个 Star！⭐</p>
          <div class="button-group">
            <el-button 
              type="primary" 
              size="small"
              @click="goToGitHub"
            >
              GitHub 仓库
            </el-button>
          </div>
        </div>
      </template>
    </el-alert>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'StarPrompt',
  emits: ['close'],
  setup(props, { emit }) {
    const goToGitHub = () => {
      window.open('https://github.com/Maicarons/azhot', '_blank')
    }

    const closePrompt = () => {
      emit('close')
    }

    return {
      goToGitHub,
      closePrompt
    }
  }
})
</script>

<style scoped>
.star-prompt {
  margin-bottom: 20px;
}

.star-alert {
  border-radius: 8px;
}

.alert-content {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.button-group {
  display: flex;
  gap: 10px;
}

/* 暗色模式下的样式 */
html.dark .star-alert {
  background-color: var(--el-bg-color-overlay);
}
</style>