package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
	"wechatNotify/models"
	"wechatNotify/pkg/cache"
	"wechatNotify/pkg/e"
	"wechatNotify/pkg/setting"
	"wechatNotify/pkg/util"
)

type ArticleController struct {
}

var articleModel = new(models.ArticleModel)

//获取单个文章
func (a ArticleController) GetArticle(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.InvalidParams
	var data interface{}
	if ! valid.HasErrors() {
		if d, ok := cache.Exists(id); !ok {
			if articleModel.ExistArticleByID(id) {
				data = articleModel.GetArticle(id)
				code = e.SUCCESS
				cache.Set(id, data)
			} else {
				code = e.ErrorArticleNotExist
			}
		} else {
			code = e.SUCCESS
			data = d
		}

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func (a ArticleController) GetByKeywords(c *gin.Context)  {
	data := make(map[string]interface{})

	keywords := c.Query("keywords")
	code := e.SUCCESS

	code = e.SUCCESS
	data["lists"] = articleModel.GetArticlesByKeywords(util.GetPage(c), setting.PageSize, keywords)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//获取多个文章
func (a ArticleController) GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	maps["state"] = models.IsStatusEnable
	maps["is_pub"] = 1

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	if arg := c.Query("tag_name"); arg != "" {
		tag := tagModel.GetTagByName(arg)
		maps["tag_id"] = tag.ID
	}

	code := e.InvalidParams
	if ! valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = articleModel.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = articleModel.GetArticleTotal(maps)

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章
func (a ArticleController) AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	isTop := c.PostForm("is_top")
	isPub := c.PostForm("is_pub")
	createdBy := util.UserInfo.Username
	state := models.IsStatusEnable

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Required(isTop, "is_top").Message("是否置顶不能为空")

	code := e.InvalidParams
	if ! valid.HasErrors() {
		if tagModel.ExistTagByID(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state
			data["is_top"] = isPub

			articleModel.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ErrorTagNotExist
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

//修改文章
func (a ArticleController) EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("id")).MustInt()
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	isPub := c.PostForm("is_pub")
	modifiedBy := util.UserInfo.Username

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	var isTop int = 0
	if arg1 := c.PostForm("is_top"); arg1 != "" {
		isTop = com.StrTo(arg1).MustInt()
		valid.Range(isTop, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.InvalidParams
	if ! valid.HasErrors() {
		if articleModel.ExistArticleByID(id) {
			if tagModel.ExistTagByID(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}

				data["is_top"] = isTop

				data["is_pub"] = isPub

				data["modified_by"] = modifiedBy

				articleModel.EditArticle(id, data)
				cache.Set(id, data)
				code = e.SUCCESS
			} else {
				code = e.ErrorTagNotExist
			}
		} else {
			code = e.ErrorArticleNotExist
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除文章
func (a ArticleController) DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.InvalidParams
	if ! valid.HasErrors() {
		if articleModel.ExistArticleByID(id) {
			articleModel.DeleteArticle(id)
			code = e.SUCCESS
			cache.Del(id)
		} else {
			code = e.ErrorArticleNotExist
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func (a ArticleController) GetTopArticle(c *gin.Context) {

	code := e.SUCCESS
	var data interface{}

	data = articleModel.GetTopArticle()

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
