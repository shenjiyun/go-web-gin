package router

import (
	"web.sjy.com/controllers/api"
	"web.sjy.com/router/middleware"
)

func apiRoute() {
	r := Gin.Group("/api")
	r.Use(middleware.Api)
	{
		r.GET("/", api.Index)
	}
}
