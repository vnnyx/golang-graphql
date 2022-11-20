package service

import (
	"github.com/golang-graphql/graph/model"
	"github.com/golang-graphql/graph/repository"
	"github.com/google/uuid"
)

type TransactionServiceImpl struct {
	repository.TransactionRepository
}

func NewTransactionService(transactionRepository *repository.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: *transactionRepository,
	}
}

func (service *TransactionServiceImpl) InsertTransaction(request model.NewTransaction) (model.Transaction, error) {
	id := uuid.New().String()
	err := service.TransactionRepository.InsertTransaction(model.TransactionEntity{
		ID:   id,
		Name: request.Name},
		request.UserID)
	if err != nil {
		return model.Transaction{}, err
	}

	response := model.Transaction{
		ID:   id,
		Name: request.Name,
	}

	return response, nil
}
