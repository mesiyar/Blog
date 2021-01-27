package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"wechatNotify/models"
	"wechatNotify/pkg/e"
	"wechatNotify/pkg/logging"
	"wechatNotify/pkg/setting"
	"wechatNotify/pkg/util"
)

type TagController struct {
}

var tagModel = new(models.TagModel)

func (t TagController) Add(c *gin.Context) {
	name := c.PostForm("name")
	createdBy := util.UserInfo.Username

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 200, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	logging.Info(name, createdBy)
	code := e.InvalidParams
	if ! valid.HasErrors() {
		if ! tagModel.ExistTagByName(name) {
			code = e.SUCCESS
			tagModel.AddTag(name, createdBy)
		} else {
			code = e.ErrorTagExist
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func (t TagController) Update(c *gin.Context) {
	name := c.PostForm("name")
	id := c.PostForm("id")
	modifiedBy := util.UserInfo.Username

	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 20, "name").Message("名称最长为20字符")

	code := e.InvalidParams
	if ! valid.HasErrors() {
		code = e.SUCCESS
		tagModel.UpdateTag(id, name, modifiedBy)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//
func (t TagController) Delete(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.InvalidParams
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if tagModel.ExistTagByID(id) {
			tagModel.DeleteTag(id)
		} else {
			code = e.ErrorTagNotExist
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func (t *TagController) List(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}

	var state int = -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = tagModel.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = tagModel.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
