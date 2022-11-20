package service

import (
	"github.com/golang-graphql/graph/model"
	"github.com/golang-graphql/graph/repository"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: *userRepository,
	}
}

func (service *UserServiceImpl) Create(request model.NewUser) (model.User, error) {
	id := uuid.New().String()
	err := service.UserRepository.InsertUser(model.UserEntity{
		ID:       id,
		Username: request.Username,
		Email:    request.Email,
	})
	if err != nil {
		return model.User{}, err
	}

	response := model.User{
		ID:       id,
		Username: request.Username,
		Email:    request.Email,
	}

	return response, nil
}

func (service *UserServiceImpl) GetAllUser() ([]*model.User, error) {
	users, err := service.UserRepository.FindAllUser()
	if err != nil {
		return []*model.User{}, err
	}

	listUser := []*model.User{}
	transaction := []*model.Transaction{}
	for _, user := range users {
		for _, t := range user.Transaction {
			transaction = append(transaction, &model.Transaction{
				ID:   t.ID,
				Name: t.Name,
			})
		}
		listUser = append(listUser, &model.User{
			ID:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			Transactions: transaction,
		})
	}
	return listUser, nil
}
