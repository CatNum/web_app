package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controller"
	"web_app/logger"
	"web_app/middlewares"
)

func SetUpRouter(mode string)( *gin.Engine){
	if mode == gin.ReleaseMode{
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	// 注册业务路由
	v1.POST("/signup",controller.SignUpHandler)
	// 登录
	v1.POST("/login",controller.LoginHandler)
	// 应用JWT认证中间件
	v1.Use(middlewares.JWTAuthMiddleware())

	{
		v1.GET("/community",controller.CommunityHandler)
		v1.GET("/communityDetail/:id",controller.CommunityDetailHandler)

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"404",
		})
	})
	return r
}

