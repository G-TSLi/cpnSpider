package main

import (
	"cpnSpider/web"
	//加载爬虫规则库
	_ "cpnSpider/lib/baidu"
	_ "cpnSpider/lib/qichacha"
)

func main()  {
	web.Run()
}