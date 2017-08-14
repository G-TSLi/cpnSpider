package downloader

import (
	"cpnSpider/app/spider"
	"cpnSpider/app/downloader/request"
)

type Downloader interface {
	Download(*spider.Spider, *request.Request) *spider.Context
}