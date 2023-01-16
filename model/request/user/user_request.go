package user

type InfoRequest struct {
	UserId string `form:"user_id" json:"user_id" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}
