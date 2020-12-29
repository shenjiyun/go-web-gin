package admin

import "github.com/gin-gonic/gin"

func init() {

}

type Index struct {
}

func (i Index) Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"text": "hello222",
	})
}
