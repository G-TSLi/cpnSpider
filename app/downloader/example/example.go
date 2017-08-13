package main

import (
	"../surfer"
	"log"
)
func main()  {
	log.Printf("********************************************* surf内核GET下载测试开始 *********************************************")
	resp, err := surfer.Download(&surfer.DefaultRequest{
		Url: "http://www.baidu.com/",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("baidu resp.Status: %s\nresp.Header: %#v\n", resp.Status, resp.Header)
}
