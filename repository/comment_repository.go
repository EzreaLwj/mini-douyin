package repository

import (
	"mini-douyin/config"
	"mini-douyin/model/domain"
)

type ICommentRepository interface {
	InsertComment(d *domain.Comment) error
	GetCommentList(videoId int64) ([]domain.Comment, error) // 添加评论

}

type CommentRepository struct {
}

// GetCommentList 获取评论列表
func (c CommentRepository) GetCommentList(videoId int64) ([]domain.Comment, error) {
	var commentList []domain.Comment
	err := config.DB.Table("tb_comment").Where("video_id = ?", videoId).Find(&commentList).Error
	return commentList, err
}

// InsertComment 添加评论
func (c CommentRepository) InsertComment(d *domain.Comment) error {
	return config.DB.Table("tb_comment").Create(d).Error
}

func NewCommentRepository() ICommentRepository {
	return CommentRepository{}
}
