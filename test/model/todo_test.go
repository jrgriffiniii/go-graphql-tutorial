package main

import (
	"testing"

	"github.com/jrgriffiniii/go-graphql-tutorial/graph/model"
)

func TestTodo(t *testing.T) {

	todo := model.Todo{
		ID:   "todo-1",
		Text: "test todo item",
		Done: false,
		User: &model.User{
			ID:   "user-1",
			Name: "test user",
		},
	}

	if todo.ID != "todo-1" {
		t.Errorf("Todo.ID should be 'todo-1'; received %s", todo.ID)
	}

	if todo.Text != "test todo item" {
		t.Errorf("Todo.Text should be 'test todo item'; received %s", todo.Text)
	}

	if todo.Done != false {
		t.Error("Todo.Done should be 'false'; received 'true'")
	}
}
