package request

type InfoRequest struct {
	UserId string `form:"user_id" json:"user_id" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}

type UserLoginRequest struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserRegisterRequest struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
