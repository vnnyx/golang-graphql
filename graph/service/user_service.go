package service

import "github.com/golang-graphql/graph/model"

type UserService interface {
	Create(request model.NewUser) (model.User, error)
	GetAllUser() ([]*model.User, error)
	FindUserById(userId string) (*model.User, error)
	UpdateUserById(request *model.UpdateUserByIDRequest) (bool, error)
	DeleteUserById(request *model.DeleteUserByIDRequest) (bool, error)
}
