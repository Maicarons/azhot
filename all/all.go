package all

import (
	"api/app"
	"bytes"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2/log"
)

// 定义所有可用的来源映射，避免循环导入
var sourceFuncMap = map[string]func() (map[string]interface{}, error){
	"360search":    app.Search360,
	"bilibili":     app.Bilibili,
	"acfun":        app.Acfun,
	"csdn":         app.CSDN,
	"dongqiudi":    app.Dongqiudi,
	"douban":       app.Douban,
	"douyin":       app.Douyin,
	"github":       app.Github,
	"guojiadili":   app.Guojiadili,
	"historytoday": app.HistoryToday, // 更新函数名和键名
	"hupu":         app.Hupu,
	"ithome":       app.Ithome,
	"lishipin":     app.Lishipin,
	"pengpai":      app.Pengpai,
	"qqnews":       app.Qqnews,
	"shaoshupai":   app.Shaoshupai,
	"sougou":       app.Sougou,
	"toutiao":      app.Toutiao,
	"v2ex":         app.V2ex,
	"wangyinews":   app.WangyiNews,
	"weibo":        app.WeiboHot,
	"xinjingbao":   app.Xinjingbao,
	"zhihu":        app.Zhihu,
	"quark":        app.Quark,
	"souhu":        app.Souhu,
	"baidu":        app.Baidu,
	"renmin":       app.Renminwang,
	"nanfang":      app.Nanfangzhoumo,
	"360doc":       app.Doc360,
	"cctv":         app.CCTV,
}

// All 获取所有平台热搜数据
//
//	@Summary		获取所有平台热搜数据
//	@Description	获取所有平台的热搜列表
//	@Tags			all
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Failure		500	{object}	map[string]interface{}
//	@Router			/all [get]
func All() map[string]interface{} {
	allResult := make(map[string]interface{})
	var wg sync.WaitGroup
	var mu sync.Mutex

	for key, fn := range sourceFuncMap {
		wg.Add(1)
		go func(k string, f func() (map[string]interface{}, error)) {
			defer wg.Done()
			result, err := f()
			if err != nil {
				var buf bytes.Buffer
				buf.WriteString(k)
				buf.WriteString(" 请求失败: ")
				buf.WriteString(err.Error())
				log.Error(buf.String())
				return
			}

			if result["code"] == 200 {
				mu.Lock()
				allResult[k] = result["obj"]
				mu.Unlock()
			}
		}(key, fn)
	}

	wg.Wait()
	var buf bytes.Buffer
	buf.WriteString("成功获取的热搜数量: ")
	buf.WriteString(strconv.Itoa(len(allResult)))
	log.Info(buf.String())

	return map[string]interface{}{
		"code": 200,
		"obj":  allResult,
	}
}

// GetAllSourceNames 获取所有可用的来源名称列表
func GetAllSourceNames() []string {
	var sources []string
	for key := range sourceFuncMap {
		sources = append(sources, key)
	}

	return sources
}
