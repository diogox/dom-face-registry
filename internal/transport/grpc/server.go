package grpc

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	pb "github.com/diogox/dom-face-recognizer/internal/pb"
	"github.com/diogox/dom-face-recognizer/internal/person"
)

type RegistryService interface {
	GetPeople(ctx context.Context) ([]person.Person, error)
	AddPerson(ctx context.Context, personInfo person.Person) error
	RemovePerson(ctx context.Context, personID uuid.UUID) error

	RecognizeFace(ctx context.Context, imgData []byte) (person.Person, error)
	AddFace(ctx context.Context, imgData []byte, personID uuid.UUID) (uuid.UUID, error)
	RemoveFace(ctx context.Context, faceID uuid.UUID) error
}

type PeopleConverter interface {
	PersonAsResponse(resPerson person.Person) *pb.Person
	PersonAsRequest(reqPerson pb.Person) person.Person
	PeopleAsResponse(reqPeople []person.Person) []*pb.Person
	PeopleAsRequest(resPeople []pb.Person) []person.Person
}

type Server struct {
	logger          logrus.FieldLogger
	service         RegistryService
	peopleConverter PeopleConverter
}

func NewServer(logger logrus.FieldLogger, registryService RegistryService, peopleConverter PeopleConverter) *Server {
	return &Server{
		logger:          logger,
		service:         registryService,
		peopleConverter: peopleConverter,
	}
}

func (s *Server) RecognizeFace(stream pb.FaceRegistry_RecognizeFaceServer) error {
	s.logger.Debug("received 'RecognizeFace' request")
	ctx := stream.Context()

	// Receive image upload
	var imgBytes []byte
	for {
		// receive data from stream
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// return will close stream from server side
				s.logger.Debug("received image successfully")
				break
			}

			return errors.Wrap(err, "failed to receive image")
		}

		imgBytes = append(imgBytes, req.ImageData...)
	}

	// Recognize face
	recognized, err := s.service.RecognizeFace(ctx, imgBytes)
	if err != nil {
		return errors.Wrap(err, "failed to recognize face")
	}

	s.logger.Debugf("recognized person: %v", recognized)
	return stream.SendAndClose(&pb.RecognizeFaceResponse{
		PersonInfo: s.peopleConverter.PersonAsResponse(recognized),
	})
}

func (s *Server) AddFace(stream pb.FaceRegistry_AddFaceServer) error {
	s.logger.Debug("received 'AddFace' request")
	ctx := stream.Context()

	// Receive image upload
	var (
		img         []byte
		personIDStr string
	)

	for {
		// receive data from stream
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// return will close stream from server side
				s.logger.Debug("received image successfully")
				break
			}

			return errors.Wrap(err, "failed to receive image")
		}

		// Check if the person id is being received
		if req.PersonId != "" {
			personIDStr = req.PersonId

			s.logger.Debug("received person id, continuing...")
			continue
		}

		img = append(img, req.FaceImage.ImageData...)
	}

	// Add Face
	personID, err := uuid.Parse(personIDStr)
	if err != nil {
		return errors.Wrap(err, "failed to parse person id")
	}

	faceID, err := s.service.AddFace(ctx, img, personID)
	if err != nil {
		return errors.Wrap(err, "failed to add face")
	}

	s.logger.Debugf("added face id: %v", faceID)
	return stream.SendAndClose(&pb.AddFaceResponse{
		Id: faceID.String(),
	})
}

func (s *Server) RemoveFace(ctx context.Context, req *pb.RemoveFaceRequest) (*pb.RemoveFaceResponse, error) {
	s.logger.Debug("received 'RemoveFace' request")

	// TODO: Validate with `validate` tag?
	if req.FaceId == "" {
		return nil, errors.New("missing face id")
	}

	faceID, err := uuid.Parse(req.FaceId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse face id")
	}

	err = s.service.RemoveFace(ctx, faceID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to remove face")
	}

	s.logger.Debug("removed face successfully for id: %v", faceID)
	return &pb.RemoveFaceResponse{}, nil
}

func (s *Server) GetPeople(ctx context.Context, req *pb.GetPeopleRequest) (*pb.GetPeopleResponse, error) {
	s.logger.Debug("received 'GetPeople' request")

	people, err := s.service.GetPeople(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get people")
	}

	s.logger.Debugf("successfully got %d people", len(people))
	return &pb.GetPeopleResponse{
		People: s.peopleConverter.PeopleAsResponse(people),
	}, nil
}

func (s *Server) AddPerson(ctx context.Context, req *pb.AddPersonRequest) (*pb.AddPersonResponse, error) {
	s.logger.Debug("received 'AddPerson' request")

	id := uuid.New()
	s.logger.Debugf("generated person id: %v", id)

	err := s.service.AddPerson(ctx, s.peopleConverter.PersonAsRequest(pb.Person{
		Id:                   id.String(),
		FirstName:            req.FirstName,
		LastName:             req.LastName,
		Roles:                req.Roles,
	}))
	if err != nil {
		return nil, errors.Wrap(err, "failed to add person")
	}

	s.logger.Debug("person added successfully")
	return &pb.AddPersonResponse{
		PersonId: id.String(),
	}, nil
}

func (s *Server) RemovePerson(ctx context.Context, req *pb.RemovePersonRequest) (*pb.RemovePersonResponse, error) {
	s.logger.Debug("received 'RemovePerson' request")

	personID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse person id")
	}

	err = s.service.RemovePerson(ctx, personID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to remove person")
	}

	s.logger.Debugf("removed person successfully for id: %v", personID)
	return &pb.RemovePersonResponse{}, nil
}
