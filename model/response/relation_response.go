package response

type FollowActionResponse struct {
	Response
}

type GetFollowerResponse struct {
	Response
	UserList []User `json:"user_list"`
}
