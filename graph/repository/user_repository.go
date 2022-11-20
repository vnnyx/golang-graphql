package repository

import "github.com/golang-graphql/graph/model"

type UserRepository interface {
	InsertUser(user model.UserEntity) error
	FindAllUser() (users model.ListUserEntity, err error)
}
