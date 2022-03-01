package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	
	//r.POST("/user/info", )

	return r
}