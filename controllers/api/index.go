package api

import (
	"github.com/gin-gonic/gin"
	"web.sjy.com/core/db"
	"web.sjy.com/model"
)

func Index(ctx *gin.Context) {

	user := &model.User{}
	db.Mysql.Find(&user)
	ctx.JSON(200, user)
}
