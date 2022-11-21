package model

type ListUserEntity []*UserEntity

type UserEntity struct {
	ID          string               `bson:"_id"`
	Username    string               `bson:"username"`
	Email       string               `bson:"email"`
	Transaction []*TransactionEntity `bson:"transactions"`
}

type UpdateUserEntity struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
}

type TransactionEntity struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func (user *UserEntity) DefaultValue() {
	if user.Transaction == nil {
		user.Transaction = []*TransactionEntity{}
	}
}
