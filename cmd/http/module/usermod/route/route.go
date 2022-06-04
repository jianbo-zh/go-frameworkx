package usermod

import (
	apiroute "goframeworkx/cmd/http/module/usermod/route/api"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	apiroute.RegisterSwagger(router)
	apiroute.RegisterV1(router)
}
