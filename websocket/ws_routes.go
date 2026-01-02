package websocket

import (
	"api/config"
	"api/service"

	websocket "github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupWebSocketRoutes 配置WebSocket路由
func SetupWebSocketRoutes(app *fiber.App, hotSearchService *service.HotSearchService, cfg *config.Config) {
	// 使用CORS中间件
	app.Use(cors.New())
	// 使用日志中间件
	app.Use(logger.New())

	// 创建WebSocket管理器
	wsManager := NewWsManager(hotSearchService)

	// 启动WebSocket管理器
	wsManager.Start()

	// WebSocket路由
	app.Get("/ws", func(c *fiber.Ctx) error {
		// 检查是否为WebSocket请求
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return websocket.New(wsManager.HandleWebSocket)(c)
		}
		return fiber.ErrUpgradeRequired
	})

	// 为每个API端点创建对应的WebSocket订阅端点
	setupWebSocketAPIRoutes(app, wsManager)
}

// setupWebSocketAPIRoutes 设置WebSocket API路由
func setupWebSocketAPIRoutes(app *fiber.App, wsManager *WsManager) {
	// 为每个API端点创建WebSocket路由
	apiEndpoints := []string{
		"/baidu", "/bilibili", "/list", "/360search", "/acfun", "/csdn",
		"/dongqiudi", "/douban", "/douyin", "/github", "/guojiadili",
		"/historytoday", "/hupu", "/ithome", "/lishipin", "/pengpai",
		"/qqnews", "/sougou", "/souhu", "/toutiao", "/v2ex", "/wangyinews",
		"/weibo", "/xinjingbao", "/zhihu", "/renmin", "/nanfang", "/360doc",
		"/cctv", "/quark", "/all",
	}

	for _, endpoint := range apiEndpoints {
		// 提取源名称（去掉开头的斜杠）
		source := endpoint[1:]

		// 创建WebSocket路由，格式为 /ws/{source}
		wsPath := "/ws" + endpoint
		app.Get(wsPath, func(source string) fiber.Handler {
			return func(c *fiber.Ctx) error {
				if websocket.IsWebSocketUpgrade(c) {
					c.Locals("allowed", true)
					// 创建一个包装函数来处理特定源的WebSocket连接
					return websocket.New(func(conn *websocket.Conn) {
						client := &Client{
							Conn: conn,
							Type: source,
						}

						wsManager.register <- client

						defer func() {
							wsManager.unregister <- conn
						}()

						// 立即发送当前数据
						wsManager.handleRequest(client, source)

						// 等待客户端断开连接
						for {
							var msg Message
							if err := conn.ReadJSON(&msg); err != nil {
								break
							}
							// 只处理心跳消息，其他消息不做处理
							if msg.Type == "ping" {
								response := Message{
									Type: "pong",
									Data: "pong",
								}
								if err := conn.WriteJSON(response); err != nil {
									break
								}
							}
						}
					})(c)
				}
				return fiber.ErrUpgradeRequired
			}
		}(source))
	}

	// 历史数据WebSocket路由
	setupWebSocketHistoryRoutes(app, wsManager)
}

// setupWebSocketHistoryRoutes 设置WebSocket历史数据路由
func setupWebSocketHistoryRoutes(app *fiber.App, wsManager *WsManager) {
	// 历史数据路由模式
	historyEndpoints := []string{
		"/history/:source/:date/:hour",
		"/history/:source/:date",
		"/history/:source",
	}

	for _, endpoint := range historyEndpoints {
		// 为历史数据路由创建对应的WebSocket路由
		wsPath := "/ws" + endpoint
		app.Get(wsPath, func(endpoint string) fiber.Handler {
			return func(c *fiber.Ctx) error {
				if websocket.IsWebSocketUpgrade(c) {
					c.Locals("allowed", true)
					source := c.Params("source")

					// 创建一个包装函数来处理历史数据的WebSocket连接
					return websocket.New(func(conn *websocket.Conn) {
						client := &Client{
							Conn: conn,
							Type: "history",
						}

						wsManager.register <- client

						defer func() {
							wsManager.unregister <- conn
						}()

						// 构建历史数据请求参数
						date := c.Params("date")
						hour := c.Params("hour")
						var historySource string

						switch {
						case hour != "":
							// /ws/history/:source/:date/:hour
							historySource = "history_" + source + "_" + date + "_" + hour
						case date != "":
							// /ws/history/:source/:date
							historySource = "history_" + source + "_" + date
						default:
							// /ws/history/:source
							historySource = "history_" + source
						}

						// 这里需要特殊处理历史数据请求
						// 由于历史数据需要通过服务层获取，我们直接调用服务层方法
						var data map[string]interface{}
						var err error

						switch {
						case hour != "":
							// 调用历史数据处理器
							data, err = wsManager.hotSearchService.GetHistoricalDataForWS(source, date, hour)
						case date != "":
							// 调用按日期历史数据处理器
							data, err = wsManager.hotSearchService.GetHistoricalDataByDateForWS(source, date)
						default:
							// 调用按来源历史数据处理器
							data, err = wsManager.hotSearchService.GetHistoricalDataBySourceForWS(source)
						}

						response := Message{
							Type:   "response",
							Source: historySource,
							Data:   data,
						}

						if err != nil {
							response.Error = err.Error()
						}

						if err := conn.WriteJSON(response); err != nil {
							return
						}

						// 等待客户端断开连接
						for {
							var msg Message
							if err := conn.ReadJSON(&msg); err != nil {
								break
							}
							// 只处理心跳消息
							if msg.Type == "ping" {
								response := Message{
									Type: "pong",
									Data: "pong",
								}
								if err := conn.WriteJSON(response); err != nil {
									break
								}
							}
						}
					})(c)
				}
				return fiber.ErrUpgradeRequired
			}
		}(endpoint))
	}
}
