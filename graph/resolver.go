package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// type Resolver struct{}

import "github.com/rajshivaram/graphql-todos/model"

type Resolver struct{
	todos []*model.Todo
}
