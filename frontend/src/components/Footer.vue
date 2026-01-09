<template>
  <el-footer
    class="footer border-t border-$el-border-color py-[40px] !m-0 w-full text-$el-text-color-regular text-0.9rem bg-$el-bg-color">
    <div class="footer-content max-w-[1200px] mx-auto px-5">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
        <!-- azhot 介绍 -->
        <div class="footer-section text-center">
          <h3 class="text-1.3rem font-bold color-$el-text-color-primary mb-3">
            azhot
          </h3>
          <p class="color-$el-text-color-secondary leading-6 mb-4">
            聚合各大平台热搜数据，提供统一API接口
          </p>
        </div>

        <!-- 快速链接 -->
        <div class="footer-section text-center">
          <h4 class="mb-4 text-[1rem] font-semibold color-$el-text-color-primary">
            快速链接
          </h4>
          <ul class="list-none p-0 m-0 space-y-2">
            <li>
              <router-link to="/"
                class="color-$el-text-color-regular no-underline transition-colors-300ms flex items-center justify-center hover:color-$el-color-primary">
                <el-icon :size="16" class="icon mr-[8px] vertical-align-middle">
                  <House />
                </el-icon>
                首页
              </router-link>
            </li>
            <li>
              <router-link to="/platforms"
                class="color-$el-text-color-regular no-underline transition-colors-300ms flex items-center justify-center hover:color-$el-color-primary">
                <el-icon :size="16" class="icon mr-[8px] vertical-align-middle">
                  <Menu />
                </el-icon>
                平台列表
              </router-link>
            </li>
            <li>
              <router-link to="/history"
                class="color-$el-text-color-regular no-underline transition-colors-300ms flex items-center justify-center hover:color-$el-color-primary">
                <el-icon :size="16" class="icon mr-[8px] vertical-align-middle">
                  <Clock />
                </el-icon>
                历史记录
              </router-link>
            </li>
          </ul>
        </div>

        <!-- 资源 -->
        <div class="footer-section text-center">
          <h4 class="mb-4 text-[1rem] font-semibold color-$el-text-color-primary">
            资源
          </h4>
          <ul class="list-none p-0 m-0 space-y-2">
            <li>
              <a href="https://github.com/maicarons/azhot" target="_blank"
                class="color-$el-text-color-regular no-underline transition-colors-300ms flex items-center justify-center hover:color-$el-color-primary">
                <el-icon :size="16" class="icon mr-[8px] vertical-align-middle">
                  <Link />
                </el-icon>
                GitHub
              </a>
            </li>
            <li>
              <a href="https://github.com/maicarons/azhot/blob/main/README.md" target="_blank"
                class="color-$el-text-color-regular no-underline transition-colors-300ms flex items-center justify-center hover:color-$el-color-primary">
                <el-icon :size="16" class="icon mr-[8px] vertical-align-middle">
                  <Document />
                </el-icon>
                文档
              </a>
            </li>
            <li>
              <a href="https://github.com/maicarons/azhot/issues" target="_blank"
                class="color-$el-text-color-regular no-underline transition-colors-300ms flex items-center justify-center hover:color-$el-color-primary">
                <el-icon :size="16" class="icon mr-[8px] vertical-align-middle">
                  <ChatLineRound />
                </el-icon>
                问题反馈
              </a>
            </li>
          </ul>
        </div>

        <!-- 联系我们 -->
        <div class="footer-section text-center">
          <h4 class="mb-4 text-[1rem] font-semibold color-$el-text-color-primary">
            联系
          </h4>
          <div class="social-links mb-4 flex justify-center space-x-4">
            <a href="https://github.com/maicarons/azhot" target="_blank" title="GitHub"
              class="w-10 h-10 rounded-full bg-$el-bg-color-overlay flex items-center justify-center hover:bg-$el-color-primary hover:text-$el-color-white transition-colors-300ms no-underline">
              <el-icon :size="20">
                <Link />
              </el-icon>
            </a>
          </div>
          <p
            class="copyright color-$el-text-color-secondary text-0.85rem mt-[15px] pt-[15px] border-t border-$el-border-color">
            © 2026 Maicarons
          </p>
        </div>
      </div>
      <div class="disclaimer mt-[20px] pt-[20px] border-t border-$el-border-color text-center pb-[15px]">
        <p class="color-$el-text-color-placeholder text-0.8rem">
          本网站仅提供内容聚合服务，展示各大平台的热搜信息，所有数据均来源于相应平台的公开接口。
        </p>
        <p class="color-$el-text-color-placeholder text-0.8rem mt-[8px]">
          本网站不对所展示内容的真实性、准确性、合法性负责，亦不认可其中任何观点或立场。
        </p>
      </div>
    </div>
  </el-footer>
</template>

<script lang="ts">
import { House, Menu, Clock, Link, Document, ChatLineRound } from '@element-plus/icons-vue'

export default {
  name: 'Footer',
  components: {
    House,
    Menu,
    Clock,
    Link,
    Document,
    ChatLineRound
  }
}
</script>

<style scoped>
/* 暗色模式下的Footer样式 */
html.dark .footer {
  background-color: var(--el-bg-color);
  border-top: 1px solid var(--el-border-color);
  color: var(--el-text-color-secondary);
}

html.dark .footer-section h3,
html.dark .footer-section h4 {
  color: var(--el-text-color-primary);
}

html.dark .footer-section p,
html.dark .footer-section a {
  color: var(--el-text-color-secondary);
}

html.dark .copyright {
  color: var(--el-text-color-secondary);
  border-top: 1px solid var(--el-border-color);
}

html.dark .disclaimer p {
  color: var(--el-text-color-placeholder);
}

html.dark .footer-section a:hover {
  color: var(--el-color-primary);
}

html.dark .social-links a:hover {
  color: var(--el-color-primary);
  background-color: var(--el-fill-color-darker) !important;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .footer-content {
    text-align: center;
  }

  .footer-section ul li {
    display: flex;
    justify-content: center;
  }

  .footer-section a {
    display: flex;
    align-items: center;
    justify-content: center;
  }
  
  .disclaimer {
    text-align: center;
  }
}
</style>