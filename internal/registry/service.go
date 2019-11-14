//go:generate mockgen -package registry -source=service.go -destination service_mocks.go

package registry

import (
	"context"
	"github.com/diogox/dom-face-registry/internal/face"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/diogox/dom-face-registry/internal/person"
)

type PersonService interface {
	GetPeople(ctx context.Context) ([]person.Person, error)
	FindPersonByID(ctx context.Context, id uuid.UUID) (person.Person, error)
	AddPerson(ctx context.Context, info person.Person) error
	RemovePerson(ctx context.Context, id uuid.UUID) error
}

type FaceService interface {
	RecognizeFace(ctx context.Context, imgBytes []byte) (uuid.UUID, error)
	FindFacesByPersonID(ctx context.Context, personID uuid.UUID) ([]face.Face, error)
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
		logger:        logger,
		personService: personService,
		faceService:   faceService,
	}
}

func (s *Service) GetPeople(ctx context.Context) ([]person.Person, error) {
	return s.personService.GetPeople(ctx)
}

func (s *Service) AddPerson(ctx context.Context, personInfo person.Person) error {
	// TODO: Handle errors better. Check for missing personInfo fields. (use `validate` tags for that?)
	if personInfo.ID == uuid.Nil {
		return errors.New(errMissingPersonInfo)
	}

	return s.personService.AddPerson(ctx, personInfo)
}

func (s *Service) RemovePerson(ctx context.Context, personID uuid.UUID) error {
	// TODO: Use 'validate' for this?

	if personID == uuid.Nil {
		return errors.New("missing person id")
	}

	faces, err := s.faceService.FindFacesByPersonID(ctx, personID)
	if err != nil {
		return errors.Wrapf(err, "failed to get faces with person id: %v", personID)
	}

	for _, f := range faces {
		err := s.faceService.RemoveFace(ctx, f.ID)
		if err != nil {
			return errors.Wrapf(err, "failed to remove face associated to person. face id: %v", f.ID)
		}
	}

	return s.personService.RemovePerson(ctx, personID)
}

func (s *Service) RecognizeFace(ctx context.Context, imgBytes []byte) (person.Person, error) {
	// TODO: Use 'validate' for this?

	if imgBytes == nil || len(imgBytes) == 0 {
		return person.Person{}, errors.New(errMissingImage)
	}

	personID, err := s.faceService.RecognizeFace(ctx, imgBytes)
	if err != nil {
		return person.Person{}, err
	}

	recognized, err := s.personService.FindPersonByID(ctx, personID)
	if err != nil {
		s.logger.Debugf("recognized person, but failed to get person by id: %v", personID)
		return person.Person{}, errors.Wrapf(err, "no match was found for person id: %v", personID)
	}

	return recognized, nil
}

func (s *Service) AddFace(ctx context.Context, imgBytes []byte, personID uuid.UUID) (uuid.UUID, error) {
	// TODO: Use 'validate' for this?

	if imgBytes == nil || len(imgBytes) == 0 {
		return uuid.Nil, errors.New(errMissingImage)
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
		return errors.New(errMissingFaceID)
	}

	return s.faceService.RemoveFace(ctx, faceID)
}
