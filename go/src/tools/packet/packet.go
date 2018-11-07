package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main()  {
	if handler,err:=pcap.OpenLive("eth0",1600,true,pcap.BlockForever);err!=nil{
		err = handler.SetBPFFilter("port 80")
		if err !=nil{
			panic(err)
		}
		source:=gopacket.NewPacketSource(handler,handler.LinkType())
		for v := range source.Packets() {
			//判断数据包是否是Payload如果是则打印,
			if payload := v.Layer(gopacket.LayerTypePayload); payload != nil {
				fmt.Println(string(payload.LayerContents()))
			}
		}


	}
}