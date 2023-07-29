package repository

import (
	"mini-douyin/config"
	"mini-douyin/model/domain"
	"strconv"
	"time"
)

type IChatRepository interface {
	AddMessage(userId int64, toUserId string, content string, actionType string) error
	ListMessage(userId int64, toUserID int64, preMsgTime int64) ([]domain.Message, error)
}

type ChatRepository struct {
}

func (c ChatRepository) ListMessage(userId int64, toUserID int64, preMsgTime int64) ([]domain.Message, error) {

	var messages []domain.Message
	timeObj := time.Unix(preMsgTime, 0)
	formattedTime := timeObj.Format("2006-01-02 15:04:05")
	err := config.DB.Table("tb_message").Where("user_id = ? and to_user_id = ? and create_time > ? ", userId, toUserID, formattedTime).Order("create_time").Find(&messages).Error
	return messages, err
}

func (c ChatRepository) AddMessage(userId int64, toUserId string, content string, actionType string) error {

	var message domain.Message
	message.UserId = userId
	message.ToUserId, _ = strconv.ParseInt(toUserId, 10, 64)
	message.Content = content
	message.ActionType = actionType
	message.CreateTime = time.Now()
	message.UpdateTime = time.Now()
	err := config.DB.Table("tb_message").Create(&message).Error
	return err
}

func NewChatRepository() IChatRepository {
	return ChatRepository{}
}
