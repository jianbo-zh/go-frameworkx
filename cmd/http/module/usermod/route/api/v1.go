package api

import (
	"goframeworkx/cmd/http/module/usermod/api/userapi"
	"goframeworkx/cmd/http/route/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterV1(router *gin.Engine) {

	guestRouter := router.Group("/user/v1", middleware.Guest)
	{
		guestRouter.POST("/login", userapi.Login)
	}
}
