package route

import (
	usermodroute "goframeworkx/cmd/http/module/usermod/route"
	"goframeworkx/cmd/http/route/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {

	// 1. 注册全局中间件
	router.Use(
		middleware.LimitBodySize,
		middleware.Log,
		gin.Recovery(),
		middleware.Cors,
	)

	// 3. 注册路由
	usermodroute.Setup(router)
}
