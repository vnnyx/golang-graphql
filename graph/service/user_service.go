package service

import "github.com/golang-graphql/graph/model"

type UserService interface {
	Create(request model.NewUser) (model.User, error)
	GetAllUser() ([]*model.User, error)
}
