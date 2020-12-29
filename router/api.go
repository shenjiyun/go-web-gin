package router

import (
	"web.sjy.com/controllers/admin"
	"web.sjy.com/middleware"
)

func apiRoute() {
	r := Gin.Group("/api")
	r.Use(middleware.Api)
	{
		r.GET("/", admin.Index{}.Hello)
	}
}
