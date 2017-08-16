package qichacha

import (
	"cpnSpider/app/downloader/request"
	. "cpnSpider/app/spider"
	"cpnSpider/common/goquery"
	"log"
	"strings"
	"regexp"
	"strconv"
)

func init() {
	QccSpider.Register()
}

var QccSpider = &Spider{
	Name:        "企查查搜索",
	Description: "企查查搜索结果 [qichacha.com]",
	RuleTree: 	&RuleTree{
		Root:func(ctx *Context) {
			ctx.Aid(map[string]interface{}{"Rule": "判断页数"}, "判断页数")
		},
		Trunk:map[string]*Rule{
			"判断页数":{
				AidFunc: func(ctx *Context, aid map[string]interface{}) interface{} {
					ctx.AddQueue(
						&request.Request{
							Url:  "https://www.qichacha.com/search?key=%E5%90%8C%E8%8A%B1%E9%A1%BA#p:1&",
							Rule: aid["Rule"].(string),
						},
					)
					return nil
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					pageCount := 0
					query.Find("script").Each(func(i int, s *goquery.Selection) {
						if strings.Contains(s.Text(), "page_count") {
							re, _ := regexp.Compile(`page_count:"\d{1,}"`)
							temp := re.FindString(s.Text())
							re, _ = regexp.Compile(`\d{1,}`)
							temp2 := re.FindString(temp)
							pageCount, _ = strconv.Atoi(temp2)
						}
					})
					ctx.Aid(map[string]interface{}{"PageCount": pageCount}, "生成请求")
				},
			},
			"生成请求": {
				AidFunc: func(ctx *Context, aid map[string]interface{}) interface{} {
					return nil
				},
			},
			"搜索结果": {
				ItemFields:[]string{
					"商标",
					"公司名",
					"法定代表人",
					"注册资本",
					"成立时间",
					"电话",
					"邮箱",
					"地址",
					"状态",
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					query.Find("#searchlist > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
						// 获取公司名
						name := s.Find("a.ma_h1").Text()
						log.Println(name)
					})
				},
			},
		},
	},
}