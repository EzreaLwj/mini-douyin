package response

type User struct {
	UserId          int64  `json:"id"`
	UserName        string `json:"name"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	FollowCount     int64  `json:"follow_count"`
	TotalFavorited  string `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
}
type InfoResponse struct {
	Response
	User User `json:"user"`
}

type LoginResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type RegisterResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
