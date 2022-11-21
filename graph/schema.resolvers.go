package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/golang-graphql/graph/generated"
	"github.com/golang-graphql/graph/model"
)

// CreateTransaction is the resolver for the createTransaction field.
func (r *mutationResolver) CreateTransaction(ctx context.Context, input model.NewTransaction) (*model.Transaction, error) {
	return r.TransactionService.InsertTransaction(input)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.UserService.Create(input)
}

// DeleteUserByID is the resolver for the deleteUserById field.
func (r *mutationResolver) DeleteUserByID(ctx context.Context, input *model.DeleteUserByIDRequest) (bool, error) {
	return r.UserService.DeleteUserById(input)
}

// UpdateUserByID is the resolver for the updateUserById field.
func (r *mutationResolver) UpdateUserByID(ctx context.Context, input *model.UpdateUserByIDRequest) (bool, error) {
	return r.UserService.UpdateUserById(input)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.UserService.GetAllUser()
}

// TransactoinByUserID is the resolver for the transactoinByUserId field.
func (r *queryResolver) TransactoinByUserID(ctx context.Context, input model.GetTransactoinByUserID) ([]*model.Transaction, error) {
	return r.TransactionService.GetTransactionByUserId(input.UserID)
}

// UserByID is the resolver for the userById field.
func (r *queryResolver) UserByID(ctx context.Context, input model.GetUserByID) (*model.User, error) {
	res, err := r.UserService.FindUserById(input.ID)
	return res, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
