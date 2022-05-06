package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/logger"
)

func SetUp()( *gin.Engine){
	r := gin.Default()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	r.GET("/",func(c *gin.Context){
		c.String(http.StatusOK,"ok")
	})
	return r
}
