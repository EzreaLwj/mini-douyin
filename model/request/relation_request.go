package request

type FollowActionRequest struct {
	Token      string `json:"token" form:"token"  binding:"required"`
	ToUserId   string `json:"to_user_id" form:"to_user_id"  binding:"required"`
	ActionType string `json:"action_type" form:"action_type"  binding:"required"`
}

type GetFollowerRequest struct {
	Token  string `json:"token" form:"token"  binding:"required"`
	UserId string `json:"user_id" form:"user_id"  binding:"required"`
}
