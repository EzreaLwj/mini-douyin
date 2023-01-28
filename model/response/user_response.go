package response

type User struct {
	UserId        int64  `json:"id"`
	UserName      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
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
