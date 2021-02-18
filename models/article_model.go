package models

import (
	"time"
	"wechatNotify/pkg/logging"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
	ViewCount  int    `json:"view_count"`
	IsTop      int    `json:"is_top"` // 是否置顶
}

type ArticleModel struct {
}

func (a ArticleModel) ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	return article.ID > 0
}

func (a ArticleModel) GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func (a ArticleModel) GetArticlesByKeywords(pageNum int, pageSize int, keywords string) (articles []Article) {
	db.Preload("Tag").Where("title like ? and state = ?", "%"+keywords+"%", IsStatusEnable).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func (a ArticleModel) GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func (a ArticleModel) GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

func (a ArticleModel) EditArticle(id int, data interface{}) bool {
	logging.Info(data)
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func (a ArticleModel) AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
		IsTop:     data["is_top"].(int),
	})

	return true
}

func (a ArticleModel) DeleteArticle(id int) bool {
	data := make(map[string]interface{})
	data["state"] = IsStatusDisable
	data["deleted_on"] = time.Now().Unix()
	db.Model(&Article{}).Where("id = ?", id).Update(data)

	return true
}

func (a ArticleModel) GetTopArticle() (articles []Article) {
	db.Preload("Tag").Where("is_top = 1").Order("created_by desc").Limit(10).Find(&articles)

	return
}
