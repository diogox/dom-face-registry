package registry

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/diogox/dom-face-recognizer/internal/person"
)

type PersonService interface {
	GetPeople(ctx context.Context) ([]person.Person, error)
	AddPerson(ctx context.Context, info person.Person) error
	RemovePerson(ctx context.Context, id uuid.UUID) error
}

type FaceService interface {
	RecognizeFace(ctx context.Context, imgBytes []byte) (uuid.UUID, error)
	AddFace(ctx context.Context, imgBytes []byte, personID uuid.UUID) (faceID uuid.UUID, err error)
	RemoveFace(ctx context.Context, faceID uuid.UUID) error
}

type Service struct {
	logger        logrus.FieldLogger
	personService PersonService
	faceService   FaceService
}

func NewRegistryService(logger logrus.FieldLogger, personService PersonService, faceService FaceService) *Service {
	return &Service{
		logger: logger,
		personService: personService,
		faceService:   faceService,
	}
}

func (s *Service) GetPeople(ctx context.Context) ([]person.Person, error) {
	return s.personService.GetPeople(ctx)
}

func (s *Service) AddPerson(ctx context.Context, personInfo person.Person) error {
	// TODO: Handle errors better. Check for missing personInfo fields. (use `validate` tags for that?)
	return s.personService.AddPerson(ctx, personInfo)
}

func (s *Service) RemovePerson(ctx context.Context, personID uuid.UUID) error {
	// TODO: Use 'validate' for this?

	if personID == uuid.Nil {
		return errors.New("missing person id")
	}

	return s.personService.RemovePerson(ctx, personID)
}

func (s *Service) RecognizeFace(ctx context.Context, imgBytes []byte) (person.Person, error) {
	// TODO: Use 'validate' for this?

	if len(imgBytes) == 0 {
		return person.Person{}, errors.New("missing image")
	}

	personID, err := s.faceService.RecognizeFace(ctx, imgBytes)
	if err != nil {
		return person.Person{}, err
	}

	// TODO: IMPORTANT! Create GetPersonByID method in PeopleService so we can get rid of this
	var recognized *person.Person
	people, _ := s.GetPeople(ctx)
	for _, p := range people {
		if p.ID == personID {
			recognized = &p
			break
		}
	}

	if recognized == nil {
		s.logger.Debugf("recognized person, but failed to get person by id: %v", personID)
		return person.Person{}, errors.Errorf("no match was found for person id: %v", personID)
	}

	return *recognized, nil
}

func (s *Service) AddFace(ctx context.Context, imgBytes []byte, personID uuid.UUID) (uuid.UUID, error) {
	// TODO: Use 'validate' for this?

	if len(imgBytes) == 0 {
		return uuid.Nil, errors.New("missing image")
	}

	if personID == uuid.Nil {
		return uuid.Nil, errors.New("missing person id")
	}

	// TODO: Need to create a store for the images only.
	return s.faceService.AddFace(ctx, imgBytes, personID)
}

func (s *Service) RemoveFace(ctx context.Context, faceID uuid.UUID) error {
	// TODO: Use 'validate' for this?

	if faceID == uuid.Nil {
		return errors.New("missing face id")
	}

	return s.faceService.RemoveFace(ctx, faceID)
}
