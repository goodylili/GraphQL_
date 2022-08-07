package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"GraphQL_/graph/generated"
	"GraphQL_/graph/model"
	"context"
	"fmt"
)

type Resolver struct {
	BookRepository model.BioRepository
}

// CreateBio is the resolver for the createTodo field.
// tell them to rename these
func (r *mutationResolver) CreateBio(ctx context.Context, input model.BioData) (*model.BioData, error) {
	person, err := r.BookRepository.CreatePerson(input)
	person_ := &model.BioData{
		ID:   person.ID,
		Name: person.Name,
		Age:  person.Age,
	}
	if err != nil {
		return nil, err
	}
	return person_, err

}

// Bios is the resolver for the todos field.
func (r *queryResolver) Bios(ctx context.Context) ([]*model.BioData, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
