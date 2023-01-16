package user

import "mini-douyin/model/domain"

type InfoResponse struct {
	StatusCode int
	StatusMsg  string
	User       domain.User
}
