package request

// CommentPostRequest 发表评论请求
type CommentPostRequest struct {
	Token       string `json:"token" form:"token" binding:"required"`
	VideoId     string `json:"video_id" form:"video_id" binding:"required"`
	ActionType  string `json:"action_type" form:"action_type" binding:"required"`
	CommentText string `json:"comment_text" form:"comment_text"`
	CommentId   string `json:"comment_id" form:"comment_id" `
}

// CommentListRequest 获取评论列表
type CommentListRequest struct {
	Token   string `json:"token" form:"token"`
	VideoId string `json:"video_id" form:"video_id"`
}
