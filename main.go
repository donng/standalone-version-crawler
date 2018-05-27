package main

import (
	"crawler/standalone-version-crawler/engine"
	"crawler/standalone-version-crawler/zhenai/parser"
)

const url = "http://www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
