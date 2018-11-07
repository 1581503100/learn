package request

import (
	"testing"
	"fmt"
)

func TestGet(t *testing.T) {
	s:=Get("http://www.baidu.com").Execute().ReturnContent().AsString()
	Post("").Body([]byte("")).Execute().ReturnContent().AsString()
	fmt.Println(s)
}
