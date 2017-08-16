package spider

import (
	"cpnSpider/app/downloader/request"
	"cpnSpider/common/goquery"
	"sync"
	"net/http"
)

type Context struct {
	spider   *Spider           // 规则
	Request  *request.Request  // 原始请求
	Response *http.Response    // 响应流，其中URL拷贝自*request.Request
	text     []byte            // 下载内容Body的字节流格式
	dom      *goquery.Document // 下载内容Body为html时，可转换为Dom的对象
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

func (self *Context) GetDom() *goquery.Document {
	return self.dom
}

func (self *Context) Aid(aid map[string]interface{},ruleName ...string) interface{} {
	_, rule, found := self.getRule(ruleName...)
	if !found {
		if len(ruleName) > 0 {

		} else {

		}
		return nil
	}
	return rule.AidFunc(self, aid)
}

func (self *Context) Parse(ruleName ...string) *Context {

	_ruleName, rule, found := self.getRule(ruleName...)
	if self.Response != nil {
		self.Request.SetRuleName(_ruleName)
	}
	if !found {
		self.spider.RuleTree.Root(self)
		return self
	}
	if rule.ParseFunc == nil {
		return self
	}
	rule.ParseFunc(self)
	return self
}

func (self *Context) getRule(ruleName ...string) (name string, rule *Rule, found bool) {
	name = ruleName[0]
	rule, found = self.spider.GetRule(name)
	return
}

func (self *Context) AddQueue(req *request.Request) *Context {
	self.spider.RequestPush(req)
	return self
}
