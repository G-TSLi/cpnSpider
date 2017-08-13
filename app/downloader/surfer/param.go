package surfer

import (
	"net/url"
	"io"
	"time"
	"net/http"
	"strings"
)

type Param struct {
	method        string
	url           *url.URL
	proxy         *url.URL
	body          io.Reader
	header        http.Header
	enableCookie  bool
	dialTimeout   time.Duration
	connTimeout   time.Duration
	tryTimes      int
	retryPause    time.Duration
	client        *http.Client
}

func NewParam(req Request) (param *Param, err error) {
	param = new(Param)
	param.url, err = UrlEncode(req.GetUrl())
	if err != nil {
		return nil, err
	}

	param.header = req.GetHeader()
	if param.header == nil {
		param.header = make(http.Header)
	}

	switch method:=strings.ToUpper(req.GetMethod());method{
	case "GET":
		param.method=method
	case "POST":
		param.method = method
		param.header.Add("Content-Type", "application/x-www-form-urlencoded")
		param.body = strings.NewReader(req.GetPostData())
	default:
		param.method = "GET"
	}
	param.dialTimeout = req.GetDialTimeout()
	if param.dialTimeout < 0 {
		param.dialTimeout = 0
	}
	param.connTimeout = req.GetConnTimeout()
	param.tryTimes = req.GetTryTimes()
	param.retryPause = req.GetRetryPause()
	return

}