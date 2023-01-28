package service

import (
	"github.com/jinzhu/copier"
	"log"
	"mini-douyin/model/domain"
	"mini-douyin/model/request"
	"mini-douyin/model/response"
	"mini-douyin/repository"
	"strconv"
	"time"
)

type ICommentService interface {
	CommentPost(commentPostRequest request.CommentPostRequest) response.CommentPostResponse
	ListComment(commentListRequest request.CommentListRequest) response.CommentListResponse
}

type CommentService struct {
	CommentRepository repository.ICommentRepository
	UserRepository    repository.IUserRepository
}

// ListComment 获取评论列表
func (c CommentService) ListComment(commentListRequest request.CommentListRequest) response.CommentListResponse {
	videoId := commentListRequest.VideoId
	parseInt, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		log.Printf("ListComment|类型转换错误|%v", err)
		return response.CommentListResponse{}
	}
	commentList, err := c.CommentRepository.GetCommentList(parseInt)
	if err != nil {
		log.Printf("ListComment|数据库获取错误|%v", err)
		return response.CommentListResponse{}
	}

	var responseComment []response.Comment
	err = copier.Copy(&responseComment, &commentList)
	if err != nil {
		log.Printf("ListComment|获取用户失败|%v", err)
		return response.CommentListResponse{}
	}

	for index, comment := range commentList {
		userId := comment.UserId

		user, err := c.UserRepository.GetUserById(userId)
		if err != nil {
			log.Printf("ListComment|获取用户失败|%v", err)
			return response.CommentListResponse{}
		}
		var responseUser response.User
		err = copier.Copy(&responseUser, &user)
		if err != nil {
			log.Printf("ListComment|对象复制错误|%v", err)
			return response.CommentListResponse{}
		}

		responseComment[index].User = responseUser

		// 格式化时间
		responseComment[index].CreateDate = comment.CreateDate.Format("01-02")
	}
	return response.CommentListResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		CommentList: responseComment,
	}
}

// CommentPost 发表评论
func (c CommentService) CommentPost(commentPostRequest request.CommentPostRequest) response.CommentPostResponse {
	var userId int64 = 1
	var userResponse response.User
	user, err := c.UserRepository.GetUserById(userId)
	if err != nil {
		log.Printf("CommentPost|类型转换错误|%v", err)
		return response.CommentPostResponse{}
	}
	err = copier.Copy(&userResponse, &user)
	if err != nil {
		log.Printf("CommentPost|对象转化错误|%v", err)
		return response.CommentPostResponse{}
	}
	videoId, err := strconv.ParseInt(commentPostRequest.VideoId, 10, 64)
	if err != nil {
		log.Printf("CommentPost|类型转换错误|%v", err)
		return response.CommentPostResponse{}
	}
	comment := domain.Comment{UserId: userId, VideoId: videoId, Content: commentPostRequest.CommentText, CreateDate: time.Now()}
	err = c.CommentRepository.InsertComment(&comment)
	if err != nil {
		log.Printf("CommentPost|数据库插入错误|%v", err)
		return response.CommentPostResponse{}
	}

	return response.CommentPostResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		Comment: response.Comment{
			CommentId:  comment.CommentId,
			Content:    comment.Content,
			CreateDate: time.Now().Format("01-02"),
			User:       userResponse,
		},
	}
}

func NewCommentService() ICommentService {
	return CommentService{
		CommentRepository: repository.NewCommentRepository(),
		UserRepository:    repository.NewUserRepository(),
	}
}
