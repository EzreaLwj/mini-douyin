package domain

import "time"

type Comment struct {
	CommentId  int64     `gorm:"Column:comment_id;AUTO_INCREMENT" json:"comment_id"`
	UserId     int64     `gorm:"Column:user_id" json:"user_id"`
	VideoId    int64     `gorm:"Column:video_id" json:"video_id"`
	Content    string    `gorm:"Column:content" json:"content"`
	CreateDate time.Time `gorm:"Column:create_date" json:"create_date"`
}
