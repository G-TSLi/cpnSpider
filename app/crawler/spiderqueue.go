package crawler

import (
	. "cpnSpider/app/spider"
)

type (
	SpiderQueue interface {
		Add(*Spider)
	}
	sq struct {
		list []*Spider
	}
)

func NewSpiderQueue() SpiderQueue {
	return &sq{
		list: []*Spider{},
	}
}

func (self *sq) Add(sp *Spider) {
	self.list = append(self.list, sp)
}