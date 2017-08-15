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
		log.Println(req)
		wsApi[util.Atoa(req["operate"])](sessID, req)
	}
}

func init()  {
	wsApi["init"] = func(sessID string, req map[string]interface{}) {
		var mode = util.Atoi(req["mode"])
		var port = util.Atoi(req["port"])
		var master = util.Atoa(req["ip"]) //服务器(主节点)地址，不含端口
		app.LogicApp = app.LogicApp.ReInit(mode, port, master) // 切换运行模式
	}
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
