package main

import (
	"conf"
	"fmt"
)

type (
	F struct {
		Name string
	}
)

func main()  {
	conf.InitWithCli([]conf.Flag{
		conf.Flag{
			Key:"port",
		},
	})
	fmt.Println(conf.String("port"))
}
