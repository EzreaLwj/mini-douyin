package domain

type User struct {
	UserId        int64  `gorm:"Column:user_id" json:"user_id"`
	UserName      string `gorm:"Column:username" json:"username"`
	Password      string `gorm:"Column:password" json:"password"`
	FollowCount   int64  `gorm:"Column:follow_count" json:"follow_count"`
	FollowerCount int64  `gorm:"Column:follower_count" json:"follower_count"`
}

func (User) TableName() string {
	return "tb_user"
}
