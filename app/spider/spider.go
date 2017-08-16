package spider

import (
	"cpnSpider/app/scheduler"
	"cpnSpider/app/downloader/request"
	"cpnSpider/runtime/status"
)

type (
	Spider struct {
		Name            string		// 名称
		Description     string		// 描述
		Keyin           string
		reqMatrix *scheduler.Matrix // 请求矩阵
		RuleTree        *RuleTree

		status    int               // 执行状态
	}

	//采集规则树
	RuleTree struct {
		Root  func(*Context)   // 根节点(执行入口)
		Trunk map[string]*Rule // 节点散列表(执行采集过程)
	}

	// 采集规则节点
	Rule struct {
		ItemFields 	[]string
		ParseFunc  	func(*Context)
		AidFunc    	func(*Context, map[string]interface{}) interface{}
	}
)

func (self Spider) Register() *Spider  {
	self.status = status.STOPPED
	return Species.Add(&self)
}

// 获取蜘蛛名称
func (self *Spider) GetName() string {
	return self.Name
}

func (self *Spider) RequestPush(req *request.Request)  {
	self.reqMatrix.Pull(req)
}

// 获取蜘蛛描述
func (self *Spider) GetDescription() string {
	return self.Description
}

func (self *Spider) RequestPull() *request.Request  {
	return nil
}

// 安全返回指定规则
func (self *Spider) GetRule(ruleName string) (*Rule, bool) {
	rule, found := self.RuleTree.Trunk[ruleName]
	return rule, found
}


