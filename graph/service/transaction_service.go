package service

import "github.com/golang-graphql/graph/model"

type TransactionService interface {
	InsertTransaction(request model.NewTransaction) (model.Transaction, error)
	GetTransactionByUserId(userId string) ([]*model.Transaction, error)
}
