package models

type UserModel struct {

}

// 用户列表
func (u *UserModel) GetUsers(page int, pageSize int, data map[string]interface{}) (users []Auth) {

	db.Where(data).Offset(page).Limit(pageSize).Find(&users)
	return
}

// 获取总数
func (u *UserModel) GetTotal(maps interface{}) (count int) {
	db.Model(&Auth{}).Where(maps).Count(&count)

	return
}