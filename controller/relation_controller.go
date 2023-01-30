package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"mini-douyin/model/request"
	"mini-douyin/service"
	"net/http"
)

type IRelationController interface {
	FollowAction(c *gin.Context) // 关注操作
	GetFollow(c *gin.Context)    // 获取关注的人
	GetFollower(c *gin.Context)  // 获取粉丝

}

type RelationController struct {
	RelationService service.IRelationService
}

// GetFollower 获取粉丝
func (r RelationController) GetFollower(c *gin.Context) {
	var getFollowerRequest request.GetFollowerRequest
	err := c.ShouldBindQuery(&getFollowerRequest)
	if err != nil {
		log.Printf("FollowAction|参数错误|%v", err)
		return
	}
	actionResponse := r.RelationService.GetFollower(getFollowerRequest)
	c.JSON(http.StatusOK, actionResponse)
}

// GetFollow 获取关注列表
func (r RelationController) GetFollow(c *gin.Context) {
	var getFollowerRequest request.GetFollowerRequest
	err := c.ShouldBindQuery(&getFollowerRequest)
	if err != nil {
		log.Printf("FollowAction|参数错误|%v", err)
		return
	}
	actionResponse := r.RelationService.GetFollow(getFollowerRequest)
	c.JSON(http.StatusOK, actionResponse)
}

// FollowAction 关注操作
func (r RelationController) FollowAction(c *gin.Context) {
	var followActionRequest request.FollowActionRequest
	err := c.ShouldBindQuery(&followActionRequest)
	if err != nil {
		log.Printf("FollowAction|参数错误|%v", err)
		return
	}
	value, _ := c.Get("userId")

	actionResponse := r.RelationService.FollowAction(value.(int64), followActionRequest)
	c.JSON(http.StatusOK, actionResponse)
}

func NewRelationController() IRelationController {
	relationController := RelationController{RelationService: service.NewRelationService()}
	return relationController
}
