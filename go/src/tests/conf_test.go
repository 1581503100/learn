package tests

import (
	"conf"
	"fmt"
	"testing"
	"reflect"
)

type B struct{
	Text string
}

type A struct {
	Port          int
	MysqlName     string
	MysqlPassword string
	Assign **B
}

func Test(t *testing.T) {
	//
	var b *B
	a:=A{
		Assign:&b,
	}

	bv:=&B{
		Text:"sdfds",
	}
	a.Assign=&bv

	a.Port=1
	fmt.Println(b.Text)

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

var queue = make(chan func())

func TestTrim(t *testing.T)  {
	var s []string
	fmt.Println(reflect.TypeOf(s))

}