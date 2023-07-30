package service

import (
	"github.com/jinzhu/copier"
	"log"
	"mini-douyin/config"
	"mini-douyin/model/request"
	"mini-douyin/model/response"
	"mini-douyin/repository"
	"strconv"
)

type IRelationService interface {
	FollowAction(userId int64, request request.FollowActionRequest) response.FollowActionResponse    // 关注操作
	GetFollow(followerRequest request.GetFollowerRequest) response.GetFollowerResponse               // 获取关注者
	GetFollower(followerRequest request.GetFollowerRequest) response.GetFollowerResponse             // 获取粉丝
	GetFriendList(userId int64, request request.GetFriendListRequest) response.GetFriendListResponse // 获取朋友列表

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
		toUserId := relation.UserId
		user, err := r.UserRepository.GetUserById(toUserId)
		if err != nil {
			log.Printf("GetFollow|数据库错误|%v", err)
			return response.GetFollowerResponse{}
		}

		var responseUser response.User
		_ = copier.Copy(&responseUser, &user)
		responseUser.TotalFavorited = strconv.FormatInt(user.TotalFavorited, 10)
		follow := r.RelationRepository.CheckIsFollow(userId, toUserId)
		responseUser.IsFollow = follow
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
		responseUser.TotalFavorited = strconv.FormatInt(user.TotalFavorited, 10)
		responseUser.IsFollow = r.RelationRepository.CheckIsFollow(userId, toUserId)
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

	begin := config.DB.Begin()
	// 关注
	if followActionRequest.ActionType == "1" {
		err = r.RelationRepository.AddFollow(userId, toUserId)
		if err != nil {
			log.Printf("FollowAction|插入数据错误|%v", err)
			begin.Rollback()
			return response.FollowActionResponse{
				Response: response.Response{StatusCode: 1, StatusMsg: "insert fail"},
			}
		}

		// 增加用户的关注数
		r.UserRepository.UpdateFollowCount(userId, true)

		// 增加用户的获赞总数
		r.UserRepository.UpdateFollowerCount(toUserId, true)
	}

	// 取消关注
	if followActionRequest.ActionType == "2" {
		err = r.RelationRepository.RemoveFollow(userId, toUserId)
		if err != nil {
			log.Printf("FollowAction|删除数据错误|%v", err)
			begin.Rollback()
			return response.FollowActionResponse{
				Response: response.Response{StatusCode: 1, StatusMsg: "delete fail"},
			}
		}

		// 减少用户的关注数
		r.UserRepository.UpdateFollowCount(userId, false)

		// 减少被关注者的粉丝数
		r.UserRepository.UpdateFollowerCount(toUserId, false)
	}

	begin.Commit()

	return response.FollowActionResponse{
		Response: response.Response{StatusCode: 0, StatusMsg: "success"},
	}
}

// GetFriendList 获取朋友列表
func (r RelationService) GetFriendList(userId int64, getFriendListRequest request.GetFriendListRequest) response.GetFriendListResponse {
	userId, err := strconv.ParseInt(getFriendListRequest.UserId, 10, 64)
	if err != nil {
		log.Printf("GetFollow|格式转换失败|%v", err)
		return response.GetFriendListResponse{}
	}

	relations, err := r.RelationRepository.GetFollower(userId)

	if err != nil {
		log.Printf("GetFollow|数据库错误|%v", err)
		return response.GetFriendListResponse{}
	}

	var getFriendListResponse response.GetFriendListResponse
	for _, relation := range relations {
		userId := relation.UserId
		user, err := r.UserRepository.GetUserById(userId)
		if err != nil {
			log.Printf("GetFollow|数据库错误|%v", err)
			return response.GetFriendListResponse{}
		}
		var responseUser response.User
		_ = copier.Copy(&responseUser, &user)

		getFriendListResponse.UserList = append(getFriendListResponse.UserList, responseUser)
	}

	getFriendListResponse.Response = response.Response{
		StatusCode: 0,
		StatusMsg:  "success",
	}

	return getFriendListResponse

}
func NewRelationService() IRelationService {
	relationService := RelationService{RelationRepository: repository.NewRelationRepository(), UserRepository: repository.NewUserRepository()}
	return relationService
}
