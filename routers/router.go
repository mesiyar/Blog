package routers

import (
	"net/http"
	_ "wechatNotify/docs"
	"wechatNotify/pkg/upload"

	"wechatNotify/middleware/cros"
	"wechatNotify/middleware/jwt"
	"wechatNotify/pkg/setting"
	"wechatNotify/routers/api"
	v1 "wechatNotify/routers/api/v1"

	"github.com/gin-gonic/gin"

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
	tag := v1.TagController{}
	r.GET("/all_tags", tag.AllTags)
	article := v1.ArticleController{}
	config := v1.ConfigController{}
	r.GET("/article", article.GetArticle)
	r.GET("/articles", article.GetArticles)
	r.GET("/search_article", article.GetByKeywords)
	r.GET("/top_articles", article.GetTopArticle)
	r.GET("/site_config", config.SiteConfig)
	// 静态文件目录
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	admin.Use(jwt.JWT())
	{
		admin.GET("/users", user.GetUsers)
		admin.POST("/disable_account", api.DisableAuth)
		admin.POST("/create_account", api.CreateAuth)
		admin.POST("/logout", api.Logout)
		admin.GET("/get_user_info", user.GetUserInfo)

		admin.GET("/tags", tag.List)
		admin.POST("/add_tag", tag.Add)
		admin.POST("/update_tag", tag.Update)
		admin.DELETE("/delete_tag", tag.Delete)
		admin.GET("/all_tags", tag.AllTags)
		billing := v1.BillController{}
		admin.POST("/add_billing", billing.AddBilling)
		// 文章管理
		admin.GET("/article", article.GetArticle)
		admin.GET("/articles", article.GetArticles)
		admin.POST("/add_article", article.AddArticle)
		admin.POST("/update_article", article.EditArticle)
		admin.DELETE("/delete_article", article.DeleteArticle)
		// 配置管理
		admin.GET("/configs", config.GetConfigs)
		admin.POST("/add_config", config.AddConfig)
		admin.POST("/update_config", config.EditConfig)
		admin.DELETE("/delete_config", config.DeleteConfig)

		// 待办事项
		todo := v1.TodolistController{}
		admin.GET("/todolist", todo.GetTodolist)
		admin.POST("/add_todolist", todo.Add)
		admin.POST("/update_todolist", todo.Update)
		admin.DELETE("/delete_todolist", todo.Delete)
		admin.POST("/update_todolist_status", todo.UpdateFinishedStatus)
	}

	return r
}
