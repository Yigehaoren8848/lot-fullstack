package ms

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	controller "meilian/controllers/base"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 在这里检查请求的来源是否在允许的列表中
		// 返回 true 表示允许连接，返回 false 表示拒绝连接
		return true
	},
}

type WebSocketController struct {
	controller.BaseController
}

// 用于存储用户ID与其对应连接的映射
var userConnections = make(map[string]*websocket.Conn)

func (c *WebSocketController) Get() {

	conn, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		beego.Error("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	// 获取用户标识符，这里假设从请求头中获取
	var userID string

	// 存储连接与用户ID的对应关系

	for {
		// Read message from the client
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			beego.Error("Failed to read message:", err)
			break
		}

		var msgJson map[string]interface{}
		json.Unmarshal(msgBytes, &msgJson)
		msgType := msgJson["type"].(string)

		// Handle different message types
		switch msgType {
		case "auth":
			// Respond with "heart" for heartbeat detection

			userID := msgJson["uid"].(string)
			err = conn.WriteMessage(websocket.TextMessage, []byte("aauth："+userID))
			userConnections[userID] = conn
			if err != nil {
				beego.Error("Failed to write message:", err)
				break
			}
		case "heart":
			// Respond with "heart" for heartbeat detection
			err = conn.WriteMessage(websocket.TextMessage, []byte("heart"))
			if err != nil {
				beego.Error("Failed to write message:", err)
				break
			}
		case "private_message":
			// Send a private message to other users
			senderID := userID
			recipientID := msgJson["receiver"].(string) // Change this to the recipient's userID
			message := msgJson["msg"].(string)

			recipientConn, found := userConnections[recipientID]
			if found {
				err := recipientConn.WriteMessage(websocket.TextMessage, []byte(senderID+": "+message))
				if err != nil {
					beego.Error("Failed to send private message:", err)
				}
			}
		}
	}
}
