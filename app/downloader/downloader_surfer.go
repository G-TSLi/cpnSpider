package downloader

import (
	"cpnSpider/app/downloader/surfer"
	"cpnSpider/app/downloader/request"
	"cpnSpider/app/spider"
	"errors"
	"net/http"
)

type Surfer struct {
	surf    surfer.Surfer
}

var SurferDownloader = &Surfer{
	surf:    surfer.New(),
}

func (self *Surfer) Download(sp *spider.Spider, cReq *request.Request) *spider.Context{


	ctx := spider.GetContext(sp, cReq)

	var resp *http.Response
	var err error

	resp, err = self.surf.Download(cReq)

	if resp.StatusCode >= 400 {
		err = errors.New("响应状态 " + resp.Status)
	}
	ctx.SetResponse(resp).SetError(err)
	return ctx
}