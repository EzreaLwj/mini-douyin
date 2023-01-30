package repository

import (
	"log"
	"mini-douyin/config"
	"mini-douyin/model/domain"
)

type IUserRepository interface {
	GetUserById(id int64) (domain.User, error)                        // 获取单个用户
	CreateUser(userName string, password string) (domain.User, error) // 创建新用户
	CheckUser(userName string, password string) (domain.User, error)  // 检查是否存在用户
}

// UserRepository 定义一个结构体
type UserRepository struct {
}

// NewUserRepository UserRepository构造函数
func NewUserRepository() IUserRepository {
	return UserRepository{}
}

// CheckUser 检查用户是否存在
func (ur UserRepository) CheckUser(userName string, password string) (domain.User, error) {
	var user domain.User
	err := config.DB.Table("tb_user").Where("username = ? and password = ?", userName, password).Scan(&user).Error
	return user, err
}

// GetUserById 获取单个用户
func (ur UserRepository) GetUserById(id int64) (domain.User, error) {
	log.Printf("GetUserById|获取用户信息|%v", id)
	var user domain.User
	err := config.DB.Where("user_id = ?", id).First(&user).Error
	return user, err
}

// CreateUser 创建新用户
func (ur UserRepository) CreateUser(userName string, password string) (domain.User, error) {
	var user domain.User
	user.UserName = userName
	user.Password = password
	user.FollowCount = 0
	user.FollowerCount = 0
	err := config.DB.Table("tb_user").Scan(&user).Error
	return user, err
}
