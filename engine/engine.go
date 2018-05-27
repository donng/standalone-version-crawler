package engine

import (
	"crawler/standalone-version-crawler/fetcher"
	"log"
)

func Run(seeds ...Request)  {
	// 生成请求的队列
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}


	for len(requests) > 0 {
		// 取出第一个 request 请求
		r := requests[0]
		requests = requests[1:]

		log.Printf("fetching url %s", r.Url)
		// 获得 url 内容
		content, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetchting url %s, %v", r.Url, err)
			continue
		}

		// 解析 url 内容
		parseResult := r.ParserFunc(content)

		// 将新的 request 添加到队列中
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got items %v", item)
		}
	}
}