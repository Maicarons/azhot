package app

import (
	"fmt"
)

// PlatformInfo 平台信息结构体
type PlatformInfo struct {
	RouteName string `json:"routeName"` // 路由名称
	Name      string `json:"name"`      // 中文名称
	Icon      string `json:"icon"`      // 图标URL
}

// ListSources 获取所有可用的来源列表
//
//	@Summary		获取所有可用的来源列表
//	@Description	获取所有可用的热搜来源名称列表
//	@Tags			list
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Failure		500	{object}	map[string]interface{}
//	@Router			/list [get]
func ListSources() (map[string]interface{}, error) {
	// 获取所有平台信息
	platforms := GetAllPlatformsInfo()
	
	// 将结构体数组转换为map数组，以确保与JSON序列化后的格式一致
	platformsMap := make([]interface{}, len(platforms))
	for i, platform := range platforms {
		platformsMap[i] = map[string]interface{}{
			"routeName": platform.RouteName,
			"name":      platform.Name,
			"icon":      platform.Icon,
		}
	}

	return map[string]interface{}{
		"code":    200,
		"message": fmt.Sprintf("共 %d 个可用平台", len(platforms)),
		"obj":     platformsMap,
	}, nil
}

// GetAllPlatformsInfo 获取所有平台的详细信息
func GetAllPlatformsInfo() []PlatformInfo {
	return []PlatformInfo{
		{RouteName: "360doc", Name: "360doc", Icon: "https://www.360doc.com/favicon.ico"},
		{RouteName: "360search", Name: "360搜索", Icon: "https://ss.360tres.com/static/121a1737750aa53d.ico"},
		{RouteName: "acfun", Name: "AcFun", Icon: "https://cdn.aixifan.com/ico/favicon.ico"},
		{RouteName: "baidu", Name: "百度", Icon: "https://www.baidu.com/favicon.ico"},
		{RouteName: "bilibili", Name: "哔哩哔哩", Icon: "https://static.hdslb.com/mobile/img/512.png"},
		{RouteName: "cctv", Name: "央视网", Icon: "https://img.cctvpic.com/favicon.ico"},
		{RouteName: "csdn", Name: "CSDN", Icon: "https://g.csdnimg.cn/static/logo/favicon32.ico"},
		{RouteName: "dongqiudi", Name: "懂球帝", Icon: "https://www.dongqiudi.com/favicon.ico"},
		{RouteName: "douban", Name: "豆瓣", Icon: "https://img3.doubanio.com/favicon.ico"},
		{RouteName: "douyin", Name: "抖音", Icon: "https://s16.douyinstatic.com/ies/resource/mandown/ies/aweme_resource/dy_favicon.ico"},
		{RouteName: "github", Name: "GitHub", Icon: "https://github.githubassets.com/favicons/favicon.png"},
		{RouteName: "guojiadili", Name: "国家地理", Icon: "http://www.dili360.com/favicon.ico"},
		{RouteName: "historytoday", Name: "历史上的今天", Icon: "https://www.lssdjt.com/favicon.ico"},
		{RouteName: "hupu", Name: "虎扑", Icon: "https://www.hupu.com/favicon.ico"},
		{RouteName: "ithome", Name: "IT之家", Icon: "https://www.ithome.com/favicon.ico"},
		{RouteName: "lishipin", Name: "梨视频", Icon: "https://page.pearvideo.com/webres/img/logo.png"},
		{RouteName: "nanfang", Name: "南方周末", Icon: "https://www.infzm.com/favicon.ico"},
		{RouteName: "pengpai", Name: "澎湃新闻", Icon: "https://www.thepaper.cn/www/static/images/favicon19x19.ico"},
		{RouteName: "qqnews", Name: "腾讯新闻", Icon: "https://vfiles.gtimg.cn/vupload/20210825/9e3c9d7e9d8e8f8c8f8c8f8c8f8c8f8c.png"},
		{RouteName: "quark", Name: "夸克", Icon: "https://quark.sm.cn/res/static/img/favicon.ico"},
		{RouteName: "renmin", Name: "人民网", Icon: "http://www.people.com.cn/favicon.ico"},
		{RouteName: "sougou", Name: "搜狗", Icon: "https://www.sogou.com/favicon.ico"},
		{RouteName: "souhu", Name: "搜狐", Icon: "https://m.sohu.com/favicon.ico"},
		{RouteName: "toutiao", Name: "今日头条", Icon: "https://sf1-ttcdn-tos.pstatp.com/img/mosaic-legacy/3793/245327~180x180.ico"},
		{RouteName: "v2ex", Name: "V2EX", Icon: "https://www.v2ex.com/static/favicon.ico"},
		{RouteName: "wangyinews", Name: "网易新闻", Icon: "https://news.163.com/favicon.ico"},
		{RouteName: "weibo", Name: "微博", Icon: "https://weibo.com/favicon.ico"},
		{RouteName: "xinjingbao", Name: "新京报", Icon: "https://www.bjnews.com.cn/favicon.ico"},
		{RouteName: "zhihu", Name: "知乎", Icon: "https://static.zhihu.com/static/favicon.ico"},
	}
}

// GetAllRouteNames 获取所有可用的路由名称
func GetAllRouteNames() []string {
	platforms := GetAllPlatformsInfo()
	routeNames := make([]string, len(platforms))
	for i, platform := range platforms {
		routeNames[i] = platform.RouteName
	}
	return routeNames
}