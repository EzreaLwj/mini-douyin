package repository

import (
	"github.com/jinzhu/gorm"
	"log"
	"mini-douyin/config"
	"mini-douyin/model/domain"
)

type IVideoRepository interface {
	GetVideoListByUserId(userId int64) ([]domain.Video, error)       // 获取视频列表
	GetVideoListByVideoIds(videoIds []int64) ([]domain.Video, error) // 获取视频列表
	FeedVideoList() ([]domain.Video, error)                          // 获取视频流列表
	SaveVideo(video *domain.Video) error                             // 保存视频
	IncreaseFavoriteCount(videoId int64) error                       // 增加点赞数量
	DecreaseFavoriteCount(videoId int64) error                       // 减少点赞数量
}

// VideoRepository 定义一个结构体
type VideoRepository struct {
}

// GetVideoListByVideoIds 根据id查找视频
func (v VideoRepository) GetVideoListByVideoIds(videoIds []int64) ([]domain.Video, error) {
	videos := make([]domain.Video, 10)
	err := config.DB.Table("tb_video").Where("video_id in (?)", videoIds).Find(&videos).Error
	return videos, err
}

// NewVideoRepository VideoRepository构造函数
func NewVideoRepository() IVideoRepository {
	return VideoRepository{}
}

// SaveVideo 保存视频
func (v VideoRepository) SaveVideo(video *domain.Video) error {
	log.Printf("SaveVideo|保存用户视频|%v", *video)
	err := config.DB.Create(video).Error
	return err
}

// GetVideoList 获取某个用户的视频
func (v VideoRepository) GetVideoListByUserId(userId int64) ([]domain.Video, error) {
	log.Printf("GetVideoListByUserId|获取用户的视频|%v", userId)
	var video []domain.Video
	err := config.DB.Where("user_id = ?", userId).Find(&video).Error

	return video, err
}

// FeedVideoList 获取视频流
func (v VideoRepository) FeedVideoList() ([]domain.Video, error) {
	log.Printf("FeedVideoList|获取视频流|")
	var video []domain.Video
	err := config.DB.Limit(30).Order("create_time desc").Find(&video).Error
	return video, err
}

// IncreaseFavoriteCount 增加点赞数量
func (v VideoRepository) IncreaseFavoriteCount(videoId int64) error {
	log.Printf("IncreaseFavoriteCount|增加点赞数量|")
	err := config.DB.Table("tb_video").Where("video_id = ?", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	return err
}

// DecreaseFavoriteCount 减少点赞数量
func (v VideoRepository) DecreaseFavoriteCount(videoId int64) error {
	log.Printf("IncreaseFavoriteCount|减少点赞数量|")
	err := config.DB.Table("tb_video").Where("video_id = ?", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	return err
}
