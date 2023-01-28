package domain

type Relation struct {
	RelationId int64 `json:"relation_id"`
	UserId     int64 `json:"user_id"`
	ToUserId   int64 `json:"to_user_id"`
}
