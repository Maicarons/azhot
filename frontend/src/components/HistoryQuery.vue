<template>
  <div class="history-query">
    <StarPrompt @close="() => {}" />
    <h2>历史记录查询</h2>
    
    <el-form :model="form" label-width="120px" class="query-form">
      <el-form-item label="平台选择">
        <el-select v-model="form.platform" placeholder="请选择要查询的平台" style="width: 100%">
          <el-option
            v-for="platform in platforms"
            :key="platform.name"
            :label="platform.title"
            :value="platform.name"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="查询方式">
        <el-radio-group v-model="form.queryType">
          <el-radio label="byDate">按日期查询</el-radio>
          <el-radio label="byDateTime">按日期和小时查询</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="选择日期">
        <el-date-picker
          v-model="form.date"
          type="date"
          placeholder="选择日期"
          format="YYYY-MM-DD"
          value-format="YYYY-MM-DD"
        />
      </el-form-item>

      <el-form-item v-if="form.queryType === 'byDateTime'" label="选择小时">
        <el-select v-model="form.hour" placeholder="选择小时">
          <el-option
            v-for="n in 24"
            :key="n"
            :label="String(n-1).padStart(2, '0') + ':00'"
            :value="n-1"
          />
        </el-select>
      </el-form-item>

      <el-form-item>
        <el-button 
          type="primary" 
          @click="onSubmit" 
          :loading="loading"
          :disabled="!form.platform || !form.date"
        >
          查询
        </el-button>
      </el-form-item>
    </el-form>

    <div v-if="result" class="result-container">
      <h3>查询结果</h3>
      <div v-if="Array.isArray(result)" class="result-list">
        <el-table 
          :data="result" 
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
        </el-table>
      </div>
      <div v-else-if="typeof result === 'object'" class="result-by-date">
        <div v-for="(items, hour) in result" :key="hour" class="hour-section">
          <h4>{{ hour }}</h4>
          <el-table 
            :data="items" 
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
          </el-table>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import StarPrompt from './StarPrompt.vue'

// 定义平台类型
interface Platform {
  name: string;
  title: string;
  description: string;
}

// 定义表单类型
interface FormModel {
  platform: string;
  queryType: 'byDate' | 'byDateTime';
  date: string;
  hour: number | null;
}

// 定义历史数据项类型
interface HistoryItem {
  index: number;
  title: string;
  url?: string;
  [key: string]: any;
}

export default {
  name: 'HistoryQuery',
  components: {
    StarPrompt
  },
  setup() {
    const form = ref<FormModel>({
      platform: '',
      queryType: 'byDate', // 'byDate' or 'byDateTime'
      date: '',
      hour: null
    })
    
    const platforms = ref<Platform[]>([])
    const result = ref<HistoryItem[] | Record<string, HistoryItem[]> | null>(null)
    const loading = ref(false)
    
    // 平台信息映射
    const platformMap: Record<string, { title: string; description: string }> = {
      '360doc': { title: '360doc', description: '获取360doc热门文章列表' },
      '360search': { title: '360搜索', description: '获取360搜索热点排行榜' },
      'acfun': { title: 'AcFun', description: '获取AcFun热门排行榜' },
      'baidu': { title: '百度', description: '获取百度热搜列表' },
      'bilibili': { title: '哔哩哔哩', description: '获取哔哩哔哩热门视频列表' },
      'cctv': { title: 'CCTV新闻', description: '获取CCTV新闻热点排行榜' },
      'csdn': { title: 'CSDN', description: '获取CSDN热门博客文章排行榜' },
      'dongqiudi': { title: '懂球帝', description: '获取懂球帝热门体育资讯排行榜' },
      'douban': { title: '豆瓣', description: '获取豆瓣热门搜索列表' },
      'douyin': { title: '抖音', description: '获取抖音热搜列表' },
      'github': { title: 'GitHub', description: '获取GitHub Trending列表' },
      'guojiadili': { title: '国家地理', description: '获取国家地理热门文章排行榜' },
      'historytoday': { title: '历史上的今天', description: '获取历史上的今天事件列表' },
      'hupu': { title: '虎扑', description: '获取虎扑热门体育资讯排行榜' },
      'ithome': { title: 'IT之家', description: '获取IT之家热搜列表' },
      'lishipin': { title: '梨视频', description: '获取梨视频热门视频排行榜' },
      'nanfangzhoumo': { title: '南方周末', description: '获取南方周末热门内容排行榜' },
      'pengpai': { title: '澎湃新闻', description: '获取澎湃新闻热点新闻排行榜' },
      'qqnews': { title: '腾讯新闻', description: '获取腾讯新闻热点排行榜' },
      'quark': { title: '夸克', description: '获取夸克热点搜索排行榜' },
      'renmin': { title: '人民网', description: '获取人民网热门新闻排行榜' },
      'shaoshupai': { title: '少数派', description: '获取少数派热门文章排行榜' },
      'sougou': { title: '搜狗', description: '获取搜狗热点搜索排行榜' },
      'souhu': { title: '搜狐', description: '获取搜狐热点新闻排行榜' },
      'toutiao': { title: '今日头条', description: '获取今日头条热搜列表' },
      'v2ex': { title: 'V2EX', description: '获取V2EX热议话题列表' },
      'wangyinews': { title: '网易新闻', description: '获取网易新闻热点排行榜' },
      'weibo': { title: '微博', description: '获取微博热搜列表' },
      'xinjingbao': { title: '新京报', description: '获取新京报热点新闻排行榜' },
      'zhihu': { title: '知乎', description: '获取知乎热搜列表' },
      'all': { title: '全部平台', description: '获取所有平台的热搜列表' },
      'list': { title: '平台列表', description: '获取所有可用的热搜来源名称列表' }
    }

    // 初始化平台列表
    platforms.value = Object.entries(platformMap).map(([name, info]) => ({
      name,
      ...info
    })).filter(p => p.title) // 过滤掉没有映射信息的平台

    const onSubmit = async () => {
      if (!form.value.platform || !form.value.date) {
        return
      }

      loading.value = true
      result.value = null

      try {
        let url = ''
        if (form.value.queryType === 'byDateTime' && form.value.hour !== null) {
          // 查询特定日期和小时的数据
          url = `/history/${form.value.platform}/${form.value.date}/${form.value.hour}`
        } else {
          // 查询特定日期的所有小时数据
          url = `/history/${form.value.platform}/${form.value.date}`
        }

        const response = await axios.get(url)
        
        if (response.data && response.data.obj) {
          result.value = response.data.obj
        } else {
          result.value = []
        }
      } catch (error) {
        console.error('查询历史记录失败:', error)
        result.value = null
      } finally {
        loading.value = false
      }
    }

    return {
      form,
      platforms,
      result,
      loading,
      onSubmit
    }
  }
}
</script>

<style scoped>
.history-query {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.query-form {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
}

/* 暗色模式下表单样式 */
html.dark .query-form {
  background: var(--el-bg-color-overlay);
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.3);
}

.result-container {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

/* 暗色模式下结果容器样式 */
html.dark .result-container {
  background: var(--el-bg-color-overlay);
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.3);
}

.result-list {
  margin-top: 20px;
}

.hour-section {
  margin-bottom: 30px;
}

.hour-section h4 {
  border-bottom: 1px solid #eee;
  padding-bottom: 5px;
  margin-bottom: 15px;
}

/* 暗色模式下小时段标题样式 */
html.dark .hour-section h4 {
  border-bottom: 1px solid var(--el-border-color);
  color: var(--el-text-color-primary);
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
</style>