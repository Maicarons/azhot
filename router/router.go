package router

import (
	"api/all"
	app_pkg "api/app"
	"api/config"
	"api/service"
	"api/websocket"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger" // swagger handler
)

// SetupRoutes 配置所有路由
func SetupRoutes(app *fiber.App, hotSearchService *service.HotSearchService, cfg *config.Config) {
	// 使用CORS中间件
	app.Use(cors.New())
	// 使用日志中间件
	app.Use(logger.New())

	if !cfg.Debug {
		app.Use(cache.New())

		app.Use(etag.New())

		app.Use(favicon.New())

		app.Use(idempotency.New())

		app.Use(limiter.New(limiter.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.IP() == "127.0.0.1"
			},
			Max:               20,
			Expiration:        30 * time.Second,
			LimiterMiddleware: limiter.SlidingWindow{},
		}))
	}

	// 仅在调试模式下启用 metrics 路由
	if cfg.Debug {
		app.Get("/metrics", monitor.New())
	}

	// Swagger API文档
	app.Get("/swagger/*", swagger.HandlerDefault)

	// 设置API路由
	setupAPIRoutes(app, hotSearchService)
	
	// 设置WebSocket路由
	websocket.SetupWebSocketRoutes(app, hotSearchService, cfg)
}

// setupAPIRoutes 设置API路由
func setupAPIRoutes(app *fiber.App, hotSearchService *service.HotSearchService) {
	// 根路径重定向到 Swagger UI
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html")
	})

	// 实时API - 直接请求接口
	app.Get("/baidu", createHandler(func() (interface{}, error) {
		return app_pkg.Baidu()
	}))
	app.Get("/bilibili", createHandler(func() (interface{}, error) {
		return app_pkg.Bilibili()
	}))
	app.Get("/list", createHandler(func() (interface{}, error) {
		return app_pkg.ListSources()
	}))
	app.Get("/360search", createHandler(func() (interface{}, error) {
		return app_pkg.Search360()
	}))
	app.Get("/acfun", createHandler(func() (interface{}, error) {
		return app_pkg.Acfun()
	}))
	app.Get("/csdn", createHandler(func() (interface{}, error) {
		return app_pkg.CSDN()
	}))
	app.Get("/dongqiudi", createHandler(func() (interface{}, error) {
		return app_pkg.Dongqiudi()
	}))
	app.Get("/douban", createHandler(func() (interface{}, error) {
		return app_pkg.Douban()
	}))
	app.Get("/douyin", createHandler(func() (interface{}, error) {
		return app_pkg.Douyin()
	}))
	app.Get("/github", createHandler(func() (interface{}, error) {
		return app_pkg.Github()
	}))
	app.Get("/guojiadili", createHandler(func() (interface{}, error) {
		return app_pkg.Guojiadili()
	}))
	// 统一历史路由为 historytoday，避免与历史记录查询冲突
	app.Get("/historytoday", createHandler(func() (interface{}, error) {
		return app_pkg.HistoryToday()
	}))
	app.Get("/hupu", createHandler(func() (interface{}, error) {
		return app_pkg.Hupu()
	}))
	app.Get("/ithome", createHandler(func() (interface{}, error) {
		return app_pkg.Ithome()
	}))
	app.Get("/lishipin", createHandler(func() (interface{}, error) {
		return app_pkg.Lishipin()
	}))
	app.Get("/pengpai", createHandler(func() (interface{}, error) {
		return app_pkg.Pengpai()
	}))
	app.Get("/qqnews", createHandler(func() (interface{}, error) {
		return app_pkg.Qqnews()
	}))
	app.Get("/sougou", createHandler(func() (interface{}, error) {
		return app_pkg.Sougou()
	}))
	app.Get("/souhu", createHandler(func() (interface{}, error) {
		return app_pkg.Souhu()
	}))
	app.Get("/toutiao", createHandler(func() (interface{}, error) {
		return app_pkg.Toutiao()
	}))
	app.Get("/v2ex", createHandler(func() (interface{}, error) {
		return app_pkg.V2ex()
	}))
	app.Get("/wangyinews", createHandler(func() (interface{}, error) {
		return app_pkg.WangyiNews()
	}))
	app.Get("/weibo", createHandler(func() (interface{}, error) {
		return app_pkg.WeiboHot()
	}))
	app.Get("/xinjingbao", createHandler(func() (interface{}, error) {
		return app_pkg.Xinjingbao()
	}))
	app.Get("/zhihu", createHandler(func() (interface{}, error) {
		return app_pkg.Zhihu()
	}))
	app.Get("/renmin", createHandler(func() (interface{}, error) {
		return app_pkg.Renminwang()
	}))
	app.Get("/nanfang", createHandler(func() (interface{}, error) {
		return app_pkg.Nanfangzhoumo()
	}))
	app.Get("/360doc", createHandler(func() (interface{}, error) {
		return app_pkg.Doc360()
	}))
	app.Get("/cctv", createHandler(func() (interface{}, error) {
		return app_pkg.CCTV()
	}))
	app.Get("/quark", createHandler(func() (interface{}, error) {
		return app_pkg.Quark()
	}))

	// 聚合API
	app.Get("/all", createHandler(func() (interface{}, error) {
		return all.All(), nil
	}))

	// 历史API - 保留历史记录查询
	// 获取指定平台、日期和小时的历史数据
	app.Get("/history/:source/:date/:hour", func(c *fiber.Ctx) error {
		return hotSearchService.GetHistoricalDataHandler(c)
	})

	// 获取指定平台、日期的所有小时数据
	app.Get("/history/:source/:date", func(c *fiber.Ctx) error {
		return hotSearchService.GetHistoricalDataByDateHandler(c)
	})

	// 获取指定平台的所有历史数据
	app.Get("/history/:source", func(c *fiber.Ctx) error {
		return hotSearchService.GetHistoricalDataBySourceHandler(c)
	})
}

// createHandler 创建处理器函数
func createHandler(f func() (interface{}, error)) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data, err := f()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"code":    500,
				"message": err.Error(),
			})
		}
		return c.JSON(data)
	}
}

// createHandlerWithCache 使用缓存创建处理器
func createHandlerWithCache(service *service.HotSearchService, source string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data, err := service.GetFromDBOrFetch(source)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"code":    500,
				"message": err.Error(),
			})
		}
		return c.JSON(data)
	}
}
