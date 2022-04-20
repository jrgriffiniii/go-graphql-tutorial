package main

import (
	"testing"

	"github.com/jrgriffiniii/go-graphql-tutorial/graph/model"
)

func TestUser(t *testing.T) {

	user := model.User{
		ID:   "user-1",
		Name: "test user",
	}

	if user.ID != "user-1" {
		t.Errorf("User.ID should be 'user-1'; received %s", user.ID)
	}

	if user.Name != "test user" {
		t.Errorf("User.Name should be 'test user'; received %s", user.Name)
	}
}
