package repository

import "github.com/golang-graphql/graph/model"

type TransactionRepository interface {
	InsertTransaction(transaction model.TransactionEntity, userId string) error
}
