package main

import (
	"polygon/configer"
	_ "polygon/driver/es"
)

func main() {
	// 初始化配置

	configer.InitConfig()
	//fmt.Println(configer.Get("es"))
	// 启动监听
}
