import { createApp } from 'vue'
import App from './App.vue'
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
// 导入UnoCSS
import "virtual:uno.css";
import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'

// 导入组件
import Home from './components/Home.vue'
import PlatformList from './components/PlatformList.vue'
import PlatformDetail from './components/PlatformDetail.vue'
import HistoryQuery from './components/HistoryQuery.vue'

// 定义路由类型
const routes: RouteRecordRaw[] = [
  { path: '/', component: Home },
  { path: '/platforms', component: PlatformList },
  { path: '/platform/:name', name: 'PlatformDetail', component: PlatformDetail, props: true },
  { path: '/history', component: HistoryQuery }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

// 创建应用
const app = createApp(App)

// 注册所有图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.use(router)

// 检查系统是否为暗色模式
if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
  document.body.classList.add('dark')
}

app.mount('#app')