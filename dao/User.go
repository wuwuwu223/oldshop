package dao

import (
	"oldshop/global"
	"oldshop/model"
)

var db = global.DB

//根据ID查找用户
func FindUserById(user *model.User) error {
	return db.Where("id = ?", user.ID).First(user).Error
}

//创建用户
func CreateUser(user *model.User) error {
	return db.Create(user).Error
}

//更新用户
func UpdateUser(user *model.User) error {
	return db.Save(user).Error
}

//根据用户名查找用户
func FindUserByUsername(user *model.User) error {
	return db.Where("username = ?", user.Username).First(user).Error
}
