package conf

import (
	"os"
	"bufio"
	"io"
	"strings"
	"strconv"
)
var confMap map[string]string
func Init(files ...string){
	confMap=make(map[string]string)
	for _,v:=range files{
		parserFile(v)
	}
}

func parserFile(file string)  {
	f,err:=os.Open(file)
	defer f.Close()
	if err !=nil{
		panic(err)
	}
	rd:=bufio.NewReader(f)
	for{
		line,err:=rd.ReadString('\n')
		if err!=nil || io.EOF ==err{
			break
		}
		if(strings.HasPrefix(line,"#")){
			continue
		}
		kvs:=strings.Split(line,"=")
		if(len(kvs)!=2){
			continue
		}
		confMap[deleteSpace(kvs[0])]=deleteSpace(kvs[1])
	}

}

func String(key string)string  {
	return confMap[key]
}

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

func Strings(key string) []string {
	return strings.Split(confMap[key],",")
}

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