package spider

import (
	"cpnSpider/app/scheduler"
	"cpnSpider/app/downloader/request"
)

type (
	Spider struct {
		Name            string
		reqMatrix *scheduler.Matrix // 请求矩阵
	}
)

// 获取蜘蛛名称
func (self *Spider) GetName() string {
	return self.Name
}

func (self *Spider) RequestPush(req *request.Request)  {
	self.reqMatrix.Pull(req)
}

func (self *Spider) RequestPull() *request.Request  {
	return nil
}