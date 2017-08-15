package web

import (
	"net/http"
	ws "cpnSpider/common/websocket"
)

// 路由
func Router() {

	http.Handle("/ws", ws.Handler(wsHandle))

	http.HandleFunc("/", web)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(assetFS())))
}