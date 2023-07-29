package domain

import "time"

type Message struct {
	MessageId  int64     `gorm:"Column:message_id;AUTO_INCREMENT" json:"message_id"`
	UserId     int64     `gorm:"Column:user_id" json:"user_id"`
	ToUserId   int64     `gorm:"Column:to_user_id" json:"to_user_id"`
	Content    string    `gorm:"Column:content" json:"content"`
	ActionType string    `gorm:"Column:action_type" json:"action_type"`
	CreateTime time.Time `gorm:"Column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"Column:update_time" json:"update_time"`
}

func (Message) TableName() string {
	return "tb_message"
}
