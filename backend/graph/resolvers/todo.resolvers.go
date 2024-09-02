package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/mashumarrow/todoapplication/backend/graph"
	"github.com/mashumarrow/todoapplication/backend/models"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, userid string) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented: Todo - todo"))
}

// Userid is the resolver for the userid field.
func (r *todoResolver) Userid(ctx context.Context, obj *models.Todo) (string, error) {
	panic(fmt.Errorf("not implemented: Userid - userid"))
}

// Todo returns graph.TodoResolver implementation.
func (r *Resolver) Todo() graph.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
