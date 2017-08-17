package spider

import (
	"bytes"
	"cpnSpider/app/downloader/request"
	"cpnSpider/common/goquery"
	"sync"
	"net/http"
	"io/ioutil"
	"mime"
	"strings"
	"io"
	"golang.org/x/net/html/charset"
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
		New: func() interface{} {
			return &Context{
			}
		},
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

// 获取当前规则名。
func (self *Context) GetRuleName() string {
	return self.Request.GetRuleName()
}

func (self *Context) SetError (err error)  {
	self.err = err
}

// GetBodyStr returns plain string crawled.
func (self *Context) initText() {
	var err error
	var contentType, pageEncode string
	// 优先从响应头读取编码类型
	contentType = self.Response.Header.Get("Content-Type")
	if _, params, err := mime.ParseMediaType(contentType); err == nil {
		if cs, ok := params["charset"]; ok {
			pageEncode = strings.ToLower(strings.TrimSpace(cs))
		}
	}
	// 响应头未指定编码类型时，从请求头读取
	if len(pageEncode) == 0 {
		contentType = self.Request.Header.Get("Content-Type")
		if _, params, err := mime.ParseMediaType(contentType); err == nil {
			if cs, ok := params["charset"]; ok {
				pageEncode = strings.ToLower(strings.TrimSpace(cs))
			}
		}
	}

	switch pageEncode {
	// 不做转码处理
	case "utf8", "utf-8":
	default:
		var destReader io.Reader

		if len(pageEncode) == 0 {
			destReader, err = charset.NewReader(self.Response.Body, "")
		} else {
			destReader, err = charset.NewReaderLabel(pageEncode, self.Response.Body)
		}
		if err == nil {
			self.text, err = ioutil.ReadAll(destReader)
			if err == nil {
				self.Response.Body.Close()
				return
			} else {
			}
		} else {
		}
	}

	// 不做转码处理
	self.text, err = ioutil.ReadAll(self.Response.Body)
	self.Response.Body.Close()
	if err != nil {
		panic(err.Error())
		return
	}
}

func (self *Context) initDom() *goquery.Document {
	if self.text == nil {
		self.initText()
	}
	var err error
	self.dom, err = goquery.NewDocumentFromReader(bytes.NewReader(self.text))
	if err != nil {
		panic(err.Error())
	}
	return self.dom
}

func (self *Context) GetDom() *goquery.Document {
	if self.dom == nil {
		self.initDom()
	}
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
	if len(ruleName) == 0 {
		if self.Response == nil {
			return
		}
		name = self.GetRuleName()
	} else {
		name = ruleName[0]
	}
	rule, found = self.spider.GetRule(name)
	return
}

func (self *Context) AddQueue(req *request.Request) *Context {
	self.spider.RequestPush(req)
	return self
}
