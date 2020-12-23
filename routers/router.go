package routers

import (
	//_ "cloud-notes/src/docs"
	//"cloud-notes/src/middleware/jwt"
	//v1 "cloud-notes/src/routers/api/v1"
	"github.com/gin-gonic/gin"
	"wechatNotify/pkg/setting"
	"wechatNotify/routers/api"

	//ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)
	r.GET("/create_account", api.CreateAuth)
	r.POST("/disable_account", api.DisableAuth)

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//apiv1 := r.Group("/api/v1")
	//
	//apiv1.Use(jwt.JWT())
	//{
	//
	//}

	return r
}
