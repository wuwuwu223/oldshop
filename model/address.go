package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Address   string `json:"address" gorm:"type:varchar(200);comment:'地址'"`
	Name      string `json:"name" gorm:"type:varchar(20);comment:'收货人'"`
	Phone     string `json:"phone" gorm:"type:varchar(20);comment:'手机号'"`
	IsDefault bool   `json:"is_default" gorm:"comment:'是否默认地址'"`
	UserID    uint   `json:"user_id" gorm:"comment:'用户ID'"`
}
