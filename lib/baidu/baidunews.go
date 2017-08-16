package lib

import (
	. "cpnSpider/app/spider"           //必需
)

func init() {
	BaiduNews.Register()
}

var rss_BaiduNews = map[string]string{
	"国内最新":  "http://news.baidu.com/n?cmd=4&class=civilnews&tn=rss",
	"国际最新":  "http://news.baidu.com/n?cmd=4&class=internews&tn=rss",
	"军事最新":  "http://news.baidu.com/n?cmd=4&class=mil&tn=rss",
	"财经最新":  "http://news.baidu.com/n?cmd=4&class=finannews&tn=rss",
	"互联网最新": "http://news.baidu.com/n?cmd=4&class=internet&tn=rss",
	"房产最新":  "http://news.baidu.com/n?cmd=4&class=housenews&tn=rss",
	"汽车最新":  "http://news.baidu.com/n?cmd=4&class=autonews&tn=rss",
	"体育最新":  "http://news.baidu.com/n?cmd=4&class=sportnews&tn=rss",
	"娱乐最新":  "http://news.baidu.com/n?cmd=4&class=enternews&tn=rss",
	"游戏最新":  "http://news.baidu.com/n?cmd=4&class=gamenews&tn=rss",
	"教育最新":  "http://news.baidu.com/n?cmd=4&class=edunews&tn=rss",
	"女人最新":  "http://news.baidu.com/n?cmd=4&class=healthnews&tn=rss",
	"科技最新":  "http://news.baidu.com/n?cmd=4&class=technnews&tn=rss",
	"社会最新":  "http://news.baidu.com/n?cmd=4&class=socianews&tn=rss",
}

var BaiduNews = &Spider{
	Name:		"百度RSS新闻",
	Description: "百度RSS新闻，实现轮询更新 [news.baidu.com]",
	//RuleTree:	&RuleTree{
	//	Root: func(ctx *Context) {
	//	},
	//	Trunk: map[string]*Rule{
	//		"新闻详情":{
	//			ItemFields: []string{
	//				"标题",
	//				"描述",
	//				"内容",
	//				"发布时间",
	//				"分类",
	//				"作者",
	//			},
	//			ParseFunc: func(ctx *Context) {
	//
	//			},
	//		},
	//	},
	//},
}





