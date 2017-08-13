package main
import (
	"github.com/henrylee2cn/surfer"
	"io/ioutil"
	"log"
)
func main() {
	// 默认使用surf内核下载
	resp, err := surfer.Download(&surfer.Request{
		Url: "http://github.com/henrylee2cn/surfer",
	})
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	log.Println(string(b), err)
	// 指定使用phantomjs内核下载
	resp, err = surfer.Download(&surfer.Request{
		Url:          "https://www.baidu.com/",
		DownloaderID: 1,
	})
	if err != nil {
		log.Fatal(err)
	}
	b, err = ioutil.ReadAll(resp.Body)
	log.Println(string(b), err)
	resp.Body.Close()
	surfer.DestroyJsFiles()
}