package web

import (
	"strconv"
	"net/http"
	"log"
)

var (
	ip 			string
	port       	int
	addr       	string
	spiderMenu 	[]map[string]string
)

func Run()  {
	ip ="127.0.0.1"
	port =9090
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