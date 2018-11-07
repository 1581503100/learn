package utils

import (
	"io/ioutil"
	"os"
)

func SaveFile(file string,data []byte)  {
	f,_:=os.Create(file)
	defer f.Close()
	f.Write(data)
}
func FileReadAll(file string) []byte {
	f,err:=os.Open(file)
	defer f.Close()
	if err !=nil{
		 panic(err)
	}
	res,err:=ioutil.ReadAll(f)
	if(err !=nil){
		panic(err)
	}
	return res
}