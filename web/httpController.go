package web

import (
	"cpnSpider/common/session"
	"cpnSpider/app"
	"net/http"
	"text/template"
	"cpnSpider/config"
	"cpnSpider/runtime/status"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"cpnSpiderSession", "enableSetCookie,omitempty": true, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 157680000, "providerConfig": ""}`)
	// go globalSessions.GC()
}

// 处理web页面请求
func web(rw http.ResponseWriter, req *http.Request) {
	sess, _ := globalSessions.SessionStart(rw, req)
	defer sess.SessionRelease(rw)
	t, err := template.ParseFiles("web/views/index.html") //解析模板文件
	if err != nil {
	}
	data := map[string]interface{}{
		"title":   config.NAME,
		"version": config.VERSION,
		"author":  config.AUTHOR,
		"mode": map[string]int{
			"offline": status.OFFLINE,
			"server":  status.SERVER,
			"client":  status.CLIENT,
			"unset":   status.UNSET,
			"curr":    app.LogicApp.GetAppConf("mode").(int),
		},
		"status": map[string]int{
			"stopped": status.STOPPED,
			"stop":    status.STOP,
			"run":     status.RUN,
			"pause":   status.PAUSE,
		},
		"port": app.LogicApp.GetAppConf("port").(int),
		"ip":   app.LogicApp.GetAppConf("master").(string),
	}
	t.Execute(rw, data) //执行模板的merger操作
}