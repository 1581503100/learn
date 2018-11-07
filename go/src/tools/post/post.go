package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"github.com/joho/godotenv"
	"tools/request"
	"tools/utils"
	"strings"
)

func main()  {
	if(len(os.Args)==1){
		fmt.Println("run post -h")
		return
	}
	if env := os.Getenv("PLUGIN_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app:=cli.NewApp()
	app.Name="post"
	app.Usage="post http"
	app.Action=run
	app.Version="1.0.0"
	app.Flags=[]cli.Flag{
		cli.StringFlag{
			Name:"u",
			Usage:"http url you want to visit like http://www.baidu.com",
		},
		cli.StringFlag{
			Name:"f",
			Usage:"file you want to save the response",
		},
		cli.StringFlag{
			Name:"grep",
			Usage:"find line you want",
		},
		cli.StringFlag{
			Name:"b",
			Usage:"post body",
		},
		cli.StringFlag{
			Name:"bf",
			Usage:"post body file",
		},
	}
	if err:=app.Run(os.Args);err!=nil{
		fmt.Println(err)
	}
}
func run(c *cli.Context)  {
	url:=os.Args[1]
	file:=c.String("f")
	if(len(url)==0){
		fmt.Println("run get -help")
		return
	}

	req:=request.Post(utils.GetUrl(url))
	if body:=c.String("b"); len(body)>0{
		req.Body([]byte(body))
	}

	if bf:=c.String("bf"); len(bf)>0{
		req.Body(utils.FileReadAll(bf))
	}
	resp:=req.Execute().ReturnContent().AsBytes()
	if(len(file)>0){
		utils.SaveFile(file,resp)
	}else {

		grep:=c.String("grep")
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