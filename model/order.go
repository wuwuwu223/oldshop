package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	GoodID    uint      `json:"good_id" gorm:"comment:'商品ID'"`
	BuyerID   uint      `json:"buyer_id" gorm:"comment:'买家ID'"`
	SellerID  uint      `json:"seller_id" gorm:"comment:'卖家ID'"`
	Price     int64     `json:"price" gorm:"comment:'订单价格'"`
	Address   string    `json:"address" gorm:"type:varchar(200);comment:'收货地址'"`
	Name      string    `json:"name" gorm:"type:varchar(20);comment:'收货人'"`
	Phone     string    `json:"phone" gorm:"type:varchar(20);comment:'收货手机号'"`
	Logistics Logistics `json:"logistics" gorm:"comment:'物流信息'"`
	Status    uint8     `json:"status" gorm:"comment:'订单状态;0.待支付;1.已支付;2.待收货;3.已完成;4.已取消'"`
}
