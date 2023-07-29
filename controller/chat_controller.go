package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mini-douyin/model/request"
	"mini-douyin/service"
	"net/http"
)

type IChatController interface {
	ChatPost(c *gin.Context)        // 发送消息
	ListChatMessage(c *gin.Context) // 聊天记录列表
}

type ChatController struct {
	ChatService service.IChatService
}

func (c2 ChatController) ListChatMessage(c *gin.Context) {
	var chatListMessageRequest request.ChatListMessageRequest

	err := c.ShouldBindQuery(&chatListMessageRequest)
	if err != nil {
		log.Printf("ChatPost|参数错误|%v", err.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	userId, _ := c.Get("userId")
	chatListMessageResponse := c2.ChatService.ListMessage(chatListMessageRequest, userId.(int64))
	c.JSON(http.StatusOK, chatListMessageResponse)
}

func (c2 ChatController) ChatPost(c *gin.Context) {
	var chatPostRequest request.ChatPostRequest

	err := c.ShouldBindQuery(&chatPostRequest)
	if err != nil {
		log.Printf("ChatPost|参数错误|%v", chatPostRequest)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	userId, _ := c.Get("userId")
	chatPostResponse := c2.ChatService.PostMessage(chatPostRequest, userId.(int64))
	c.JSON(http.StatusOK, chatPostResponse)

}

func NewChatController() IChatController {
	return ChatController{ChatService: service.NewChatService()}
}
