package core

import "web.sjy.com/router"
import "web.sjy.com/config"

type Params struct {
	Addr  string
	Mysql bool
	Redis bool
}

func Run(p Params) {

	//加载配置文件
	config.Load()

	//初始化Mysql
	if p.Mysql {

	}

	//启动路由
	router.Gin.Run(p.Addr)
}
