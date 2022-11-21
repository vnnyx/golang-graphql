package repository

import (
	"github.com/golang-graphql/graph/model"
	"github.com/golang-graphql/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &UserRepositoryImpl{
		Collection: database.Collection("users"),
	}
}

func (repository *UserRepositoryImpl) InsertUser(user model.UserEntity) error {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	user.DefaultValue()
	_, err := repository.Collection.InsertOne(ctx, user)
	return err
}

func (repository *UserRepositoryImpl) FindAllUser() (users model.ListUserEntity, err error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}

	err = cursor.All(ctx, &users)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (repository *UserRepositoryImpl) FindUserById(userId string) (user model.UserEntity, err error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	err = repository.Collection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	return user, err
}

func (repository *UserRepositoryImpl) UpdateUserById(user model.UpdateUserEntity) (bool, error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	_, err := repository.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	return true, nil
}
