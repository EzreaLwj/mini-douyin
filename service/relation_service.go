package service

import (
	"github.com/jinzhu/copier"
	"log"
	"mini-douyin/model/request"
	"mini-douyin/model/response"
	"mini-douyin/repository"
	"strconv"
)

type IRelationService interface {
	FollowAction(userId int64, request request.FollowActionRequest) response.FollowActionResponse // 关注操作
	GetFollow(followerRequest request.GetFollowerRequest) response.GetFollowerResponse            // 获取关注者
	GetFollower(followerRequest request.GetFollowerRequest) response.GetFollowerResponse          // 获取粉丝

}

type RelationService struct {
	RelationRepository repository.IRelationRepository
	UserRepository     repository.IUserRepository
}

// GetFollower 获取粉丝
func (r RelationService) GetFollower(followerRequest request.GetFollowerRequest) response.GetFollowerResponse {
	userId, err := strconv.ParseInt(followerRequest.UserId, 10, 64)
	if err != nil {
		log.Printf("GetFollow|格式转换失败|%v", err)
		return response.GetFollowerResponse{}
	}

	relations, err := r.RelationRepository.GetFollower(userId)
	if err != nil {
		log.Printf("GetFollow|数据库错误|%v", err)
		return response.GetFollowerResponse{}
	}

	var followerResponse response.GetFollowerResponse
	for _, relation := range relations {
		userId := relation.UserId
		user, err := r.UserRepository.GetUserById(userId)
		if err != nil {
			log.Printf("GetFollow|数据库错误|%v", err)
			return response.GetFollowerResponse{}
		}
		var responseUser response.User
		_ = copier.Copy(&responseUser, &user)

		followerResponse.UserList = append(followerResponse.UserList, responseUser)
	}

	followerResponse.Response = response.Response{
		StatusCode: 0,
		StatusMsg:  "success",
	}

	return followerResponse
}

// GetFollow 获取关注者
func (r RelationService) GetFollow(followerRequest request.GetFollowerRequest) response.GetFollowerResponse {
	userId, err := strconv.ParseInt(followerRequest.UserId, 10, 64)
	if err != nil {
		log.Printf("GetFollow|格式转换失败|%v", err)
		return response.GetFollowerResponse{}
	}

	relations, err := r.RelationRepository.GetFollow(userId)
	if err != nil {
		log.Printf("GetFollow|数据库错误|%v", err)
		return response.GetFollowerResponse{}
	}

	var followerResponse response.GetFollowerResponse
	for _, relation := range relations {
		toUserId := relation.ToUserId
		user, err := r.UserRepository.GetUserById(toUserId)
		if err != nil {
			log.Printf("GetFollow|数据库错误|%v", err)
			return response.GetFollowerResponse{}
		}
		var responseUser response.User
		_ = copier.Copy(&responseUser, &user)

		followerResponse.UserList = append(followerResponse.UserList, responseUser)
	}

	followerResponse.Response = response.Response{
		StatusCode: 0,
		StatusMsg:  "success",
	}

	return followerResponse
}

// FollowAction 关注操作
func (r RelationService) FollowAction(userId int64, followActionRequest request.FollowActionRequest) response.FollowActionResponse {
	toUserId, err := strconv.ParseInt(followActionRequest.ToUserId, 10, 64)
	if err != nil {
		log.Printf("FollowAction|格式转换失败|%v", err)
		return response.FollowActionResponse{}
	}
	err = r.RelationRepository.AddFollow(userId, toUserId)
	if err != nil {
		log.Printf("FollowAction|插入数据错误|%v", err)
		return response.FollowActionResponse{}
	}
	return response.FollowActionResponse{
		Response: response.Response{StatusCode: 0, StatusMsg: "success"},
	}
}

func NewRelationService() IRelationService {
	relationService := RelationService{RelationRepository: repository.NewRelationRepository(), UserRepository: repository.NewUserRepository()}
	return relationService
}
