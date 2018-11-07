package request

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


type (
	Request struct {
		method     string
		url        string
		returnData []byte
		timeout    time.Duration
		body       []byte
		header     map[string]string
		cookies    []*http.Cookie
	}

	ReturnContent struct {
		data []byte
	}
)

func inits(met,url string) *Request  {
	req :=&Request{
		method:  met,
		url:     url,
		timeout: 3*time.Second,
		header: make(map[string]string),
	}
	return req
}
func Get(url string) *Request {
	return inits("GET",url)
}

func Post(url string)*Request  {
	return inits("POST",url)
}

func Put(url string)*Request  {
	return inits("PUT",url)
}

func Delete(url string)*Request  {
	return inits("DELETE",url)
}


func (this *Request)Execute() *Request {
	return this.request(this.method,this.url,this.body)
}
func(this *Request) Body(b []byte)  *Request{
	this.body=b
	return this
}

func (this *Request) AddHeader(key,val string)*Request  {
	this.header[key]=val
	return this
}

func (this *Request)AddCookie(c *http.Cookie) *Request {
	this.cookies=append(this.cookies,c)
	return this;
}

func (this *Request)Timeout(t time.Duration) *Request {
	this.timeout=t
	return this
}

func (this *Request)request(method ,url string,body []byte) *Request {
	defer func() {
		if err:=recover();err!=nil{
			log.Println(err)
		}
	}()
	client:=http.Client{Timeout:this.timeout}
	req,err:=http.NewRequest(method,url,bytes.NewReader(body))
	//add header
	for k,v:=range this.header{
		req.Header.Set(k,v)
	}

	for _,v:=range this.cookies{
		req.AddCookie(v)
	}
	if(err !=nil){
		panic(err)
	}
	resp,err:=client.Do(req)
	if(err !=nil){
		panic(err)
	}
	defer resp.Body.Close()
	result,err:=ioutil.ReadAll(resp.Body)
	if(err!=nil){
		panic(err)
	}
	this.returnData = result

	return this
}

func (this *Request)ReturnContent() *ReturnContent  {
	return &ReturnContent{data:this.returnData}
}



func (this *ReturnContent) AsString() string {
	if(this.data !=nil){
		return string(this.data)
	}
	return ""
}
func (this *ReturnContent)AsBytes() []byte  {
	return this.data
}
func (this *ReturnContent)AsReader() io.Reader  {
	return bytes.NewReader(this.data)
}
