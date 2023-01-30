package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mini-douyin/model/request"
	"mini-douyin/model/response"
	"mini-douyin/service"
	"net/http"
)

type IFavoriteController interface {
	FavoriteAction(c *gin.Context)    // 点赞操作
	FavoriteVideoList(c *gin.Context) // 获取点赞列表
}

type FavoriteController struct {
	FavoriteService service.IFavoriteService
}

func NewFavoriteController() IFavoriteController {
	favoriteController := FavoriteController{FavoriteService: service.NewFavoriteService()}
	return favoriteController
}

// FavoriteAction 点赞操作
func (f FavoriteController) FavoriteAction(c *gin.Context) {

	var favoriteRequest request.FavoriteActionRequest
	err := c.ShouldBindQuery(&favoriteRequest)
	if err != nil {
		log.Printf("FavoriteAction|参数错误|%v", err)
		return
	}

	value, _ := c.Get("userId")
	f.FavoriteService.FavoriteAction(value.(int64), favoriteRequest)

	c.JSON(http.StatusOK, response.FavoriteActionResponse{
		Response: response.Response{StatusCode: 0, StatusMsg: "success"},
	})
}

// FavoriteVideoList 获取用户点赞的视频
func (f FavoriteController) FavoriteVideoList(c *gin.Context) {
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
	video := f.FavoriteService.FavoriteVideoList(&videoRequest)
	c.JSON(http.StatusOK, video)
}
