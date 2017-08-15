package web

import (
	"cpnSpider/app"
	"cpnSpider/app/spider"
	"cpnSpider/common/util"
	ws "cpnSpider/common/websocket"
	"sync"
	"log"
)

type SocketController struct {
	connPool     map[string]*ws.Conn
	wchanPool    map[string]*Wchan
	connRWMutex  sync.RWMutex
	wchanRWMutex sync.RWMutex
}

func (self *SocketController) GetConn(sessID string) *ws.Conn {
	self.connRWMutex.RLock()
	defer self.connRWMutex.RUnlock()
	return self.connPool[sessID]
}

func (self *SocketController) GetWchan(sessID string) *Wchan {
	self.wchanRWMutex.RLock()
	defer self.wchanRWMutex.RUnlock()
	return self.wchanPool[sessID]
}

func (self *SocketController) Add(sessID string, conn *ws.Conn) {
	self.connRWMutex.Lock()
	self.wchanRWMutex.Lock()
	defer self.connRWMutex.Unlock()
	defer self.wchanRWMutex.Unlock()

	self.connPool[sessID] = conn
	self.wchanPool[sessID] = newWchan()
}


func (self *SocketController) Write(sessID string, void map[string]interface{}, to ...int) {
	self.wchanRWMutex.RLock()
	defer self.wchanRWMutex.RUnlock()

	// to为1时，只向当前连接发送；to为-1时，向除当前连接外的其他所有连接发送；to为0时或为空时，向所有连接发送
	var t int = 0
	if len(to) > 0 {
		t = to[0]
	}

	void["mode"] = app.LogicApp.GetAppConf("mode").(int)

	switch t {
	case 1:
		wc := self.wchanPool[sessID]
		if wc == nil {
			return
		}
		void["initiative"] = true
		wc.wchan <- void

	case 0, -1:
		l := len(self.wchanPool)
		for _sessID, wc := range self.wchanPool {
			if t == -1 && _sessID == sessID {
				continue
			}
			_void := make(map[string]interface{}, l)
			for k, v := range void {
				_void[k] = v
			}
			if _sessID == sessID {
				_void["initiative"] = true
			} else {
				_void["initiative"] = false
			}
			log.Println(_void)
			wc.wchan <- _void
		}
	}
}


type Wchan struct {
	wchan chan interface{}
}

func newWchan() *Wchan {
	return &Wchan{
		wchan: make(chan interface{}, 1024),
	}
}

var (
	wsApi = map[string]func(string, map[string]interface{}){}
	Sc    = &SocketController{
		connPool:  make(map[string]*ws.Conn),
		wchanPool: make(map[string]*Wchan),
	}
)

func wsHandle(conn *ws.Conn) {
	sess, _ := globalSessions.SessionStart(nil, conn.Request())
	sessID := sess.SessionID()
	if Sc.GetConn(sessID) == nil {
		Sc.Add(sessID, conn)
	}

	go func() {
		var err error
		for info := range Sc.GetWchan(sessID).wchan {
			if _, err = ws.JSON.Send(conn, info); err != nil {
				return
			}
		}
	}()

	for {
		var req map[string]interface{}
		if err := ws.JSON.Receive(conn, &req); err != nil {
			return
		}
		wsApi[util.Atoa(req["operate"])](sessID, req)
	}
}

func init()  {
	wsApi["init"] = func(sessID string, req map[string]interface{}) {
		var mode = util.Atoi(req["mode"])
		var port = util.Atoi(req["port"])
		var master = util.Atoa(req["ip"]) //服务器(主节点)地址，不含端口
		app.LogicApp = app.LogicApp.ReInit(mode, port, master) // 切换运行模式
		// 写入发送通道
		Sc.Write(sessID, tplData(mode))
	}
}

func tplData(mode int) map[string]interface{} {
	var info = map[string]interface{}{"operate": "init", "mode": mode}

	// 蜘蛛家族清单
	info["spiders"] = map[string]interface{}{
		"menu": spiderMenu,
		"curr": func() interface{} {
			l := app.LogicApp.GetSpiderQueue().Len()
			if l == 0 {
				return 0
			}
			var curr = make(map[string]bool, l)
			for _, sp := range app.LogicApp.GetSpiderQueue().GetAll() {
				curr[sp.GetName()] = true
			}

			return curr
		}(),
	}

	// 并发协程上限
	info["ThreadNum"] = map[string]int{
		"max":  999999,
		"min":  1,
		"curr": app.LogicApp.GetAppConf("ThreadNum").(int),
	}

	// 暂停区间/ms(随机: Pausetime/2 ~ Pausetime*2)
	info["Pausetime"] = map[string][]int64{
		"menu": {0, 100, 300, 500, 1000, 3000, 5000, 10000, 15000, 20000, 30000, 60000},
		"curr": []int64{app.LogicApp.GetAppConf("Pausetime").(int64)},
	}

	// 代理IP更换的间隔分钟数
	info["ProxyMinute"] = map[string][]int64{
		"menu": {0, 1, 3, 5, 10, 15, 20, 30, 45, 60, 120, 180},
		"curr": []int64{app.LogicApp.GetAppConf("ProxyMinute").(int64)},
	}

	return info
}


func setConf(req map[string]interface{})  {
	if tn :=util.Atoi(req["ThreadNum"]); tn == 0 {
		app.LogicApp.SetAppConf("ThreadNum", 1)
	}else {
		app.LogicApp.SetAppConf("ThreadNum", tn)
	}
	app.LogicApp.
		SetAppConf("Pausetime", int64(util.Atoi(req["Pausetime"]))).
		SetAppConf("ProxyMinute", int64(util.Atoi(req["ProxyMinute"]))).
		SetAppConf("Limit", int64(util.Atoi(req["Limit"])))

	setSpiderQueue(req)
}

func setSpiderQueue(req map[string]interface{}) {
	spiders := []*spider.Spider{}
	app.LogicApp.SpiderPrepare(spiders)
}
