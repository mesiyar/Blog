package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"wechatNotify/pkg/e"
)

type Notes struct {

}

// 获取tags
//func getNotes(c *gin.Context) {
//	name := c.Query("id")
//	maps := make(map[string]interface{})
//	data := make(map[string]interface{})
//	if name != "" {
//		maps["name"] = name
//	}
//
//	var state int = -1
//
//	if arg := c.Query("state"); arg != "" {
//		state = com.StrTo(arg).MustInt()
//		maps["state"] = state
//	}
//
//	code := e.SUCCESS
//
//	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
//	data["total"] = models.GetTagTotal(maps)
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": data,
//	})
//}



// 添加文章标签
func AddNote(c *gin.Context) {
	name := c.Query("name")
	isStatus := 1
	parentId := 0
	noteType := c.Query("note_type")
	//noteTitle := c.Query("note_title")
	//noteContents := c.Query("note_contents")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Range(noteType, 0, 2, "note_type").Message("笔记类型 只能是 文件夹/笔记/markdown")

	code := e.InvalidParams
	if ! valid.HasErrors() {
		data := make(map[string]interface{})
		data["user_id"] = 0
		data["parent_id"] = parentId
		data["note_type"] = noteType
		data["is_status"] = isStatus
		code = e.SUCCESS
		models.AddNote(data)

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.InvalidParams
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.InvalidParams
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
