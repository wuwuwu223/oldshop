package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string         `json:"username" gorm:"type:varchar(100);uniqueIndex;comment:'用户名'"`
	Password    string         `json:"password" gorm:"type:varchar(100);comment:'密码'"`
	Phone       string         `json:"phone" gorm:"type:varchar(20);uniqueIndex;comment:'手机号'"`
	Avatar      string         `json:"avatar" gorm:"type:longtext;comment:'头像'"`
	Role        string         `json:"role" gorm:"default:user;type:varchar(20);comment:'角色'"`
	RealProfile RealProfile    `json:"real_profile" gorm:"comment:'实名信息'"`
	Addresses   []Address      `json:"addresses"`
	Goods       []Good         `json:"goods"`
	BuyOrders   []Order        `json:"buy_orders" gorm:"foreignKey:BuyerID"`
	SellOrders  []Order        `json:"sell_orders" gorm:"foreignKey:SellerID"`
	CommentTo   []Conversation `json:"comment_to" gorm:"foreignKey:ToID"`
	CommentFrom []Conversation `json:"comment_from" gorm:"foreignKey:FromID"`
}

type RealProfile struct {
	gorm.Model
	RealName string `json:"real_name" gorm:"type:varchar(100);comment:'名字'"`
	IDCard   string `json:"id_card" gorm:"type:varchar(30);uniqueIndex;comment:'身份证号'"`
	Verified int8   `json:"verified" gorm:"comment:'状态;0.未验证;1.已验证;2.已取消'"`
	UserID   uint   `json:"user_id" gorm:"comment:'用户ID'"`
}
