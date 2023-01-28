package request

import (
	"mime/multipart"
)

// PostRequest 提交请求参数
type PostRequest struct {
	Token string                `form:"token"`
	Title string                `form:"title"`
	Data  *multipart.FileHeader `form:"data"`
}

// ListRequest 视频列表请求接口
type ListRequest struct {
	Token string `form:"token"`
	UseId int64  `form:"user_id"`
}

type FeedRequest struct {
	Token      string `form:"token"`
	LatestTime string `form:"latest_time"`
}
