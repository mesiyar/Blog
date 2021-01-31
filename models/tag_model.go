// 文章标签
package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

type TagModel struct {
}

func (tag *TagModel) GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func (tag *TagModel) GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func (tag *TagModel) ExistTagByName(name string) bool {
	var t Tag
	db.Select("id").Where("name = ? and state = ?", name, IsStatusEnable).First(&t)
	if t.ID > 0 {
		return true
	}

	return false
}

func (tag *TagModel) GetTagByName(name string) Tag {
	var t Tag
	db.Select("id").Where("name = ? and state = ?", name, IsStatusEnable).First(&t)
	return t
}


func (tag *TagModel) AddTag(name string, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     IsStatusEnable,
		CreatedBy: createdBy,
	})

	return true
}

func (tag *TagModel) UpdateTag(id string, name string, by string) bool {
	data := make(map[string]interface{})
	data["name"] = name
	data["modified_by"] = by
	db.Model(&Tag{}).Where("id = ?", id).Update(data)
	return true
}

func (tag *TagModel) ExistTagByID(id int) bool {
	var t Tag
	db.Select("id").Where("id = ?", id).First(&t)
	if t.ID > 0 {
		return true
	}

	return false
}

func (tag *TagModel) DeleteTag(id int) bool {
	data := make(map[string]interface{})
	data["state"] = IsStatusDisable
	db.Model(&Tag{}).Where("id = ?", id).Update(data)

	return true
}

func (tag *TagModel) GetAll() (tags []Tag) {
	db.Find(&tags)

	return
}
