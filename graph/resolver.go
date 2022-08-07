package graph

import (
	"GraphQL_/graph/generated"
	"GraphQL_/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	BioRepository model.BioRepository
}

func (r Resolver) Mutation() generated.MutationResolver {
	//TODO implement me
	panic("implement me")
}

func (r Resolver) Query() generated.QueryResolver {
	//TODO implement me
	panic("implement me")
}
