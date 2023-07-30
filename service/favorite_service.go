package service

import (
	"github.com/jinzhu/copier"
	"log"
	"mini-douyin/config"
	"mini-douyin/model/domain"
	"mini-douyin/model/request"
	"mini-douyin/model/response"
	"mini-douyin/repository"
	"strconv"
)

type IFavoriteService interface {
	FavoriteAction(userId int64, favoriteRequest request.FavoriteActionRequest)
	FavoriteVideoList(r *request.ListRequest) response.VideoList // 点赞操作
}

type FavoriteService struct {
	VideoRepository    repository.IVideoRepository
	FavoriteRepository repository.IFavoriteRepository
	UserRepository     repository.IUserRepository
}

func (f FavoriteService) FavoriteVideoList(r *request.ListRequest) response.VideoList {
	userId := r.UseId
	user, err := f.UserRepository.GetUserById(userId)
	if err != nil {
		log.Printf("获取用户信息失败")
		return response.VideoList{}
	}
	var userResponse response.User
	err = copier.Copy(&userResponse, &user)
	userResponse.TotalFavorited = strconv.FormatInt(user.TotalFavorited, 10)
	if err != nil {
		return response.VideoList{}
	}

	likes, err := f.FavoriteRepository.GetUserLikeVideoId(userId)
	if err != nil {
		return response.VideoList{}
	}

	videoIds := make([]int64, len(likes))
	for idx, like := range likes {
		videoIds[idx] = like.VideoId
	}

	videos, err := f.VideoRepository.GetVideoListByVideoIds(videoIds)
	if err != nil {
		return response.VideoList{}
	}

	videosResponse := make([]response.Video, len(videos))

	err = copier.Copy(&videosResponse, &videos)
	if err != nil {
		log.Printf("拷贝失败")
		return response.VideoList{}
	}
	for i := range videosResponse {
		videosResponse[i].User = userResponse
	}

	return response.VideoList{
		Video:    videosResponse,
		Response: response.Response{StatusMsg: "success", StatusCode: 0},
	}
}

// FavoriteAction 点赞操作
func (f FavoriteService) FavoriteAction(userId int64, favoriteRequest request.FavoriteActionRequest) {
	actionType := favoriteRequest.ActionType
	videoId, err := strconv.ParseInt(favoriteRequest.VideoId, 10, 64)
	if err != nil {
		log.Printf("FavoriteAction|格式转换错误|%v", err)
	}
	begin := config.DB.Begin()

	// 查询作品作者 id
	workerId := f.VideoRepository.GetUserIdByVideoId(videoId)
	if actionType == "1" {

		// 增加视频的点赞数量
		err := f.VideoRepository.IncreaseFavoriteCount(videoId)
		if err != nil {
			log.Printf("FavoriteAction|点赞失败|%v", err)
			begin.Rollback()
			return
		}

		// 添加点赞记录
		err = f.FavoriteRepository.AddFavoriteItem(domain.Like{
			VideoId: videoId,
			UserId:  userId,
		})
		if err != nil {
			log.Printf("FavoriteAction|点赞失败|%v", err)
			begin.Rollback()
			return
		}

		// 更新点赞数量
		f.UserRepository.UpdateFavoriteCount(userId, true)

		// 增加作者的获赞总数
		f.UserRepository.UpdateTotalFavoritedCount(workerId, true)

	} else if actionType == "2" {

		// 减少点赞数量
		err := f.VideoRepository.DecreaseFavoriteCount(videoId)
		if err != nil {
			log.Printf("FavoriteAction|取消点赞失败|%v", err)
			begin.Rollback()
			return
		}

		// 删除点赞记录
		err = f.FavoriteRepository.DeleteFavoriteItem(userId, videoId)
		if err != nil {
			log.Printf("FavoriteAction|取消点赞失败|%v", err)
			begin.Rollback()
			return
		}

		// 减少点赞数量
		f.UserRepository.UpdateFavoriteCount(userId, false)

		// 减少作者的获赞总数
		//f.UserRepository.UpdateTotalFavoritedCount(workerId, false)
	} else {
		log.Printf("FavoriteAction|参数错误|actionType=%v", actionType)
	}
	begin.Commit()
}

func NewFavoriteService() IFavoriteService {
	favoriteService := FavoriteService{
		VideoRepository:    repository.NewVideoRepository(),
		FavoriteRepository: repository.NewFavoriteRepository(),
		UserRepository:     repository.NewUserRepository(),
	}
	return favoriteService
}
