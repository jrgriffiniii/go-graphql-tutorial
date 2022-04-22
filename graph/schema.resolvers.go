package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/jrgriffiniii/go-graphql-tutorial/graph/generated"
	"github.com/jrgriffiniii/go-graphql-tutorial/graph/model"
	"github.com/jrgriffiniii/go-graphql-tutorial/internal/todos"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var todo todos.Todo

	todo.Text = input.Text
	todo.Done = false
	todoID := todo.Save()

	return &model.Todo{ID: strconv.FormatInt(todoID, 10), Text: todo.Text, Done: todo.Done}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {

	var resultTodos []*model.Todo
	var dbTodos []todos.Todo
	dbTodos = todos.GetAll()
	for _, todo := range dbTodos {
		resultTodos = append(resultTodos, &model.Todo{ID: todo.ID, Text: todo.Text, Done: todo.Done})
	}
	return resultTodos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
