package repository

import (
	"mini-douyin/config"
	"mini-douyin/model/domain"
)

type IFavoriteRepository interface {
	AddFavoriteItem(like domain.Like) error               // 增加点赞记录
	DeleteFavoriteItem(userId int64, videoId int64) error // 取消点赞记录
	GetUserLikeVideoId(useId int64) ([]domain.Like, error)
	CheckIsLike(userid int64, id2 int64) (bool, error)
}

// FavoriteRepository UserRepository 定义一个结构体
type FavoriteRepository struct {
}

// CheckIsLike 检查视频是否被点赞
func (f FavoriteRepository) CheckIsLike(userId int64, videoId int64) (bool, error) {
	var count int64 = 0
	err := config.DB.Table("tb_like").Where("user_id = ? and video_id = ?", userId, videoId).Count(&count).Error
	if count == 0 {
		return false, err
	}
	return true, err
}

func (f FavoriteRepository) GetUserLikeVideoId(useId int64) ([]domain.Like, error) {
	//var videoIds []int64
	var likes []domain.Like
	//err := config.DB.Table("tb_like").Select([]string{"video_id"}).Where("user_id = ?", useId).Scan(&videoIds).Error
	err := config.DB.Table("tb_like").Select([]string{"video_id"}).Where("user_id = ?", useId).Scan(&likes).Error
	return likes, err
}

// AddFavoriteItem 增加点赞记录
func (f FavoriteRepository) AddFavoriteItem(like domain.Like) error {
	err := config.DB.Table("tb_like").Create(&like).Error
	return err
}

// DeleteFavoriteItem 删除点赞记录
func (f FavoriteRepository) DeleteFavoriteItem(usrId int64, videoId int64) error {
	err := config.DB.Table("tb_like").Where("user_id = ? and video_id =  ?", usrId, videoId).Delete(&domain.Like{}).Error
	return err
}

// NewFavoriteRepository UserRepository构造函数
func NewFavoriteRepository() IFavoriteRepository {
	return FavoriteRepository{}
}
