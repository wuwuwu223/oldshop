package model

import "gorm.io/gorm"

type Logistics struct {
	gorm.Model
	OrderID uint   `json:"order_id" gorm:"comment:'订单ID'"`
	Company string `json:"company" gorm:"type:varchar(255);comment:'物流公司'"`
	No      string `json:"no" gorm:"type:varchar(255);comment:'物流单号'"`
}
