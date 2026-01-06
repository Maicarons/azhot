import Unocss from "unocss/vite";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    Unocss(),
    vue(),
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
  ],
  server: {
    host: "0.0.0.0",
    port: 3000,
    open: true,
  },
  build: {
    chunkSizeWarningLimit: 1000, // 增加到1000kb，避免块大小警告
    rollupOptions: {
      output: {
        manualChunks: {
          // 将 Vue 和 Vue Router 分割成单独的块
          'vue': ['vue', 'vue-router'],
          // 将 Element Plus 单独分割成一个块
          'element-plus': ['element-plus'],
          // 将 UI 相关库分割到一起
          'ui-lib': ['@element-plus/icons-vue'],
          // 将工具库分割到一起
          'utils': ['axios'],
        }
      }
    },
  },
});