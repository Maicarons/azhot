package websocket

import (
	"api/service"
	"sync"

	websocket "github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
)

// Client 客户端结构
type Client struct {
	Conn *websocket.Conn
	Type string // api类型，如 baidu, bilibili 等
}

// WsManager WebSocket管理器
type WsManager struct {
	clients          map[*websocket.Conn]*Client
	broadcast        chan Message
	register         chan *Client
	unregister       chan *websocket.Conn
	hotSearchService *service.HotSearchService
	mutex            sync.RWMutex
}

// Message WebSocket消息结构
type Message struct {
	Type   string      `json:"type"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error,omitempty"`
	Source string      `json:"source,omitempty"`
}

// NewWsManager 创建新的WebSocket管理器
func NewWsManager(hotSearchService *service.HotSearchService) *WsManager {
	return &WsManager{
		clients:          make(map[*websocket.Conn]*Client),
		broadcast:        make(chan Message),
		register:         make(chan *Client),
		unregister:       make(chan *websocket.Conn),
		hotSearchService: hotSearchService,
	}
}

// Start 启动WebSocket管理器
func (manager *WsManager) Start() {
	go func() {
		for {
			select {
			case client := <-manager.register:
				manager.mutex.Lock()
				manager.clients[client.Conn] = client
				manager.mutex.Unlock()
				log.Info("客户端连接: ", client.Conn.RemoteAddr())

			case client := <-manager.unregister:
				manager.mutex.Lock()
				if _, ok := manager.clients[client]; ok {
					delete(manager.clients, client)
					client.Close()
				}
				manager.mutex.Unlock()
				log.Info("客户端断开连接: ", client.RemoteAddr())

			case message := <-manager.broadcast:
				manager.mutex.RLock()
				for conn := range manager.clients {
					if err := conn.WriteJSON(message); err != nil {
						log.Error("发送消息失败: ", err)
						delete(manager.clients, conn)
						conn.Close()
					}
				}
				manager.mutex.RUnlock()
			}
		}
	}()
}

// HandleWebSocket 处理WebSocket连接
func (manager *WsManager) HandleWebSocket(c *websocket.Conn) {
	client := &Client{
		Conn: c,
		Type: "default", // 默认类型，后续可通过消息设置
	}

	manager.register <- client

	defer func() {
		manager.unregister <- c
	}()

	for {
		var msg Message
		if err := c.ReadJSON(&msg); err != nil {
			log.Error("读取消息失败: ", err)
			break
		}

		// 根据消息类型处理请求
		switch msg.Type {
		case "subscribe":
			// 订阅特定类型的实时数据
			client.Type = msg.Source
			manager.handleSubscribe(client, msg.Source)
		case "unsubscribe":
			// 取消订阅
			client.Type = "default"
		case "request":
			// 请求一次性数据
			manager.handleRequest(client, msg.Source)
		case "ping":
			// 心跳响应
			response := Message{
				Type: "pong",
				Data: "pong",
			}
			if err := c.WriteJSON(response); err != nil {
				log.Error("发送pong失败: ", err)
				break
			}
		}
	}
}

// handleSubscribe 处理订阅请求
func (manager *WsManager) handleSubscribe(client *Client, source string) {
	// 立即发送当前数据
	manager.handleRequest(client, source)
}

// handleRequest 处理数据请求
func (manager *WsManager) handleRequest(client *Client, source string) {
	var data map[string]interface{}
	var err error

	// 检查特殊处理的源
	switch source {
	case "all":
		data, err = manager.hotSearchService.GetAllFromDBOrFetch()
	case "list":
		routeNames := manager.hotSearchService.GetRouteNames()
		data = map[string]interface{}{
			"code":    200,
			"message": "sources",
			"obj":     routeNames,
		}
	default:
		// 使用映射来处理常规API调用，减少嵌套条件
		apiSourceMap := map[string]string{
			"360search":    "360search",
			"bilibili":     "bilibili",
			"acfun":        "acfun",
			"csdn":         "csdn",
			"dongqiudi":    "dongqiudi",
			"douban":       "douban",
			"douyin":       "douyin",
			"github":       "github",
			"guojiadili":   "guojiadili",
			"historytoday": "historytoday",
			"hupu":         "hupu",
			"ithome":       "ithome",
			"lishipin":     "lishipin",
			"pengpai":      "pengpai",
			"qqnews":       "qqnews",
			"shaoshupai":   "shaoshupai",
			"sougou":       "sougou",
			"souhu":        "souhu",
			"toutiao":      "toutiao",
			"v2ex":         "v2ex",
			"wangyinews":   "wangyinews",
			"weibo":        "weibo",
			"xinjingbao":   "xinjingbao",
			"zhihu":        "zhihu",
			"quark":        "quark",
			"kuake":        "quark", // kuake映射到quark
			"baidu":        "baidu",
			"renmin":       "renmin",
			"nanfang":      "nanfang",
			"360doc":       "360doc",
			"cctv":         "cctv",
		}

		if apiSource, exists := apiSourceMap[source]; exists {
			// 对于"kuake"，我们映射到"quark"，所以需要特殊处理
			if source == "kuake" {
				data, err = manager.hotSearchService.FetchDataFromAPI("quark")
			} else {
				data, err = manager.hotSearchService.FetchDataFromAPI(apiSource)
			}
		} else {
			// 默认返回空结果
			data = map[string]interface{}{
				"code":    200,
				"message": source,
				"obj":     []map[string]interface{}{},
			}
		}
	}

	response := Message{
		Type:   "response",
		Source: source,
		Data:   data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	if client.Conn != nil {
		if err := client.Conn.WriteJSON(response); err != nil {
			log.Error("发送响应失败: ", err)
		}
	}
}
