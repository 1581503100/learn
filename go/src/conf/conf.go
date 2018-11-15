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

//init the config with args from commond line
func InitWithCli(flags []Flag) {

	for i := 0; i < len(flags); i++ {
		flags[i].val = flag.String(flags[i].Key, flags[i].Default, flags[i].Usage)
	}
	flag.Parse()
	for _, v := range flags {
		Set(v.Key, *v.val)
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
	if (strings.HasPrefix(line, "#") || len(line) == 0 || !strings.Contains(line, "=")) {
		return
	}
	lis := strings.SplitN(line, "=", 2)
	if (len(lis) != 2) {
		return
	}
	confMap[deleteSpace(lis[0])] = deleteSpace(lis[1])
}

//return int value of giving key
func String(key string) string {
	return confMap[key]
}

func StringD(k, def string) string {
	v := confMap[k]
	if v == "" {
		return def
	}
	return v
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
	return trimStrs(strings.Split(confMap[key], ","))
}

// return string array value of specific key spliting by specific sep
func StringsS(k, sep string) []string {
	if confMap[k] == "" {
		return []string{}
	}
	return trimStrs(strings.Split(confMap[k], sep))
}
func Ints(k string)[]int  {
	if confMap[k] == "" {
		return []int{}
	}
	vs:=Strings(k)
	var vals []int
	for _,v:=range vs{
		iv,err:=strconv.Atoi(v)
		if(err !=nil){
			panic(err)
		}
		vals = append(vals,iv)
	}
	return vals
}
func trimStrs(ss []string) []string {
	for i := 0; i < len(ss); i++ {
		ss[i] = deleteSpace(ss[i])
	}
	return ss
}

//delete space '\n '\t' '\r' of begin and end of string
func deleteSpace(s string) string {
	return strings.Trim(s, " \r\n\t")
}

func AddReloadHanler(f func()) {
	reloadFuns = append(reloadFuns, f)
}

func toLine(s string) string {
	if (len(s) == 0) {
		return ""
	}
	var bf bytes.Buffer
	sp := ('A' - 'a')
	if (s[1] >= 'A' && s[0] <= 'Z') {
		bf.WriteByte(s[0] - byte(sp))
	} else {
		bf.WriteByte(s[0])
	}
	for i := 1; i < len(s); i++ {
		if (s[i] <= 'Z' && s[i] >= 'A') {
			bf.WriteByte('_')
			bf.WriteByte(s[i] - byte(sp))
		} else {
			bf.WriteByte(s[i])
		}
	}
	return bf.String()
}

func toUp(s string) string {
	if (s == "") {
		return s
	}
	sp := 'A' - 'a'
	var bf bytes.Buffer
	if (s[0] <= 'Z' && s[0] >= 'A') {
		bf.WriteByte(s[0] - byte(sp))
	} else {
		bf.WriteByte(s[0])
	}
	bf.WriteString(s[1:])
	return bf.String()
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
		key := prefix + tag
		switch field.Type.Name() {
		case "string":
			v.Field(i).Set(reflect.ValueOf(String(key)));
			break
		case "int":
			v.Field(i).Set(reflect.ValueOf(Int(key, 0)));
			break
		case "float64":
			v.Field(i).Set(reflect.ValueOf(Float64(key, 0)));
			break
		case "bool":
			v.Field(i).Set(reflect.ValueOf(Bool(key, false)));
			break
		case "[]string":
			v.Field(i).Set(reflect.ValueOf(String(key)));
			break
		case "[]int":
			v.Field(i).Set(reflect.ValueOf(Ints(key)));
			break
		}
	}
}
