package main

import (
	"conf"
	"fmt"
)

func main()  {
	conf.InitWithCli([]conf.Flag{
		conf.Flag{
			Key:"port",
			Default:"9090",
			Usage:"port of http ",
		},
		conf.Flag{
			Key:"port2",
			Default:"90901",
			Usage:"port of http2 ",
		},
	})
	fmt.Println(conf.String("port"))
}
