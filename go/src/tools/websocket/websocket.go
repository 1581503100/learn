package websocket

import (
	"net"
	"log"
	"strings"
	"crypto/sha1"
	"io"
	"encoding/base64"
)

func Server()  {
	ln,err:=net.Listen("tcp","9090")
	if(err !=nil){
		panic(err)
	}

	for{
		con,err:=ln.Accept()
		if(err !=nil){
			log.Println(err)
			continue
		}
		go handlerConnection(con)
	}
}
func handlerConnection(conn net.Conn)  {
	content:=make([]byte,1024)
	_,err:=conn.Read(content)
	if(err !=nil){
		log.Println(err)
	}
	isHttp:=false
	if(string(content[0:3])=="GET"){
		isHttp =true
	}
	if isHttp{
		headers:=parseHandshake(string(content))
		websocketKey:=headers["Sec-WebSocket-Key"]
		guid:="258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
		h:=sha1.New()
		log.Println("accept raw:"+websocketKey+guid)
		io.WriteString(h,websocketKey+guid)
		accept:=make([]byte,28)
		base64.StdEncoding.Encode(accept,h.Sum(nil))
		log.Println(string(accept))

		response:="HTTP/1.1 101 Switching Protocols\r\n"
		response = response+"Sec-WebSocket-Accept: "+string(accept)+"\r\n"
		response+="Connection: Upgrade\r\n"
		response+= "Upgrade: WebSocket\r\n\r\n"

		log.Println(response)

		ws:=NewWsSocket(conn)
		for{
			data,err:=ws.ReadFrame()
		}
	}
}

type WsSocker struct {
	MaskingKey []byte
	Conn net.Conn
}

func NewWsSocket(conn net.Conn) *WsSocker {
	return &WsSocker{Conn:conn}
}

func (this *WsSocker)ReadFrame()(data []byte,err error)  {
	err = nil
	opcodeByte:=make([]byte,1)
	this.Conn.Read(opcodeByte)
	FIN:=opcodeByte[0]>>7
	RSV1:=opcodeByte[0]>>6&1
	RSV2:=opcodeByte[0]>>5 & 1
	RSV3:=opcodeByte[0]>>4 & 1
	OPCODE:=opcodeByte[0] & 15

	log.Println(RSV1,RSV2,RSV3,OPCODE)
	playloadLenByte:=make([]byte,1)
	this.Conn.Read(playloadLenByte)
	playloadLen:=int(playloadLenByte[0] & 0x7f)
	mask:=playloadLenByte[0] >>7
	if (playloadLen==127){
		extendeByte:=make([]byte,4)
		this.Conn.Read(extendeByte)
	}

	maskingByte:=make([]byte,4)
	if(mask == 1){
		this.Conn.Read(maskingByte)
	}
	payloadDataByte := make([]byte, playloadLen)
	this.Conn.Read(payloadDataByte)
	log.Println("data:", payloadDataByte)

	dataByte:=make([]byte,playloadLen)
	for i:=0;i<playloadLen;i++{
		if mask ==1 {
			dataByte[i]=payloadDataByte[i]^maskingByte[i%4]
		}else {
			dataByte[i] = payloadDataByte[i]
		}
	}
	if FIN == 1{
		data = dataByte
	}
	nextData,err:=this.ReadFrame()
	if(err !=nil){
		return
	}
	data = append(data,nextData...)
	return
}

func parseHandshake(content string) map[string]string {
	headers := make(map[string]string, 10)
	lines := strings.Split(content, "\r\n")

	for _,line := range lines {
		if len(line) >= 0 {
			words := strings.Split(line, ":")
			if len(words) == 2 {
				headers[strings.Trim(words[0]," ")] = strings.Trim(words[1], " ")
			}
		}
	}
	return headers
}