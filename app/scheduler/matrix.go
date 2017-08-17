package scheduler

import (
	"cpnSpider/app/downloader/request"
	"sync"
	"sort"
)

type Matrix struct {
	maxPage         int64                       // 最大采集数
	spiderName      string                      // 所属Spider
	reqs            map[int][]*request.Request  // [优先级]队列，优先级默认为0
	priorities      []int                       // 优先级顺序，从低到高
	sync.Mutex
}

func newMatrix(spiderName string) *Matrix {
	matrix :=&Matrix{
		spiderName:  spiderName,
		reqs:        make(map[int][]*request.Request),
		priorities:  []int{},
	}
	return matrix
}

// 添加请求到队列，并发安全
func (self *Matrix) Push(req * request.Request)  {
	// 达到请求上限，停止该规则运行
	//if self.maxPage >= 0 {
	//	return
	//}
	//
	var priority = req.GetPriority()

	if _, found := self.reqs[priority]; !found {
		self.priorities = append(self.priorities, priority)
		sort.Ints(self.priorities) // 从小到大排序
		self.reqs[priority] = []*request.Request{}
	}

	// 添加请求到队列
	self.reqs[priority] = append(self.reqs[priority], req)
}

func (self *Matrix) Pull() (req *request.Request)  {
	self.Lock()
	defer self.Unlock()

	// 按优先级从高到低取出请求
	for i :=len(self.reqs) -1; i >=0 ;i--{
		idx := self.priorities[i]
		if len(self.reqs[idx]) > 0 {
			req = self.reqs[idx][0]
			self.reqs[idx] = self.reqs[idx][1:]
		}
	}
	return
}
