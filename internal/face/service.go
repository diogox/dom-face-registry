package face

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/diogox/dom-face-recognizer/internal/face/recognizer"
)

const (
	errFaceNotRecognized = "could not recognize face"
)

type Store interface {
	GetFaces(ctx context.Context) ([]Face, error)
	AddFace(ctx context.Context, encoding recognizer.FaceEncoding, imgBytes []byte, personID uuid.UUID) (faceID uuid.UUID, err error)
	RemoveFace(ctx context.Context, faceID uuid.UUID) error
}

type Service struct {
	store      Store
	recognizer recognizer.Recognizer
}

func NewService(store Store, faceRecognizer recognizer.Recognizer) *Service {
	return &Service{
		store:      store,
		recognizer: faceRecognizer,
	}
}

func (s *Service) RecognizeFace(ctx context.Context, imgBytes []byte) (uuid.UUID, error) {
	faces, err := s.store.GetFaces(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	var peopleIDs []uuid.UUID
	var allFaceEncodings []recognizer.FaceEncoding
	for _, f := range faces {
		peopleIDs = append(peopleIDs, f.PersonID)
		allFaceEncodings = append(allFaceEncodings, f.Encoding)
	}

	targetFaceEncoding, err := s.recognizer.EncodeFace(imgBytes)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to encode target face")
	}

	personID, err := s.recognizer.Recognize(targetFaceEncoding, peopleIDs, allFaceEncodings)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, errFaceNotRecognized)
	}

	return personID, nil
}

func (s *Service) AddFace(ctx context.Context, imgBytes []byte, personID uuid.UUID) (faceID uuid.UUID, err error) {
	faceEncoding, err := s.recognizer.EncodeFace(imgBytes)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to encode face to be added")
	}

	return s.store.AddFace(ctx, faceEncoding, imgBytes, personID)
}

func (s *Service) RemoveFace(ctx context.Context, faceID uuid.UUID) error {
	return s.store.RemoveFace(ctx, faceID)
}
