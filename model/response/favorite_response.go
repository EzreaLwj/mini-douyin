package response

type FavoriteActionResponse struct {
	Response
}

type FavoriteListResponse struct {
	Response
	Video []Video `json:"video_list"`
}
