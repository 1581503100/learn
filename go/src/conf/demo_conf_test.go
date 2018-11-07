package conf

import (
	"fmt"
	"testing"
)

type A struct {
	Port          int
	MysqlName     string
	MysqlPassword string
}

func Test(t *testing.T) {
	//init the files you want to load
	Init("app.conf")
	fmt.Println(Int("port", 9090))
	fmt.Println(String("mysql_name"))
	fmt.Println(String("JAVA_HOME"))
	var a A
	Unmarshal(&a, "mysql.")
	fmt.Println(a)
}
func TestName(t *testing.T) {
	fmt.Println(toLine("JavaHome"))
}
