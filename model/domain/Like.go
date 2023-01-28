package domain

type Like struct {
	LikeId  int64 `gorm:"Column:like_id;AUTO_INCREMENT" json:"like_id"`
	UserId  int64 `gorm:"Column:user_id" json:"user_id"`
	VideoId int64 `gorm:"Column:video_id" json:"video_id"`
}

func (Like) TableName() string {
	return "tb_like"
}
