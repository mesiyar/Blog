package routers

import (
	_ "wechatNotify/docs"
	//"cloud-notes/src/middleware/jwt"
	//v1 "cloud-notes/src/routers/api/v1"
	"github.com/gin-gonic/gin"
	"wechatNotify/middleware/cros"
	"wechatNotify/middleware/jwt"
	"wechatNotify/pkg/setting"
	"wechatNotify/routers/api"
	v1 "wechatNotify/routers/api/v1"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.Use(cros.Cros())

	gin.SetMode(setting.RunMode)

	r.POST("/auth", api.GetAuth)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//
	admin := r.Group("/admin")
	user := v1.UserController{}

	admin.Use(jwt.JWT())
	{
		admin.GET("/users", user.GetUsers)
		admin.POST("/disable_account", api.DisableAuth)
		admin.POST("/create_account", api.CreateAuth)
		admin.POST("/logout", api.Logout)
		admin.GET("/get_user_info", user.GetUserInfo)
		tag := v1.TagController{}
		admin.GET("/tags", tag.List)
		admin.POST("/add_tag", tag.Add)
		admin.POST("/update_tag", tag.Update)
		admin.DELETE("/delete_tag", tag.Delete)
		admin.GET("/all_tags", tag.AllTags)
		billing := v1.BillController{}
		admin.POST("/add_billing", billing.AddBilling)
		// 文章管理
		article := v1.ArticleController{}
		admin.GET("/article", article.GetArticle)
		admin.GET("/articles", article.GetArticles)
		admin.POST("/add_article", article.AddArticle)
		admin.POST("/update_article", article.EditArticle)
		admin.DELETE("/delete_article", article.DeleteArticle)
	}

	return r
}
