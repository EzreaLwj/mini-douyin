package request

// ChatPostRequest 发表聊天请求
type ChatPostRequest struct {
	Token      string `json:"token" form:"token" binding:"required"`
	ToUserId   string `json:"to_user_id" form:"to_user_id" binding:"required"`
	ActionType string `json:"action_type" form:"action_type" binding:"required"`
	Content    string `json:"content" form:"content"`
}

type ChatListMessageRequest struct {
	ToUserID   string `json:"to_user_id" form:"to_user_id" binding:"required"` // 对方用户id
	Token      string `json:"token" form:"token" binding:"required"`           // 用户鉴权token
	PreMsgTime int64  `json:"pre_msg_time" form:"pre_msg_time"`                // 上次最新消息的时间
}
