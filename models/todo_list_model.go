package models

import (
	"time"
	"wechatNotify/pkg/logging"
	"wechatNotify/pkg/util"
)

type TodoList struct {
	Model

	Todo           string `json:"todo"`
	State          int    `json:"state"`
	FinishedOn     int    `json:"finished_on"`
	FinishedStatus int    `json:"finished_status"`
	CreatedBy      int    `json:"created_by"`
}

const (
	FinishedStatusNot = iota
	FindishedStatusYes
)

type TodoListModel struct {
}

func (todo *TodoListModel) GetTodoLists(pageNum int, pageSize int, maps interface{}) (todos []TodoList) {

	if err := db.Where(maps).
		Where("state = ?", IsStatusEnable).
		Where("created_on >= ?", util.TodayZeroTime()).
		Offset(pageNum).Limit(pageSize).Find(&todos).Error; err != nil {
		logging.Error("get todo list error: %v", err)
		return nil
	}
	return
}

func (t *TodoListModel) GetTodoListTotal(maps interface{}) (count int) {
	db.Model(&TodoList{}).Where(maps).Count(&count)
	return
}

func (t *TodoListModel) AddTodoList(todo string, createdBy int) bool {
	db.Create(&TodoList{
		Todo:      todo,
		State:     IsStatusEnable,
		CreatedBy: createdBy,
	})
	return true
}

func (t *TodoListModel) UpdateTodoList(id int, todo string, updatedBy int) bool {
	data := make(map[string]interface{})
	data["todo"] = todo
	data["modified_by"] = updatedBy
	db.Model(&TodoList{}).Where("id = ?", id).Update(data)
	return true
}

func (t *TodoListModel) DeleteTodoList(id int) bool {
	data := make(map[string]interface{})
	data["state"] = IsStatusDisable
	db.Model(&TodoList{}).Where("id = ?", id).Update(data)
	return true
}

func (t *TodoListModel) UpdateFinishedStatus(id int) bool {
	data := make(map[string]interface{})
	data["finished_on"] = time.Now().Unix()
	data["finished_status"] = FindishedStatusYes
	db.Model(&TodoList{}).Where("id = ?", id).Update(data)

	return true
}

func (t *TodoListModel) ExistTodoListById(id int) bool {
	var todoList TodoList
	db.Select("id").Where("id = ?", id).First(&todoList)
	return todoList.ID > 0
}
