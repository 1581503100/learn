package conf

import (
	"bufio"
	"bytes"
	"github.com/astaxie/beego/logs"
	"io"
	"os"
	"strconv"
	"strings"
)
var confMap = map[string]string{}
var reloadFuns []func()


func Init(files ...string){
	for _,v:=range files{
		parserFile(v)
	}
}
func InitWithBytes(b []byte)  {
	reader:=bytes.NewReader(b)
	parser(reader)
}
func parserFile(file string)  {
	f,err:=os.Open(file)
	defer f.Close()
	if err !=nil{
		panic(err)
	}
	parser(f)
}

func parser(reader io.Reader)  {
	scaner:=bufio.NewScanner(reader)
	for scaner.Scan(){
		line:=scaner.Text()
		if(strings.HasPrefix(line,"#") || len(line)==0){
			continue
		}
		idx:=strings.Index(line,"=")
		if(idx ==0 ){
			continue
		}
		confMap[deleteSpace(line[0:idx])]=deleteSpace(line[idx+1:])
	}

	logs.Info("------------------------load config:",confMap)
	for _,v:=range reloadFuns{
		v()
	}
}
//return int value of giving key
func String(key string)string  {
	return confMap[key]
}
//return int value of giving key and return defaultVal by default
func Int(key string,defaultVal int) int {
	if(len(confMap[key])==0){
		return defaultVal
	}
	val,err:=strconv.Atoi(confMap[key])
	if err!=nil{
		panic(err)
	}
	return val
}
//return bool value of giving key and return defaultVal by default
func Bool(key string,defaultVal bool) bool  {
	if(len(confMap[key])==0){
		return defaultVal
	}
	val,err:=strconv.ParseBool(confMap[key])
	if err!=nil{
		panic(err)
	}
	return val
}

//return float64 value of giving key and return defaultVal by default
func Float64(key string,defaultVal float64) float64  {
	if(len(confMap[key])==0){
		return defaultVal
	}
	val,err:=strconv.ParseFloat(confMap[key],64)
	if err!=nil{
		panic(err)
	}
	return val
}
//return string array value of giving key and using ',' to split each item
func Strings(key string) []string {
	if(confMap[key]==""){
		return []string{}
	}
	return strings.Split(confMap[key],",")
}
//delete space '\n '\t' '\r' of begin and end of string
func deleteSpace(s string) string {
	var st,ed int
	for i,_:=range s{
		if( s[i] !=' ' && s[i]!='\n' && s[i]!='\r' && s[i]!='\t'){
			st=i;
			break
		}
	}
	for i:=len(s)-1;i>=0;i--{
		if(s[i] !=' ' && s[i]!='\n' && s[i]!='\r' && s[i]!='\t'){
			ed=i+1
			break
		}
	}
	return s[st:ed]
}

func AddReloadHanler(f func()){
	reloadFuns =append(reloadFuns,f)
}