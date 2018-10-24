package conf

import (
	"testing"
	"fmt"
)

func Test(t *testing.T) {
	//init the files you want to load
	Init("app.conf")
	fmt.Println(Int("port",9090))
	fmt.Println(String("mysql_name"))
}
