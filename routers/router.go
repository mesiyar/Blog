package routers

import (
	//_ "cloud-notes/src/docs"
	//"cloud-notes/src/middleware/jwt"
	//v1 "cloud-notes/src/routers/api/v1"
	"github.com/gin-gonic/gin"
	"wechatNotify/middleware/cros"
	"wechatNotify/middleware/jwt"
	"wechatNotify/pkg/setting"
	"wechatNotify/routers/api"
	v1 "wechatNotify/routers/api/v1"

	//ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.Use(cros.Cros())

	gin.SetMode(setting.RunMode)

	r.POST("/auth", api.GetAuth)

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//
	V1 := r.Group("/api/v1")

	V1.Use(jwt.JWT())
	{
		user := v1.UserController{}
		V1.GET("/users", user.GetUsers)
		V1.POST("/disable_account", api.DisableAuth)
		V1.POST("/create_account", api.CreateAuth)
		V1.POST("/logout", api.Logout)
		billing := v1.BillController{}
		V1.POST("/add_billing", billing.AddBilling)

		tag := v1.TagController{}
		V1.POST("/add_tag", tag.Add)
		V1.POST("/update_tag", tag.Update)
		V1.DELETE("/delete_tag", tag.Delete)
	}

	return r
}
