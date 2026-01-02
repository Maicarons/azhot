<template>
  <div class="platform-detail">
    <el-page-header :content="platformTitle" @back="goBack" />
    
    <div v-if="loading" class="loading">
      <el-skeleton :rows="6" animated />
    </div>
    
    <div v-else-if="error" class="error">
      <el-alert title="获取数据失败" type="error" :description="error" show-icon />
    </div>
    
    <div v-else class="content">
      <el-table 
        :data="platformData" 
        stripe 
        style="width: 100%"
        :default-sort="{ prop: 'index', order: 'ascending' }"
      >
        <el-table-column prop="index" label="#" width="80" sortable />
        <el-table-column prop="title" label="标题" min-width="300">
          <template #default="scope">
            <a 
              v-if="scope.row.url" 
              :href="scope.row.url" 
              target="_blank" 
              class="title-link"
            >
              {{ scope.row.title }}
            </a>
            <span v-else>{{ scope.row.title }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="hot" label="热度" width="120" sortable />
        <el-table-column prop="desc" label="描述" width="200" />
        <el-table-column label="标签" width="120">
          <template #default="scope">
            <el-tag 
              v-if="scope.row.tag" 
              :type="getTagType(scope.row.tag)" 
              size="small"
            >
              {{ scope.row.tag }}
            </el-tag>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'

// 定义平台数据类型
interface PlatformItem {
  index: number;
  title: string;
  hot?: string | number;
  desc?: string;
  url?: string;
  tag?: string;
  [key: string]: any; // 支持其他动态属性
}

// 设置axios的基础URL，优先从环境变量获取后端服务器地址
const BACKEND_BASE_URL = 
  // @ts-ignore
  import.meta.env?.VITE_API_BASE_URL || 
  'http://localhost:8080' // 默认值

axios.defaults.baseURL = BACKEND_BASE_URL

export default {
  name: 'PlatformDetail',
  props: {
    name: {
      type: String,
      required: true
    }
  },
  setup(props) {
    const route = useRoute()
    const router = useRouter()
    const platformData = ref<PlatformItem[]>([])
    const loading = ref(true)
    const error = ref<string | null>(null)
    
    // 平台标题映射
    const platformMap: Record<string, string> = {
      '360doc': '360doc热门文章',
      '360search': '360搜索热点',
      'acfun': 'AcFun热门',
      'baidu': '百度热搜',
      'bilibili': '哔哩哔哩热门视频',
      'cctv': 'CCTV新闻热点',
      'csdn': 'CSDN热门博客',
      'dongqiudi': '懂球帝热门体育资讯',
      'douban': '豆瓣热门搜索',
      'douyin': '抖音热搜',
      'github': 'GitHub Trending',
      'guojiadili': '国家地理热门文章',
      'historytoday': '历史上的今天',
      'hupu': '虎扑热门体育资讯',
      'ithome': 'IT之家热搜',
      'lishipin': '梨视频热门视频',
      'nanfangzhoumo': '南方周末热门内容',
      'pengpai': '澎湃新闻热点新闻',
      'qqnews': '腾讯新闻热点',
      'quark': '夸克热点搜索',
      'renmin': '人民网热门新闻',
      'shaoshupai': '少数派热门文章',
      'sougou': '搜狗热点搜索',
      'souhu': '搜狐热点新闻',
      'toutiao': '今日头条热搜',
      'v2ex': 'V2EX热议话题',
      'wangyinews': '网易新闻热点',
      'weibo': '微博热搜',
      'xinjingbao': '新京报热点新闻',
      'zhihu': '知乎热搜',
      'all': '所有平台热搜',
      'list': '平台列表'
    }

    const platformTitle = computed(() => {
      return platformMap[props.name] || `${props.name} - 热搜`
    })

    const getTagType = (tag: string): 'primary' | 'success' | 'warning' | 'info' | 'danger' | undefined => {
      if (!tag) return undefined
      
      const tagLower = tag.toLowerCase()
      if (tagLower.includes('热') || tagLower.includes('新') || tagLower.includes('沸')) return 'danger'
      if (tagLower.includes('荐') || tagLower.includes('精')) return 'warning'
      
      return 'info'
    }

    const goBack = () => {
      router.go(-1)
    }

    onMounted(async () => {
      try {
        const response = await axios.get(`/${props.name}`)
        
        // 处理API返回的数据
        let data = response.data
        
        // 检查是否是HTML响应（错误页面）
        if (typeof data === 'string' && data.includes('<!DOCTYPE html>')) {
          platformData.value = [{
            index: 1,
            title: 'API返回了错误页面，可能路径不正确',
            hot: '',
            desc: '请检查API路径',
            url: '',
            tag: 'error'
          }]
          return
        }
        
        // 根据API结构处理数据，如果返回的是数组则直接使用，否则提取其中的数组部分
        if (Array.isArray(data)) {
          platformData.value = data.map((item, index) => ({
            index: index + 1,
            ...item
          }))
        } else if (data.obj && Array.isArray(data.obj)) {
          // 正确的API响应格式是 {code: 200, message: "...", obj: [...]}
          platformData.value = data.obj.map((item: any, index: number) => ({
            index: index + 1,
            ...item
          }))
        } else {
          // 如果是对象形式，尝试提取可能的数组字段
          const arrayKeys = Object.keys(data).filter(key => Array.isArray(data[key]))
          if (arrayKeys.length > 0) {
            platformData.value = data[arrayKeys[0]].map((item: any, index: number) => ({
              index: index + 1,
              ...item
            }))
          } else {
            // 如果没有找到数组字段，将整个对象作为一行显示
            platformData.value = [{
              index: 1,
              title: data.message || JSON.stringify(data),
              hot: '',
              desc: '数据详情',
              url: '',
              tag: ''
            }]
          }
        }
      } catch (err: any) {
        error.value = err.message || '获取数据失败'
        console.error('获取平台数据失败:', err)
      } finally {
        loading.value = false
      }
    })

    return {
      platformData,
      loading,
      error,
      platformTitle,
      getTagType,
      goBack
    }
  }
}
</script>

<style scoped>
.platform-detail {
  max-width: 1200px;
  margin: 0 auto;
}

.content {
  margin-top: 20px;
}

.title-link {
  color: #409EFF;
  text-decoration: none;
}

.title-link:hover {
  text-decoration: underline;
}

/* 暗色模式下链接颜色 */
html.dark .title-link {
  color: var(--el-color-primary);
}

.loading {
  padding: 20px;
}

.error {
  margin-top: 20px;
}
</style>