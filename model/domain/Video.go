package domain

import "time"

type Video struct {
	VideoId       int64     `gorm:"Column:video_id" json:"video_id"`
	UserId        int64     `gorm:"Column:user_id" json:"user_id"`
	Title         string    `gorm:"Column:title" json:"title"`
	PlayUrl       string    `gorm:"Column:play_url" json:"play_url"`
	CoverUrl      string    `gorm:"Column:cover_url" json:"cover_url"`
	FavoriteCount int64     `gorm:"Column:favorite_count" json:"favorite_count"`
	CommentCount  int64     `gorm:"Column:comment_count" json:"comment_count"`
	CreateTime    time.Time `gorm:"Column:create_time" json:"create_time"`
	UpdateTime    time.Time `gorm:"Column:update_time" json:"update_time"`
}

func (Video) TableName() string {
	return "tb_video"
}
