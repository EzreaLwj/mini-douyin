package domain

type Relation struct {
	RelationId int64 `json:"relation_id" gorm:"relation_id"`
	UserId     int64 `json:"user_id" gorm:"user_id"`
	ToUserId   int64 `json:"to_user_id" gorm:"to_user_id"`
}

func (Relation) TableName() string {
	return "tb_message"
}
