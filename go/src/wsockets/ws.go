package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"os"
	"fmt"
)

var wsm = WsMannager{Clients:map[string]*Client{}}

type (
	WsMannager struct {
		Clients map[string]*Client
	}

	Client struct {
		conn *websocket.Conn
		ws WsMannager
	}
	Message struct {
		From string
		To string
		Msg string
	}
)

func main()  {

	http.HandleFunc("/ws",wsPage)
	http.HandleFunc("/i",page)
	http.ListenAndServe(":12345",nil)
}


func(this * WsMannager)Add (client *Client)  {
	fmt.Println(client.conn.RemoteAddr().String())
	this.Clients[client.conn.RemoteAddr().String()]=client
}

func (c *Client)read()  {
	defer func() {
			c.conn.Close()
			addr:=c.conn.RemoteAddr().String()
			delete(c.ws.Clients,addr)
			fmt.Println(addr ,"offline-")
	}()
	for{
		var msg Message
		err:=c.conn.ReadJSON(&msg)
		if(err !=nil){
			c.conn.Close()
			break
		}
		c.sendAll(msg)
	}
}

func (c* Client)write(msg Message)  {
	c.conn.WriteJSON(msg)
}
func (c *Client)sendAll(msg Message)  {
	for _,v:=range c.ws.Clients{
		v.write(msg)
	}
}


func wsPage(w http.ResponseWriter,r *http.Request)  {
	conn,err:=(&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}).Upgrade(w,r,nil)
	if(err!=nil){
		http.NotFound(w,r)
		return
	}
	c:=&Client{conn:conn,ws:wsm}
	wsm.Add(c)
	c.read()
}

func page(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("index")
	f,err:=os.Open("D:\\learn\\learn\\go\\src\\wsockets\\index.html")
	fmt.Println(err)
	defer f.Close()
	bs,_:=ioutil.ReadAll(f)
	w.Write(bs)

}