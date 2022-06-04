package api

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/jianbo-zh/gin-swagger"
	"github.com/jianbo-zh/gin-swagger/swaggerFiles"

	_ "goframeworkx/cmd/http/module/usermod/doc"
)

func RegisterSwagger(router *gin.Engine) {
	router.GET("/swagger/user/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		// ginSwagger.URL("doc.json"),
		ginSwagger.InstanceName("user"), // swag init --instanceName xxx
		ginSwagger.DocExpansion("list"),
		ginSwagger.DeepLinking(true),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))
}
