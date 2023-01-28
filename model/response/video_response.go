package response

type Video struct {
	VideoId       int64  `json:"id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
	User          User   `json:"author"`
}

type VideoList struct {
	Response
	Video []Video `json:"video_list"`
}

type VideoFeedList struct {
	Response
	NextTime int64   `json:"next_time"`
	Video    []Video `json:"video_list"`
}

type VideoPostResponse struct {
	Response
}
