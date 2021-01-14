package admin

import (
	"github.com/gin-gonic/gin"
	"web.sjy.com/core/db"
	"web.sjy.com/model"
)

func init() {

}

type Index struct {
}

func (i Index) Hello(ctx *gin.Context) {
	users := &model.Users{}
	db.Mysql.First(&users)
	ctx.JSON(200, users)
}
