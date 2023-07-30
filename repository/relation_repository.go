package repository

import (
	"log"
	"mini-douyin/config"
	"mini-douyin/model/domain"
)

type IRelationRepository interface {
	AddFollow(userId int64, toUserId int64) error        // 添加关注记录
	GetFollow(userId int64) ([]domain.Relation, error)   // 获取关注者
	GetFollower(userId int64) ([]domain.Relation, error) // 获取关注者
	CheckIsFollow(userId int64, toUserId int64) bool     // 检查是否关注
	RemoveFollow(userId int64, toUserId int64) error     //取消关注
}

type RelationRepository struct {
}

func (r RelationRepository) RemoveFollow(userId int64, toUserId int64) error {
	var relation domain.Relation
	err := config.DB.Table("tb_relation").Where("user_id = ? and to_user_id = ?", userId, toUserId).Delete(&relation).Error
	return err
}

// CheckIsFollow 检查是否被关注
func (r RelationRepository) CheckIsFollow(userId int64, toUserId int64) bool {
	var count int64
	err := config.DB.Table("tb_relation").Where("user_id = ? and to_user_id = ?", userId, toUserId).Count(&count).Error
	if err != nil {
		log.Printf("CheckIsFollow|数据库获取数量错误|%v", err)
		return false
	}
	return count == 1
}

func (r RelationRepository) GetFollower(userId int64) ([]domain.Relation, error) {
	var relations []domain.Relation
	err := config.DB.Table("tb_relation").Where("to_user_id = ?", userId).Find(&relations).Error
	return relations, err
}

func (r RelationRepository) GetFollow(userId int64) ([]domain.Relation, error) {
	var relations []domain.Relation
	err := config.DB.Table("tb_relation").Where("user_id = ?", userId).Find(&relations).Error
	return relations, err
}

func (r RelationRepository) AddFollow(userId int64, toUserId int64) error {
	relationDomain := domain.Relation{ToUserId: toUserId, UserId: userId}
	err := config.DB.Table("tb_relation").Create(&relationDomain).Error
	return err
}

func NewRelationRepository() IRelationRepository {
	relationRepository := RelationRepository{}
	return relationRepository
}
