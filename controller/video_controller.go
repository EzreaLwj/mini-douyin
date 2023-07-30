package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mini-douyin/model/request"
	"mini-douyin/model/response"
	"mini-douyin/service"
	"mini-douyin/utils/jwt"
	"net/http"
)

type IVideoController interface {
	PostVideo(c *gin.Context) // 提交/发布视频
	ListVideo(c *gin.Context) // 展示视频列表
	FeedVideo(c *gin.Context) // 视频流接口
}

type VideoController struct {
	VideoService service.IVideoService
}

func NewVideoController() IVideoController {
	videoService := service.NewVideoService()
	videoController := VideoController{VideoService: videoService}
	return videoController
}

// PostVideo 发布视频
func (v VideoController) PostVideo(c *gin.Context) {
	// 判断参数
	var postRequest request.PostRequest
	err := c.ShouldBind(&postRequest)
	if err != nil {
		log.Printf("PostVideo|参数错误|%v", err)
		c.JSON(http.StatusOK, response.VideoPostResponse{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	token := postRequest.Token
	claims, err := jwt.ParseToken(token)
	if err != nil {
		log.Printf("PostVideo|token解析错误|%v", err)
		c.JSON(http.StatusUnauthorized, response.ErrorMessage{
			Response: response.Response{
				StatusMsg:  "unauthorized",
				StatusCode: 1,
			},
		})
		return
	}

	userID := claims.UserID
	videoPostResponse := v.VideoService.PostVideo(userID, &postRequest, c)
	c.JSON(http.StatusOK, videoPostResponse)
}

// ListVideo 展示视频
func (v VideoController) ListVideo(c *gin.Context) {
	var videoRequest request.ListRequest
	err := c.ShouldBindQuery(&videoRequest)
	if err != nil {
		log.Printf("ListVideo|请求参数错误|%v", err)
		c.JSON(http.StatusBadRequest, response.ErrorMessage{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "参数错误",
			},
		})
		return
	}
	video := v.VideoService.ListVideo(&videoRequest)
	c.JSON(http.StatusOK, video)
}

// FeedVideo 视频流接口
func (v VideoController) FeedVideo(c *gin.Context) {
	var videoRequest request.FeedRequest
	err := c.ShouldBindQuery(&videoRequest)
	if err != nil {
		log.Printf("ListVideo|请求参数错误|%v", err)
		c.JSON(http.StatusBadRequest, response.ErrorMessage{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "参数错误",
			},
		})
		return
	}
	video := v.VideoService.FeedVideo(c, &videoRequest)
	c.JSON(http.StatusOK, video)
}
