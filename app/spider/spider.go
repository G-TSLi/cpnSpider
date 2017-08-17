package spider

import (
	"cpnSpider/app/scheduler"
	"cpnSpider/app/downloader/request"
	"cpnSpider/runtime/status"
)

type (
	// 蜘蛛规则
	Spider struct {
		Name            string		// 名称（应保证唯一性）
		Description     string		// 描述
		Pausetime       int64       // 暂停区间
		Limit           int64       // 默认限制请求数，0为不限
		Keyin           string		// 自定义参数
		RuleTree        *RuleTree	// 采集规则树

		id        int               // 自动分配的SpiderQueue中的索引
		reqMatrix *scheduler.Matrix // 请求矩阵
		status    int               // 执行状态
	}

	//采集规则树
	RuleTree struct {
		Root  func(*Context)   // 根节点(执行入口)
		Trunk map[string]*Rule // 节点散列表(执行采集过程)
	}

	// 采集规则节点
	Rule struct {
		ItemFields 	[]string												// 结果字段列表
		ParseFunc  	func(*Context)											// 内容解析函数
		AidFunc    	func(*Context, map[string]interface{}) interface{}	// 通用辅助函数
	}
)

// 添加自身到蜘蛛菜单
func (self *Spider) Register() *Spider  {
	self.status = status.STOPPED
	return Species.Add(self)
}





















func (self *Spider) ReqmatrixInit() *Spider {
	self.reqMatrix = scheduler.AddMatrix(self.GetName())
	return self
}

// 获取蜘蛛名称
func (self *Spider) GetName() string {
	return self.Name
}

func (self *Spider) RequestPush(req *request.Request)  {
	self.reqMatrix.Push(req)
}

// 获取蜘蛛描述
func (self *Spider) GetDescription() string {
	return self.Description
}

func (self *Spider) RequestPull() *request.Request {
	return self.reqMatrix.Pull()
}

// 安全返回指定规则
func (self *Spider) GetRule(ruleName string) (*Rule, bool) {
	rule, found := self.RuleTree.Trunk[ruleName]
	return rule, found
}

func (self *Spider) Start() {
	self.RuleTree.Root(GetContext(self, nil))
}


