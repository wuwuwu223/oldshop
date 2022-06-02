package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Message        string `json:"message" gorm:"comment:'消息内容'"`
	IsRead         bool   `json:"is_read" gorm:"comment:'是否已读'"`
	ConversationID uint   `json:"conversation_id"`
}
type Conversation struct {
	gorm.Model
	FromID   uint      `json:"from_id" gorm:"comment:'发送人'"`
	ToID     uint      `json:"to_id" gorm:"comment:'接收人'"`
	GoodID   uint      `json:"good_id" gorm:"comment:'最后关联商品ID'"`
	Messages []Message `json:"messages"`
}
