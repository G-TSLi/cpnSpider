package web

import (
	"strconv"
	"net/http"
	"log"
	"cpnSpider/app"
)

var (
	ip 			string
	port       	int
	addr       	string
	spiderMenu 	[]map[string]string
)

func Run()  {
	appInit()

	ip ="127.0.0.1"
	port =9091
	// web服务器地址
	addr = ip + ":" + strconv.Itoa(port)

	// 预绑定路由
	Router()

	log.Printf("Server running on %v\n", addr)

	//监听端口
	err := http.ListenAndServe(addr, nil) //设置监听的端口
	if err != nil {
	}
}

func appInit() {
	spiderMenu = func() (spmenu []map[string]string) {
		// 获取蜘蛛家族
		for _, sp := range app.LogicApp.GetSpiderLib() {
			spmenu = append(spmenu, map[string]string{"name": sp.GetName(), "description": sp.GetDescription()})
		}
		return spmenu
	}()
}