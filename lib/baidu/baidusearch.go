package lib

import (
	"cpnSpider/app/downloader/request"
	. "cpnSpider/app/spider"
	"cpnSpider/common/goquery"
	"regexp"
	"strconv"
	"math"
)

func init() {
	BaiduSearch.Register()
}

var BaiduSearch = &Spider{
	Name:        "百度搜索",
	Description: "百度搜索结果 [www.baidu.com]",
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.Aid(map[string]interface{}{"loop": [2]int{0, 1}, "Rule": "生成请求"}, "生成请求")
		},
		Trunk: map[string]*Rule{
			"生成请求": {
				AidFunc: func(ctx *Context, aid map[string]interface{}) interface{} {
					for loop := aid["loop"].([2]int); loop[0] < loop[1]; loop[0]++ {
						ctx.AddQueue(&request.Request{
							Url:        "http://www.baidu.com/s?ie=utf-8&nojc=1&wd=123&rn=50&pn=" + strconv.Itoa(50*loop[0]),
							Rule:       aid["Rule"].(string),
						})
					}
					return nil
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					total1 := query.Find(".nums").Text()
					re, _ := regexp.Compile(`[\D]*`)
					total1 = re.ReplaceAllString(total1, "")
					total2, _ := strconv.Atoi(total1)
					total := int(math.Ceil(float64(total2) / 50))
					// 调用指定规则下辅助函数
					ctx.Aid(map[string]interface{}{"loop": [2]int{1, total}, "Rule": "搜索结果"})
					// 用指定规则解析响应流
					ctx.Parse("搜索结果")
				},
			},

			"搜索结果": {
				//注意：有无字段语义和是否输出数据必须保持一致
				ItemFields: []string{
					"标题",
					"内容",
					"不完整URL",
					"百度跳转",
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					query.Find("#content_left .c-container").Each(func(i int, s *goquery.Selection) {
						title := s.Find(".t").Text()
						content := s.Find(".c-abstract").Text()
						re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
						title = re.ReplaceAllString(title, "")
						content = re.ReplaceAllString(content, "")
					})
				},
			},
		},
	},
}
