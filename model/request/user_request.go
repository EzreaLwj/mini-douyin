package request

type InfoRequest struct {
	UserId string `form:"user_id" json:"user_id" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}

type UserLoginRequest struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}
