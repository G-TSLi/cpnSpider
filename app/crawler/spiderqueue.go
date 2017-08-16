package crawler

import (
	. "cpnSpider/app/spider"
)

type (
	SpiderQueue interface {
		Add(*Spider)
		GetAll() []*Spider
		GetByIndex(int) *Spider
		Len() int
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

func (self *sq) GetAll() []*Spider {
	return self.list
}

func (self *sq) Len() int {
	return len(self.list)
}

func (self *sq) GetByIndex(idx int) *Spider {
	return self.list[idx]
}