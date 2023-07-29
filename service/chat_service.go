package service

import (
	"log"
	"mini-douyin/model/request"
	"mini-douyin/model/response"
	"mini-douyin/repository"
	"strconv"
)

type IChatService interface {
	PostMessage(request request.ChatPostRequest, userId int64) response.ChatPostResponse
	ListMessage(messageRequest request.ChatListMessageRequest, userId int64) response.ChatListMessageResponse
}

type ChatService struct {
	ChatRepository repository.IChatRepository
}

func (c ChatService) ListMessage(messageRequest request.ChatListMessageRequest, userId int64) response.ChatListMessageResponse {

	toUserID, _ := strconv.ParseInt(messageRequest.ToUserID, 10, 64)

	messages, err := c.ChatRepository.ListMessage(userId, toUserID, messageRequest.PreMsgTime)
	if err != nil {
		log.Printf("ListMessage|查询消息列表失败|%v", err)
		return response.ChatListMessageResponse{
			MessageList: nil,
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "get message list fail",
			}}
	}
	log.Printf("数据库消息列表为|%v", messages)
	var messageList []response.Message
	messageList = make([]response.Message, len(messages))
	for idx, message := range messages {
		messageList[idx].CreateTime = message.CreateTime.Unix()
		messageList[idx].ID = message.MessageId
		messageList[idx].ToUserID = message.ToUserId
		messageList[idx].FromUserID = message.UserId
		messageList[idx].Content = message.Content
	}

	log.Printf("返回消息列表为|%v", messageList)
	return response.ChatListMessageResponse{
		MessageList: messageList,
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		}}
}

func (c ChatService) PostMessage(request request.ChatPostRequest, userId int64) response.ChatPostResponse {

	content := request.Content
	toUserId := request.ToUserId
	actionType := request.ActionType

	err := c.ChatRepository.AddMessage(userId, toUserId, content, actionType)

	if err != nil {
		log.Printf("PostMessage|增加消息失败|%v", err)
		return response.ChatPostResponse{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "post message fail",
			},
		}
	}

	return response.ChatPostResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "post message success",
		},
	}

}

func NewChatService() IChatService {
	return ChatService{
		ChatRepository: repository.NewChatRepository(),
	}
}
