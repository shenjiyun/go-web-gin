package middleware

import "github.com/gin-gonic/gin"

func Api(c *gin.Context) {
	c.String(200, "中间件")
}