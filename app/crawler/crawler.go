package crawler

import (
	"cpnSpider/app/spider"
	"cpnSpider/app/downloader"
	"cpnSpider/app/downloader/request"
)


type (
	Crawler interface {

	}
	crawler struct {
		*spider.Spider                 //执行的采集规则
		downloader.Downloader          //全局公用的下载器
	}
)

func New() Crawler {
	return &crawler{
		Downloader: downloader.SurferDownloader,
	}
}

// 任务执行入口
func (self *crawler) Run()  {
	go func() {
		self.run()
	}()
}

func (self *crawler) run()  {
	for {
		// 队列中取出一条请求并处理
		req := self.GetOne()
		if req == nil {
			self.Process(req)
		}
	}
}

func (self *crawler) Process(req *request.Request)  {
	var (
		sp      = self.Spider
	)
	self.Downloader.Download(sp, req) // download p
}

func (self *crawler) GetOne() *request.Request  {
	return self.Spider.RequestPull()
}


