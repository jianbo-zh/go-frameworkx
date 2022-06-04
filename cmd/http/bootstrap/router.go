package bootstrap

import (
	"goframeworkx/cmd/http/route"

	"github.com/gin-gonic/gin"
)

var appRouter *gin.Engine = gin.New()

func initRouter() {
	route.Setup(appRouter)
}
