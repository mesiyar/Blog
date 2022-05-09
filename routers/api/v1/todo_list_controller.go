package v1

import (
	"net/http"

	"wechatNotify/models"
	"wechatNotify/pkg/e"
	"wechatNotify/pkg/setting"
	"wechatNotify/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type TodolistController struct {
}
type addTodoListReq struct {
	Todo string `json:"todo" binding:"required"`
}

type updateTodoListReq struct {
	ID   int    `json:"id" binding:"required"`
	Todo string `json:"todo" binding:"required"`
}
type finishedTodoReq struct {
	ID int `json:"id" binding:"required"`
}

var todoListModel = new(models.TodoListModel)

func (t TodolistController) GetTodolist(c *gin.Context) {
	userId := util.UserInfo.ID
	query := make(map[string]interface{})
	query["created_by"] = userId
	res := todoListModel.GetTodoLists(util.GetPage(c), setting.PageSize, query)
	data := make(map[string]interface{})
	data["list"] = res
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": data,
	})
}

func (t TodolistController) Add(c *gin.Context) { // 新增待办事项
	var addTodo addTodoListReq
	c.BindJSON(&addTodo)
	todo := addTodo.Todo
	valid := validation.Validation{}
	valid.Required(todo, "todo").Message("待办事项不能为空")

	code := e.InvalidParams
	if !valid.HasErrors() {
		code = e.SUCCESS
		todoListModel.AddTodoList(todo, util.UserInfo.ID)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

func (t TodolistController) Delete(c *gin.Context) { // 删除待办事项
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.InvalidParams
	if !valid.HasErrors() {
		if todoListModel.ExistTodoListById(id) {
			code = e.SUCCESS
			todoListModel.DeleteTodoList(id)
		} else {
			code = e.ErrorTodoListNotExist
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func (t TodolistController) Update(c *gin.Context) { // 更新待办事项
	var updateTodo updateTodoListReq
	c.BindJSON(&updateTodo)
	id := updateTodo.ID
	todo := updateTodo.Todo
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")
	valid.Required(todo, "todo").Message("待办事项不能为空")

	code := e.InvalidParams
	if !valid.HasErrors() {
		if todoListModel.ExistTodoListById(id) {
			code = e.SUCCESS
			todoListModel.UpdateTodoList(id, todo, util.UserInfo.ID)
		} else {
			code = e.ErrorTodoListNotExist
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func (t TodolistController) UpdateFinishedStatus(c *gin.Context) { // 更新待办事项完成状态
	var finishedTodo finishedTodoReq
	c.BindJSON(&finishedTodo)
	id := finishedTodo.ID
	valid := validation.Validation{}
	valid.Required(id, "id").Message("id不能为空")

	code := e.InvalidParams
	if !valid.HasErrors() {
		if todoListModel.ExistTodoListById(id) {
			code = e.SUCCESS
			todoListModel.UpdateFinishedStatus(id)
		} else {
			code = e.ErrorTodoListNotExist
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
