package main

import (
	"fmt"
	"regexp"
)
type (
	S struct {
		Username string
		Password string
		Port int
	}
)
func main()  {

	rex:=regexp.MustCompile("[a-z][a-z0-9_]*")

	fmt.Println(rex.FindAllStringSubmatch("1sdfddsf",-1))
	fmt.Println(rex.Match([]byte("1sd")))


	//conf.InitWithCli([]conf.Flag{
	//	conf.Flag{
	//		Key:"mysql.username",
	//		Default:"9090",
	//		Usage:"port of http ",
	//	},
	//	conf.Flag{
	//		Key:"mysql.password",
	//		Default:"90901",
	//		Usage:"port of http2 ",
	//	},
	//})
	//conf.Init("..\\conf\\app.conf")
	//var s S
	//conf.Unmarshal(&s,"mysql.")
	//fmt.Println(s)
	//fmt.Println(conf.String("mysql.password"))
}
