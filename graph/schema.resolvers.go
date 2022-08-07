package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"GraphQL_/graph/model"
	"context"
)

//type Resolver struct {
//	BioRepository model.BioRepository
//}

// CreateBio is the resolver for the createTodo field.
// tell them to rename these
func (r *mutationResolver) CreateBio(ctx context.Context, input model.BioData) (*model.BioData, error) {
	person, err := r.BioRepository.CreatePerson(input)
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
func (r *queryResolver) Bios(ctx context.Context) ([]model.BioData, error) {
	bio, err := r.BioRepository.GetAllPersons()
	if err != nil {
		return nil, err
	}
	return bio, nil
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id int, input model.BioData) (string, error) {
	err := r.BioRepository.UpdatePerson(input, id)
	if err != nil {
		return "nil", err
	}
	message := "successfully updated row"

	return message, nil
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (string, error) {
	err := r.BioRepository.DeletePerson(id)
	if err != nil {
		return "", err
	}
	message := "successfully deleted row"
	return message, nil
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
