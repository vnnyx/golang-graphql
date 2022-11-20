package model

type ListUserEntity []*UserEntity

type UserEntity struct {
	ID          string               `bson:"_id"`
	Username    string               `bson:"username"`
	Password    string               `bson:"password"`
	Email       string               `bson:"email"`
	Transaction []*TransactionEntity `bson:"transactions"`
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
