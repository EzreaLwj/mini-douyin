package request

// FavoriteActionRequest 点赞操作请求
type FavoriteActionRequest struct {
	Token      string `json:"token" form:"token" binding:"required"`
	VideoId    string `json:"video_id" form:"video_id" binding:"required"`
	ActionType string `json:"action_type" form:"action_type" binding:"required"`
}
