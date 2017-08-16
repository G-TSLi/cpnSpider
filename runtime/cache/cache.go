package cache

type AppConf struct {
	Mode           	int    	// 	节点角色
	Port           	int    	// 	主节点端口
	Master         	string 	// 	服务器(主节点)地址，不含端口
	ThreadNum      	int    	// 	全局最大并发量
	Pausetime      	int64  	// 	暂停时长参考/ms(随机: Pausetime/2 ~ Pausetime*2)
	Limit  			int64 	//	采集上限，0为不限
	ProxyMinute    	int64  	// 	代理IP更换的间隔分钟数
	Keyins 			string 	// 	自定义输入，后期切分为多个任务的Keyin自定义配置
}

// 该初始值即默认值
var Task = new(AppConf)