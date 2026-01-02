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

	switch source {
	case "360search":
		data, err = manager.hotSearchService.FetchDataFromAPI("360search")
	case "bilibili":
		data, err = manager.hotSearchService.FetchDataFromAPI("bilibili")
	case "acfun":
		data, err = manager.hotSearchService.FetchDataFromAPI("acfun")
	case "csdn":
		data, err = manager.hotSearchService.FetchDataFromAPI("csdn")
	case "dongqiudi":
		data, err = manager.hotSearchService.FetchDataFromAPI("dongqiudi")
	case "douban":
		data, err = manager.hotSearchService.FetchDataFromAPI("douban")
	case "douyin":
		data, err = manager.hotSearchService.FetchDataFromAPI("douyin")
	case "github":
		data, err = manager.hotSearchService.FetchDataFromAPI("github")
	case "guojiadili":
		data, err = manager.hotSearchService.FetchDataFromAPI("guojiadili")
	case "historytoday":
		data, err = manager.hotSearchService.FetchDataFromAPI("historytoday")
	case "hupu":
		data, err = manager.hotSearchService.FetchDataFromAPI("hupu")
	case "ithome":
		data, err = manager.hotSearchService.FetchDataFromAPI("ithome")
	case "lishipin":
		data, err = manager.hotSearchService.FetchDataFromAPI("lishipin")
	case "pengpai":
		data, err = manager.hotSearchService.FetchDataFromAPI("pengpai")
	case "qqnews":
		data, err = manager.hotSearchService.FetchDataFromAPI("qqnews")
	case "shaoshupai":
		data, err = manager.hotSearchService.FetchDataFromAPI("shaoshupai")
	case "sougou":
		data, err = manager.hotSearchService.FetchDataFromAPI("sougou")
	case "souhu":
		data, err = manager.hotSearchService.FetchDataFromAPI("souhu")
	case "toutiao":
		data, err = manager.hotSearchService.FetchDataFromAPI("toutiao")
	case "v2ex":
		data, err = manager.hotSearchService.FetchDataFromAPI("v2ex")
	case "wangyinews":
		data, err = manager.hotSearchService.FetchDataFromAPI("wangyinews")
	case "weibo":
		data, err = manager.hotSearchService.FetchDataFromAPI("weibo")
	case "xinjingbao":
		data, err = manager.hotSearchService.FetchDataFromAPI("xinjingbao")
	case "zhihu":
		data, err = manager.hotSearchService.FetchDataFromAPI("zhihu")
	case "quark", "kuake":
		data, err = manager.hotSearchService.FetchDataFromAPI("quark")
	case "baidu":
		data, err = manager.hotSearchService.FetchDataFromAPI("baidu")
	case "renmin":
		data, err = manager.hotSearchService.FetchDataFromAPI("renmin")
	case "nanfang":
		data, err = manager.hotSearchService.FetchDataFromAPI("nanfang")
	case "360doc":
		data, err = manager.hotSearchService.FetchDataFromAPI("360doc")
	case "cctv":
		data, err = manager.hotSearchService.FetchDataFromAPI("cctv")
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
		// 默认返回空结果
		data = map[string]interface{}{
			"code":    200,
			"message": source,
			"obj":     []map[string]interface{}{},
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

	if err := client.Conn.WriteJSON(response); err != nil {
		log.Error("发送响应失败: ", err)
	}
}
