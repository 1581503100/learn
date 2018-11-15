package tests

import (
	"testing"
	"fmt"
)
type (
	A1 struct {
		AssginTo ** A1
		Hello string
	}
)
func Test_s(t *testing.T)  {
	var as *A1
	a:=A1{
		AssginTo:&as,
	}
	as.Hello = "sss"
	fmt.Println(a.Hello)
}
