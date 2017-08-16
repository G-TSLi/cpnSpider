package main

import (
	"cpnSpider/web"
	_ "cpnSpider/lib/baidu"
	_ "cpnSpider/lib/qichacha"

)

func main()  {
	web.Run()
}