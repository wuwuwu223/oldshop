package model

import "gorm.io/gorm"

//轮播图结构体
type HomePic struct {
	gorm.Model
	Pic   string `json:"pic" gorm:"type:varchar(255);comment:'图片地址'"`
	Index int    `json:"index" gorm:"type:int;comment:'图片排序'"`
}
