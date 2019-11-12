package person

import (
	"context"

	"github.com/google/uuid"
)

type Store interface {
	GetPeople(ctx context.Context) ([]Person, error)
	FindPersonByID(ctx context.Context, id uuid.UUID) (Person, error)
	CreatePerson(ctx context.Context, person Person) error
	DeletePerson(ctx context.Context, id uuid.UUID) error
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}

func (p *Service) GetPeople(ctx context.Context) ([]Person, error) {
	return p.store.GetPeople(ctx)
}

func (p *Service) FindPersonByID(ctx context.Context, id uuid.UUID) (Person, error) {
	return p.store.FindPersonByID(ctx, id)
}

func (p *Service) AddPerson(ctx context.Context, person Person) error {
	return p.store.CreatePerson(ctx, person)
}

func (p *Service) RemovePerson(ctx context.Context, id uuid.UUID) error {
	return p.store.DeletePerson(ctx, id)
}
