package spider

import (
	"cpnSpider/app/downloader/request"
	"sync"
	"net/http"
)

type Context struct {
	spider   *Spider           // 规则
	Request  *request.Request  // 原始请求
	Response *http.Response    // 响应流，其中URL拷贝自*request.Request
	text     []byte            // 下载内容Body的字节流格式
	err      error             // 错误标记
}


var (
	contextPool = &sync.Pool{
		
	}
)

func GetContext(sp *Spider, req *request.Request) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.spider = sp
	ctx.Request = req
	return ctx
}

func (self *Context) SetResponse(resp *http.Response) *Context {
	self.Response = resp
	return self
}

func (self *Context) SetError (err error)  {
	self.err = err
}

