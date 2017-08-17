package crawler

import (
	"cpnSpider/app/spider"
	"cpnSpider/app/downloader"
	"cpnSpider/app/downloader/request"
	"time"
)


type (
	Crawler interface {
		Init(*spider.Spider) Crawler //初始化采集引擎
		Run()
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

func (self *crawler) Init(sp *spider.Spider) Crawler {
	self.Spider = sp.ReqmatrixInit()
	return self
}

// 任务执行入口
func (self *crawler) Run()  {

	// 运行处理协程
	c := make(chan bool)
	go func() {
		self.run()
		close(c)
	}()

	// 启动任务
	self.Spider.Start()

	<-c // 等待处理协程退出


}

func (self *crawler) run()  {
	for {
		// 队列中取出一条请求并处理
		req := self.GetOne()
		if req == nil {
			time.Sleep(20 * time.Millisecond)
			continue
		}

		go func(req *request.Request) {
			self.Process(req)
		}(req)
	}
}

func (self *crawler) Process(req *request.Request)  {
	var (
		sp      = self.Spider
	)
	var ctx = self.Downloader.Download(sp, req) // download page
	// 过程处理，提炼数据
	ctx.Parse(req.GetRuleName())

}

func (self *crawler) GetOne() *request.Request  {
	return self.Spider.RequestPull()
}


