package main

import (
	"web.sjy.com/core"
)

func main() {
	core.Run(core.Params{
		Addr:  ":9090",
		Redis: true,
		Mysql: true,
	})
}
