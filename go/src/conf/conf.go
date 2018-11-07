package conf

import (
	"bufio"
	"bytes"
	"flag"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type (
	Flag struct {
		Key     string
		Default string
		Usage   string
		val     *string
	}
)

var confMap = map[string]string{}
var reloadFuns []func()

func Init(files ...string) {
	for _, v := range files {
		parserFile(v)
	}
}
func InitWithBytes(b []byte) {
	reader := bytes.NewReader(b)
	parser(reader)
}

func InitWithReader(reader io.Reader) {
	parser(reader)
}

func InitWithCli(flags []Flag)  {

	for i:=0;i< len(flags);i++{
		flags[i].val=flag.String(flags[i].Key,flags[i].Default,flags[i].Usage)
	}
	flag.Parse()
	for _,v:=range flags{
		Set(v.Key,*v.val)
	}
}
func parserFile(file string) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	parser(f)
}

func parser(reader io.Reader) {
	scaner := bufio.NewScanner(reader)
	for scaner.Scan() {
		line := scaner.Text()
		addProperity(line)
	}
	//callback functions while config is reload or init
	for _, v := range reloadFuns {
		v()
	}
}

func Set(k, v string) {
	confMap[k] = v
}

func addProperity(line string) {
	line = deleteSpace(line)
	if (strings.HasPrefix(line, "#") || len(line) == 0) {
		return
	}
	idx := strings.Index(line, "=")
	if (idx == 0) {
		return
	}
	confMap[deleteSpace(line[0:idx])] = deleteSpace(line[idx+1:])
}

//return int value of giving key
func String(key string) string {
	return confMap[key]
}

//return int value of giving key and return defaultVal by default
func Int(key string, defaultVal int) int {
	if (len(confMap[key]) == 0) {
		return defaultVal
	}
	val, err := strconv.Atoi(confMap[key])
	if err != nil {
		panic(err)
	}
	return val
}

//return bool value of giving key and return defaultVal by default
func Bool(key string, defaultVal bool) bool {
	if (len(confMap[key]) == 0) {
		return defaultVal
	}
	val, err := strconv.ParseBool(confMap[key])
	if err != nil {
		panic(err)
	}
	return val
}

//return float64 value of giving key and return defaultVal by default
func Float64(key string, defaultVal float64) float64 {
	if (len(confMap[key]) == 0) {
		return defaultVal
	}
	val, err := strconv.ParseFloat(confMap[key], 64)
	if err != nil {
		panic(err)
	}
	return val
}

//return string array value of giving key and using ',' to split each item
func Strings(key string) []string {
	if (confMap[key] == "") {
		return []string{}
	}
	return strings.Split(confMap[key], ",")
}

//delete space '\n '\t' '\r' of begin and end of string
func deleteSpace(s string) string {
	var st, ed int
	for i, _ := range s {
		if (s[i] != ' ' && s[i] != '\n' && s[i] != '\r' && s[i] != '\t') {
			st = i;
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if (s[i] != ' ' && s[i] != '\n' && s[i] != '\r' && s[i] != '\t') {
			ed = i + 1
			break
		}
	}
	return s[st:ed]
}

func AddReloadHanler(f func()) {
	reloadFuns = append(reloadFuns, f)
}
func toLine(s string) string {
	if (len(s) == 0) {
		return ""
	}
	var bf []byte
	sp := 'A' - 'a'
	if (s[1] >= 'A' && s[0] <= 'Z') {
		bf = append(bf, s[0]-byte(sp))
	} else {
		bf = append(bf, s[0])
	}
	for i := 1; i < len(s); i++ {
		if (s[i] <= 'Z' && s[i] >= 'A') {
			bf = append(bf, '_', s[i]-byte(sp))
		} else {
			bf = append(bf, s[i])
		}
	}
	return string(bf)
}
func Unmarshal(i interface{}, prefix string) {
	t := reflect.TypeOf(i).Elem()
	v := reflect.ValueOf(i).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if (tag == "") {
			tag = toLine(field.Name)
		}
		switch field.Type.Name() {
			case "string":
				v.Field(i).Set(reflect.ValueOf(String(prefix + tag)));
				break
			case "int":
				v.Field(i).Set(reflect.ValueOf(Int(prefix+tag, 0)));
				break
			case "float64":
				v.Field(i).Set(reflect.ValueOf(Float64(prefix+tag, 0)));
				break
			case "bool":
				v.Field(i).Set(reflect.ValueOf(Bool(prefix+tag, false)));
				break
		}
	}

}

