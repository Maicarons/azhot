<template>
  <div class="platform-list">
    <StarPrompt @close="() => {}" />
    <h2>所有平台列表</h2>
    <el-table 
      :data="platforms" 
      stripe 
      style="width: 100%"
      @row-click="handleRowClick"
    >
      <el-table-column prop="routeName" label="平台路由名" width="150" />
      <el-table-column prop="name" label="平台名称" width="150" />
      <el-table-column label="图标" width="100">
        <template #default="scope">
          <img :src="scope.row.icon" :alt="scope.row.name" class="platform-icon" />
        </template>
      </el-table-column>
      <el-table-column prop="description" label="描述" />
      <el-table-column label="操作" width="150">
        <template #default="scope">
          <el-button 
            size="small" 
            @click.stop="goToPlatform(scope.row.routeName)"
          >
            查看
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import StarPrompt from './StarPrompt.vue'

// 定义平台类型
interface Platform {
  routeName: string;
  name: string;
  icon: string;
  description?: string;
}

// 设置axios的基础URL，优先从环境变量获取后端服务器地址
const BACKEND_BASE_URL = import.meta.env.VITE_API_BASE_URL ||  
                         'http://localhost:8080' // 默认值

axios.defaults.baseURL = BACKEND_BASE_URL

export default {
  name: 'PlatformList',
  components: {
    StarPrompt
  },
  setup() {
    const router = useRouter()
    const platforms = ref<Platform[]>([])

    // 平台描述信息映射
    const platformDescriptions: Record<string, string> = {
      '360doc': '获取360doc热门文章列表',
      '360search': '获取360搜索热点排行榜',
      'acfun': '获取AcFun热门排行榜',
      'baidu': '获取百度热搜列表',
      'bilibili': '获取哔哩哔哩热门视频列表',
      'cctv': '获取CCTV新闻热点排行榜',
      'csdn': '获取CSDN热门博客文章排行榜',
      'dongqiudi': '获取懂球帝热门体育资讯排行榜',
      'douban': '获取豆瓣热门搜索列表',
      'douyin': '获取抖音热搜列表',
      'github': '获取GitHub Trending列表',
      'guojiadili': '获取国家地理热门文章排行榜',
      'historytoday': '获取历史上的今天事件列表',
      'hupu': '获取虎扑热门体育资讯排行榜',
      'ithome': '获取IT之家热搜列表',
      'lishipin': '获取梨视频热门视频排行榜',
      'nanfang': '获取南方周末热门内容排行榜',
      'pengpai': '获取澎湃新闻热点新闻排行榜',
      'qqnews': '获取腾讯新闻热点排行榜',
      'quark': '获取夸克热点搜索排行榜',
      'renmin': '获取人民网热门新闻排行榜',
      'sougou': '获取搜狗热点搜索排行榜',
      'souhu': '获取搜狐热点新闻排行榜',
      'toutiao': '获取今日头条热搜列表',
      'v2ex': '获取V2EX热议话题列表',
      'wangyinews': '获取网易新闻热点排行榜',
      'weibo': '获取微博热搜列表',
      'xinjingbao': '获取新京报热点新闻排行榜',
      'zhihu': '获取知乎热搜列表',
      'all': '获取所有平台的热搜列表',
      'list': '获取所有可用的热搜来源名称列表'
    }

    const goToPlatform = (platformName: string): void => {
      router.push(`/platform/${platformName}`)
    }

    const handleRowClick = (row: Platform): void => {
      goToPlatform(row.routeName)
    }

    onMounted(async () => {
      try {
        // 获取平台列表
        const response = await axios.get('/list')
        platforms.value = response.data.obj.map((platform: Platform) => ({
          ...platform,
          description: platformDescriptions[platform.routeName] || '获取平台热搜列表'
        }))
      } catch (error) {
        console.error('获取平台列表失败:', error)
        // 如果API调用失败，显示预定义的平台列表
        platforms.value = Object.entries(platformDescriptions).map(([routeName, description]) => ({
          routeName,
          name: routeName,
          icon: 'https://via.placeholder.com/16x16', // 使用占位符图标
          description
        }))
      }
    })

    return {
      platforms,
      goToPlatform,
      handleRowClick
    }
  }
}
</script>

<style scoped>
.platform-list {
  max-width: 1200px;
  margin: 0 auto;
}

h2 {
  margin-bottom: 20px;
}

.platform-icon {
  width: 16px;
  height: 16px;
  vertical-align: middle;
}

/* 暗色模式下标题颜色 */
html.dark h2 {
  color: var(--el-text-color-primary);
}
</style>