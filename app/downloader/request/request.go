package request

import (
	"net/http"
	"time"
	"strings"
)

type Request struct {
	Url           string          //目标URL，必须设置
	Rule          string
	Method        string          //GET POST POST-M HEAD
	Header        http.Header     //请求头信息
	EnableCookie  bool            //是否使用cookies，在Spider的EnableCookie设置
	PostData      string          //POST values
	DialTimeout   time.Duration   //创建连接超时 dial tcp: i/o timeout
	ConnTimeout   time.Duration   //连接状态超时 WSARecv tcp: i/o timeout
	TryTimes      int             //尝试下载的最大次数
	RetryPause    time.Duration   //下载失败后，下次尝试下载的等待时间
	Priority      int             //指定调度优先级，默认为0（最小优先级为0）
	//Surfer下载器内核ID
	//0为Surf高并发下载器，各种控制功能齐全
	//1为PhantomJS下载器，特点破防力强，速度慢，低并发
	DownloaderID int

	proxy  string //当用户界面设置可使用代理IP时，自动设置代理
}


// 获取Url
func (self *Request) GetUrl() string {
	return self.Url
}

// 获取Http请求的方法名称 (注意这里不是指Http GET方法)
func (self *Request) GetMethod() string {
	return self.Method
}

// 设定Http请求方法的类型
func (self *Request) SetMethod(method string) *Request {
	self.Method = strings.ToUpper(method)
	return self
}

func (self *Request) SetRuleName(ruleName string) *Request {
	self.Rule = ruleName
	return self
}


func (self *Request) SetUrl(url string) *Request {
	self.Url = url
	return self
}

func (self *Request) GetPostData() string {
	return self.PostData
}

func (self *Request) GetHeader() http.Header {
	return self.Header
}

func (self *Request) SetHeader(key, value string) *Request {
	self.Header.Set(key, value)
	return self
}

func (self *Request) AddHeader(key, value string) *Request {
	self.Header.Add(key, value)
	return self
}

func (self *Request) GetEnableCookie() bool {
	return self.EnableCookie
}

func (self *Request) SetEnableCookie(enableCookie bool) *Request {
	self.EnableCookie = enableCookie
	return self
}

func (self *Request) GetCookies() string {
	return self.Header.Get("Cookie")
}

func (self *Request) SetCookies(cookie string) *Request {
	self.Header.Set("Cookie", cookie)
	return self
}

func (self *Request) GetDialTimeout() time.Duration {
	return self.DialTimeout
}

func (self *Request) GetConnTimeout() time.Duration {
	return self.ConnTimeout
}

func (self *Request) GetTryTimes() int {
	return self.TryTimes
}

func (self *Request) GetRetryPause() time.Duration {
	return self.RetryPause
}

func (self *Request) GetProxy() string {
	return self.proxy
}

func (self *Request) SetProxy(proxy string) *Request {
	self.proxy = proxy
	return self
}

func (self *Request) GetPriority() int {
	return self.Priority
}

func (self *Request) SetPriority(priority int) *Request {
	self.Priority = priority
	return self
}

func (self *Request) GetDownloaderID() int {
	return self.DownloaderID
}

func (self *Request) SetDownloaderID(id int) *Request {
	self.DownloaderID = id
	return self
}