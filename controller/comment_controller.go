package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mini-douyin/model/request"
	"mini-douyin/service"
	"net/http"
)

type ICommentController interface {
	CommentPost(c *gin.Context) // 发表评论
	ListComment(c *gin.Context) // 获取评论列表
}

type CommentController struct {
	CommentService service.ICommentService
}

func (c2 CommentController) ListComment(c *gin.Context) {
	var listRequest request.CommentListRequest
	err := c.ShouldBindQuery(&listRequest)
	if err != nil {
		log.Printf("PostComment|参数错误|%v", listRequest)
		return
	}

	commentListResponse := c2.CommentService.ListComment(listRequest)
	c.JSON(http.StatusOK, commentListResponse)
}

func (c2 CommentController) CommentPost(c *gin.Context) {
	var postRequest request.CommentPostRequest
	err := c.ShouldBindQuery(&postRequest)
	if err != nil {
		log.Printf("PostComment|参数错误|%v", postRequest)
		return
	}

	commentPostResponse := c2.CommentService.CommentPost(postRequest)
	c.JSON(http.StatusOK, commentPostResponse)
}

func NewCommentController() ICommentController {
	return CommentController{CommentService: service.NewCommentService()}
}
