package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"tools/request"
	"tools/utils"
	"strings"
	"conf"
)

func main()  {
	if(len(os.Args)==1){
		fmt.Println("run get -h")
		return
	}

	conf.InitWithCli([]conf.Flag{
		conf.Flag{
			Key:   "u",
			Usage: "http url you want to visit like http://www.baidu.com",
		},
		conf.Flag{
			Key:   "f",
			Usage: "file you want to save the response",
		},
		conf.Flag{
			Key:   "grep",
			Usage: "find line you want",
		},
	})
	fmt.Println("grep:"+conf.String("grep"))
	run()
}
func run()  {
	url:=os.Args[1]
	file:=conf.String("f")
	if(len(url)==0){
		fmt.Println("run get -help")
		return
	}
	resp:=request.Get(utils.GetUrl(url)).Execute().ReturnContent().AsBytes()
	if(len(file)>0){
		utils.SaveFile(file,resp)
	}else {

		grep:=conf.String("grep")
		fmt.Println(grep)
		if(len(grep)==0){
			fmt.Println(string(resp))
			return
		}
		reader:=bytes.NewReader(resp)
		scanner:=bufio.NewScanner(reader)
		for scanner.Scan(){
			line:=scanner.Text()
			if(strings.Contains(line,grep)){
				fmt.Println(line)
			}

		}
	}
}