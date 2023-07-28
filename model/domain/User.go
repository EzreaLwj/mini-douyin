package domain

type User struct {
	UserId          int64  `gorm:"Column:user_id" json:"user_id"`
	UserName        string `gorm:"Column:username" json:"username"`
	Password        string `gorm:"Column:password" json:"password"`
	FollowCount     int64  `gorm:"Column:follow_count" json:"follow_count"`
	FollowerCount   int64  `gorm:"Column:follower_count" json:"follower_count"`
	Avatar          string ` gorm:"Column:avatar" json:"avatar"`
	BackgroundImage string `gorm:"Column:background_image"  json:"background_image"`
	Signature       string `gorm:"Column:signature" json:"signature"`
	TotalFavorited  string `gorm:"Column:total_favorited" json:"total_favorited"`
	WorkCount       int64  `gorm:"Column:work_count" json:"work_count"`
	FavoriteCount   int64  `gorm:"Column:favorite_count" json:"favorite_count"`
}

func (User) TableName() string {
	return "tb_user"
}
