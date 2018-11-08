package websocket

import (
	"github.com/gorilla/websocket"
	"encoding/json"
)

type (
	ClientManager struct {
		clients map[*Client]bool
		broadcast chan []byte
		register chan *Client
		unregister chan *Client
	}

	Client struct {
		id string
		socket websocket.Conn
		send chan []byte
	}

	Message struct {
		Sender string `json:"sender"`
		Recipient string `json:"recipient"`
		Content string `json:"content"`
	}

)
var (
	manager = ClientManager{
		broadcast:make(chan []byte),
		register:make(chan *Client),
		unregister:make(chan *Client),
		clients:map[*Client]bool{},
	}
)

func (manager *ClientManager)start()  {
	for{
		select {
			case conn:=<-manager.register:
				manager.clients[conn]=true
				json_message,_:=json.Marshal(&Message{Content:"new socket has contected"})
				manager.send(json_message,conn)
		}
	}

}
func (this *ClientManager)send(msg []byte,ignore *Client)  {
	for conn:=range manager.clients{
		if conn != ignore{
			conn.send <-msg
		}
	}
}
func Serve()  {

	c:=websocket.Conn{}
}
