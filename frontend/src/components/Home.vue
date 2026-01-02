<template>
  <div class="home">
    <h2>azhot</h2>
    <p>获取各大平台的实时热搜数据</p>

    <el-row :gutter="20">
      <el-col :span="8" v-for="platform in platforms" :key="platform.routeName">
        <el-card class="platform-card" @click="goToPlatform(platform.routeName)">
          <div class="platform-info">
            <div class="platform-header">
              <img :src="platform.icon" :alt="platform.name" class="platform-icon" />
              <h3>{{ platform.name }}</h3>
            </div>
            <p>{{ getDescription(platform.routeName) }}</p>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 添加历史记录查询入口 -->
    <el-row :gutter="20" style="margin-top: 30px;">
      <el-col :span="8" :offset="8">
        <el-card class="platform-card" @click="goToHistoryQuery">
          <div class="platform-info">
            <h3>历史记录查询</h3>
            <p>查询历史上的热搜数据</p>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

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
  name: 'Home',
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

    const getDescription = (platformName: string): string => {
      return platformDescriptions[platformName] || '获取平台热搜列表'
    }

    const goToPlatform = (platformName: string): void => {
      if (platformName === 'list' || platformName === 'all') {
        router.push(`/platforms`)
      } else {
        router.push(`/platform/${platformName}`)
      }
    }

    const goToHistoryQuery = (): void => {
      router.push('/history')
    }

    onMounted(async () => {
      try {
        // 获取平台列表
        const response = await axios.get('/list')
        platforms.value = response.data.obj
      } catch (error) {
        console.error('获取平台列表失败:', error)
        // 如果API调用失败，显示预定义的平台列表
        platforms.value = Object.entries(platformDescriptions).map(([routeName, description]) => ({
          routeName,
          name: routeName,
          icon: 'https://placeholder.im/16x16/AZHOT/47a6ff/000000', // 使用占位符图标
          description
        }))
      }
    })

    return {
      platforms,
      goToPlatform,
      goToHistoryQuery,
      getDescription
    }
  }
}
</script>

<style scoped>
.home {
  max-width: 1200px;
  margin: 0 auto;
}

h2 {
  margin-bottom: 10px;
}

p {
  color: #606266;
  margin-bottom: 30px;
}

/* 暗色模式下文本颜色 */
html.dark p {
  color: var(--el-text-color-regular);
}

.platform-card {
  margin-bottom: 20px;
  cursor: pointer;
  transition: all 0.3s;
}

.platform-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15);
}

.platform-header {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.platform-icon {
  width: 24px;
  height: 24px;
  margin-right: 10px;
  vertical-align: middle;
}

.platform-info h3 {
  margin: 0 0 10px 0;
  color: #409EFF;
  display: inline-block;
}

/* 暗色模式下标题颜色 */
html.dark .platform-info h3 {
  color: var(--el-color-primary);
}

.platform-info p {
  color: #909399;
  font-size: 14px;
  line-height: 1.5;
}

/* 暗色模式下平台信息文本颜色 */
html.dark .platform-info p {
  color: var(--el-text-color-secondary);
}
</style>