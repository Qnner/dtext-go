package main

import (
	"dtext-go/model"
	"dtext-go/utils"
	"embed"
)

// 使用embed 将静态资源嵌入go可执行文件中
//go:embed templates/*
var temp embed.FS

//go:embed static/*
var static embed.FS

func main() {
	// 读取默认仅读的页面，并写入sqlite
	model.WriteDefaultPages(utils.RelativePath("defaultPages.json"))
	server := model.InitServer(temp, static)
	
	utils.Log.Info("server run!")

	// 默认端口 5252
	server.Run(":5252")
}
