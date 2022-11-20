package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/golang-graphql/graph/generated"
	"github.com/golang-graphql/graph/model"
)

// CreateTransaction is the resolver for the createTransaction field.
func (r *mutationResolver) CreateTransaction(ctx context.Context, input model.NewTransaction) (*model.Transaction, error) {
	res, err := r.TransactionService.InsertTransaction(input)
	return &res, err
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	res, err := r.UserService.Create(input)
	return &res, err
}

// DeleteUserByID is the resolver for the deleteUserById field.
func (r *mutationResolver) DeleteUserByID(ctx context.Context, input *model.DeleteUserByIDRequest) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUserByID - deleteUserById"))
}

// UpdateUserByID is the resolver for the updateUserById field.
func (r *mutationResolver) UpdateUserByID(ctx context.Context, input *model.UpdateUserByIDRequest) (bool, error) {
	panic(fmt.Errorf("not implemented: UpdateUserByID - updateUserById"))
}

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context) ([]*model.Transaction, error) {
	panic(fmt.Errorf("not implemented: Transactions - transactions"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	res, err := r.UserService.GetAllUser()
	return res, err
}

// TransactionByID is the resolver for the transactionById field.
func (r *queryResolver) TransactionByID(ctx context.Context, input model.GetTransactoinByID) (*model.Transaction, error) {
	panic(fmt.Errorf("not implemented: TransactionByID - transactionById"))
}

// UserByID is the resolver for the userById field.
func (r *queryResolver) UserByID(ctx context.Context, input model.GetUserByID) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserByID - userById"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
