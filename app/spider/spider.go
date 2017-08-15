package spider

import (
	"cpnSpider/app/scheduler"
	"cpnSpider/app/downloader/request"
)

type (
	Spider struct {
		reqMatrix *scheduler.Matrix // 请求矩阵
	}
)

func (self *Spider) RequestPush(req *request.Request)  {
	self.reqMatrix.Pull(req)
}

func (self *Spider) RequestPull() *request.Request  {
	return nil
}