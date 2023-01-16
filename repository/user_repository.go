package repository

import (
	"fmt"
	"mini-douyin/common"
	"mini-douyin/model/domain"
)

type IUserRepository interface {
	GetUserById(id int64) (domain.User, error) // 获取单个用户
}

// UserRepository 定义一个结构体
type UserRepository struct {
}

// NewUserRepository UserRepository构造函数
func NewUserRepository() IUserRepository {
	return UserRepository{}
}

// GetUserById 获取单个用户
func (ur UserRepository) GetUserById(id int64) (domain.User, error) {
	fmt.Println("GetUserById---")
	var user domain.User
	err := common.DB.Where("user_id = ?", id).First(&user).Error
	return user, err
}
