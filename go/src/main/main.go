package main

import (
	"conf"
	"fmt"
)
type (
	S struct {
		Username string
		Password string
		Port int
	}
)
func main()  {

	conf.InitWithCli([]conf.Flag{
		conf.Flag{
			Key:"mysql.username",
			Default:"9090",
			Usage:"port of http ",
		},
		conf.Flag{
			Key:"mysql.password",
			Default:"90901",
			Usage:"port of http2 ",
		},
	})
	conf.Init("..\\conf\\app.conf")
	var s S
	conf.Unmarshal(&s,"mysql.")
	fmt.Println(s)
	fmt.Println(conf.String("mysql.password"))
}
