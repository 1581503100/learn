package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {
	var fileInfos []os.FileInfo
	if(len(os.Args)==1){
		fileInfos,_ =ioutil.ReadDir(".")
	}else{
		fileInfos,_=ioutil.ReadDir(os.Args[1])
	}
	output(fileInfos)
}
func output(fileinfos []os.FileInfo)  {
	fmt.Println("mod              isdir       name")
	for _,v:=range fileinfos{
		fmt.Printf("%s       %t       %s \n",v.Mode(),v.IsDir(),v.Name())
	}
}