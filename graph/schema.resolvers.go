package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jrgriffiniii/hackernews/graph/generated"
	"github.com/jrgriffiniii/hackernews/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var user model.User
	var todo model.Todo

	// This would be replaced by an ORM query to load the model instance
	user.ID = input.UserID

	todo.User = &user
	todo.Text = input.Text
	// One would also assign the model instance ID here
	todo.ID = "todo-2"

	return &todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo

	// This would be replaced by an ORM query to load the model instances
	fixtureTodo := model.Todo{
		ID:   "todo-1",
		Text: "test todo item",
		Done: false,
		User: &model.User{
			ID:   "user-1",
			Name: "test user",
		},
	}
	todos = append(todos, &fixtureTodo)

	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
