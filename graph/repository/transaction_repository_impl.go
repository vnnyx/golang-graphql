package repository

import (
	"github.com/golang-graphql/graph/model"
	"github.com/golang-graphql/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewTransactionRepository(database *mongo.Database) TransactionRepository {
	return &TransactionRepositoryImpl{
		Collection: database.Collection("users"),
	}
}

func (repository *TransactionRepositoryImpl) InsertTransaction(transaction model.TransactionEntity, userId string) error {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	filter := bson.M{"_id": userId}
	update := bson.M{"$push": bson.M{"transactions": transaction}}
	_, err := repository.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (repository *TransactionRepositoryImpl) GetTransactionByUserId(userId string) ([]*model.TransactionEntity, error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	user := model.UserEntity{}

	filter := bson.M{"_id": userId}
	err := repository.Collection.FindOne(ctx, filter).Decode(&user)

	return user.Transaction, err
}
