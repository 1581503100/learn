package tests

import (
	"conf"
	"fmt"
	"testing"
)

type A struct {
	Port          int
	MysqlName     string
	MysqlPassword string
}

func Test(t *testing.T) {
	//
	fmt.Println()
	conf.Init("../conf/app.conf")
	var a A
	conf.Unmarshal(&a, "mysql.")
	fmt.Println(a)
}

func TestCLi(t *testing.T)  {

	conf.InitWithCli([]conf.Flag{
		conf.Flag{
			Key:"port",
			Default:"8080",
			Usage:"listen port",
		},
		conf.Flag{
			Key:"name",
			Default:"Josn",
			Usage:"author",
		},
	})
	fmt.Println(conf.String("port"))
	fmt.Println(conf.String("name"))
}
