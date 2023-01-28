package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"log"
	"mini-douyin/common"
	"mini-douyin/model/domain"
	"mini-douyin/model/request"
	"mini-douyin/model/response"
	"mini-douyin/repository"
	"path/filepath"
	"time"
)

type IVideoService interface {
	PostVideo(r *request.PostRequest, c *gin.Context) response.VideoPostResponse // 发布
	ListVideo(r *request.ListRequest) response.VideoList                         // 获取某个用户的视频
	FeedVideo(r *request.FeedRequest) response.VideoFeedList                     // 获取视频流
}

type VideoService struct {
	VideoRepository repository.IVideoRepository
	UserRepository  repository.IUserRepository
}

// NewVideoService 构造函数
func NewVideoService() IVideoService {
	videoRepository := repository.NewVideoRepository()
	userRepository := repository.NewUserRepository()
	videoService := VideoService{
		VideoRepository: videoRepository,
		UserRepository:  userRepository,
	}
	return videoService
}

// PostVideo 上传视频
func (v VideoService) PostVideo(r *request.PostRequest, c *gin.Context) response.VideoPostResponse {
	log.Printf("PostVideo|上传视频|%v", *r)
	filename := filepath.Base(r.Data.Filename)
	saveFile := filepath.Join("./public/", filename)

	// 保存到本地
	if err := c.SaveUploadedFile(r.Data, saveFile); err != nil {
		log.Printf("PostVideo|存储错误|%v", err)

		return response.VideoPostResponse{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		}
	}

	// todo 可以用消息队列优化
	// 保存到COS
	go func() {
		playUrl := common.SaveFile(r.Data, c)
		// userId 和 coverUrl先写死
		videoDomain := domain.Video{
			UserId:     1,
			Title:      r.Title,
			PlayUrl:    playUrl,
			CoverUrl:   "https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/3874/logo.png",
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		err := v.VideoRepository.SaveVideo(&videoDomain)
		if err != nil {
			log.Printf("PostVideo|数据库保存错误|%v", err)
			return
		}
		log.Printf("PostVideo|视频上传成功|%v", *r)
	}()

	return response.VideoPostResponse{
		Response: response.Response{
			StatusCode: 0,
			StatusMsg:  "post success",
		},
	}
}

// ListVideo 用户的发布列表
func (v VideoService) ListVideo(r *request.ListRequest) response.VideoList {
	// 查询用户
	useId := r.UseId
	user, err := v.UserRepository.GetUserById(useId)
	if err != nil {
		log.Printf("ListVideo|获取实体失败|%v", err)
		return response.VideoList{}
	}

	var responseUser response.User

	err = copier.Copy(&responseUser, &user)
	if err != nil {
		log.Printf("ListVideo|格式转换失败|%v", err)
		return response.VideoList{}
	}

	list, err := v.VideoRepository.GetVideoListByUserId(useId)
	if err != nil {
		log.Printf("ListVideo|数据库获取错误|%v", err)
		return response.VideoList{}
	}

	var responseVideoList []response.Video
	err = copier.Copy(&responseVideoList, &list)
	if err != nil {
		log.Printf("ListVideo|格式转换失败|%v", err)
		return response.VideoList{}
	}

	for index := range responseVideoList {
		responseVideoList[index].User = responseUser
	}

	var videoList response.VideoList

	videoList.Video = responseVideoList
	videoList.Response = response.Response{StatusMsg: "成功返回信息", StatusCode: 0}

	return videoList
}

// FeedVideo 获取视频流接口
func (v VideoService) FeedVideo(r *request.FeedRequest) response.VideoFeedList {

	list, err := v.VideoRepository.FeedVideoList()
	if err != nil {
		log.Printf("ListVideo|数据库获取错误|%v", err)
		return response.VideoFeedList{}
	}

	var videoList response.VideoFeedList

	// 遍历 添加用户信息
	for index := range list {
		userId := list[index].UserId
		user, err := v.UserRepository.GetUserById(userId)
		if err != nil {
			log.Printf("ListVideo|数据库获取错误|%v", err)
			return response.VideoFeedList{}
		}

		var responseUser response.User
		err = copier.Copy(&responseUser, &user)
		if err != nil {
			log.Printf("ListVideo|格式转换失败|%v", err)
			return response.VideoFeedList{}
		}
		var responseVideo response.Video
		err = copier.Copy(&responseVideo, &list[index])
		if err != nil {
			log.Printf("ListVideo|格式转换失败|%v", err)
			return response.VideoFeedList{}
		}
		responseVideo.User = responseUser
		videoList.Video = append(videoList.Video, responseVideo)
	}

	videoList.Response = response.Response{
		StatusMsg:  "Feed",
		StatusCode: 0,
	}
	videoList.NextTime = time.Now().Unix()

	return videoList
}
