package grpc

import (
	"context"
	"errors"
	"github.com/diogox/dom-face-registry/internal/person"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	pb "github.com/diogox/dom-face-registry/internal/pb"
)

func TestServer_RemoveFace(t *testing.T) {
	faceID := uuid.New()
	mockReq := &pb.RemoveFaceRequest{
		FaceId: faceID.String(),
	}

	t.Run("should recognize face", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRegistry := NewMockRegistryService(ctrl)
		mockPeopleConverter := NewMockPeopleConverter(ctrl)
		service := NewServer(logrus.New(), mockRegistry, mockPeopleConverter)

		ctx := context.Background()
		mockRegistry.EXPECT().RemoveFace(ctx, faceID).Return(nil)

		res, err := service.RemoveFace(ctx, mockReq)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("service returns error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRegistry := NewMockRegistryService(ctrl)
			mockPeopleConverter := NewMockPeopleConverter(ctrl)
			service := NewServer(logrus.New(), mockRegistry, mockPeopleConverter)

			ctx := context.Background()
			expectedErr := errors.New("error")

			mockRegistry.EXPECT().RemoveFace(ctx, faceID).Return(expectedErr)

			_, err := service.RemoveFace(ctx, mockReq)
			assert.Error(t, err)
		})

		t.Run("id cannot be parsed", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRegistry := NewMockRegistryService(ctrl)
			mockPeopleConverter := NewMockPeopleConverter(ctrl)
			service := NewServer(logrus.New(), mockRegistry, mockPeopleConverter)

			mockRegistry.EXPECT().RemoveFace(gomock.Any(), gomock.Any()).Times(0)

			_, err := service.RemoveFace(context.Background(), &pb.RemoveFaceRequest{
				FaceId: "1",
			})
			assert.Error(t, err)
		})

		t.Run("id is missing", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRegistry := NewMockRegistryService(ctrl)
			mockPeopleConverter := NewMockPeopleConverter(ctrl)
			service := NewServer(logrus.New(), mockRegistry, mockPeopleConverter)

			mockRegistry.EXPECT().RemoveFace(gomock.Any(), gomock.Any()).Times(0)

			_, err := service.RemoveFace(context.Background(), &pb.RemoveFaceRequest{
				FaceId: "",
			})
			assert.Error(t, err)
		})
	})
}

func TestServer_GetPeople(t *testing.T) {
	personID := uuid.New()
	mockPerson := person.Person{
		ID:        personID,
		FirstName: "first",
		LastName:  "last",
		Roles:     []string{person.RoleInhabitant},
	}
	mockPersonResponse := pb.Person{
		Id:        personID.String(),
		FirstName: "first",
		LastName:  "last",
		Roles:     []pb.PersonRole{pb.PersonRole_INHABITANT},
	}

	t.Run("should get people", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRegistry := NewMockRegistryService(ctrl)
		mockPeopleConverter := NewMockPeopleConverter(ctrl)
		service := NewServer(logrus.New(), mockRegistry, mockPeopleConverter)

		ctx := context.Background()
		mockRegistry.EXPECT().GetPeople(ctx).Return([]person.Person{mockPerson}, nil)
		mockPeopleConverter.EXPECT().PeopleAsResponse(gomock.Any()).Return([]*pb.Person{&mockPersonResponse})

		res, err := service.GetPeople(ctx, &pb.GetPeopleRequest{})
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("service returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRegistry := NewMockRegistryService(ctrl)
			mockPeopleConverter := NewMockPeopleConverter(ctrl)
			service := NewServer(logrus.New(), mockRegistry, mockPeopleConverter)

			ctx := context.Background()
			expectedErr := errors.New("error")

			mockRegistry.EXPECT().GetPeople(ctx).Return(nil, expectedErr)
			mockPeopleConverter.EXPECT().PeopleAsResponse(gomock.Any()).Times(0)

			_, err := service.GetPeople(ctx, &pb.GetPeopleRequest{})
			assert.Error(t, err)
			assert.Contains(t, err.Error(), expectedErr.Error())
		})
	})
}

func TestServer_AddPerson(t *testing.T) {
	mockReq := &pb.AddPersonRequest{
		FirstName: "first",
		LastName:  "last",
		Roles:     []pb.PersonRole{pb.PersonRole_INHABITANT},
	}
	mockPerson := person.Person{
		ID:        uuid.New(),
		FirstName: "first",
		LastName:  "last",
		Roles:     []string{person.RoleInhabitant},
	}

	t.Run("should add person", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRegistry := NewMockRegistryService(ctrl)
		mockPeopleConverter := NewMockPeopleConverter(ctrl)
		service := NewServer(logrus.New(), mockRegistry, mockPeopleConverter)

		ctx := context.Background()
		mockPeopleConverter.EXPECT().PersonAsRequest(gomock.Any()).Return(mockPerson)
		mockRegistry.EXPECT().AddPerson(ctx, gomock.Any()).Return(nil)

		res, err := service.AddPerson(ctx, mockReq)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("service returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRegistry := NewMockRegistryService(ctrl)
			mockPeopleConverter := NewMockPeopleConverter(ctrl)
			service := NewServer(logrus.New(), mockRegistry, mockPeopleConverter)

			expectedErr := errors.New("error")

			mockPeopleConverter.EXPECT().PersonAsRequest(gomock.Any()).Return(mockPerson)
			mockRegistry.EXPECT().AddPerson(gomock.Any(), gomock.Any()).Return(expectedErr)

			_, err := service.AddPerson(context.Background(), mockReq)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), expectedErr.Error())
		})
	})
}

func TestServer_RemovePerson(t *testing.T) {
	personID := uuid.New()

	t.Run("should remove person", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRegistry := NewMockRegistryService(ctrl)
		service := NewServer(logrus.New(), mockRegistry, nil)

		ctx := context.Background()
		mockRegistry.EXPECT().RemovePerson(ctx, personID).Return(nil)

		_, err := service.RemovePerson(ctx, &pb.RemovePersonRequest{
			Id: personID.String(),
		})
		assert.NoError(t, err)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("service returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRegistry := NewMockRegistryService(ctrl)
			service := NewServer(logrus.New(), mockRegistry, nil)

			ctx := context.Background()
			expectedErr := errors.New("error")

			mockRegistry.EXPECT().RemovePerson(ctx, gomock.Any()).Return(expectedErr)

			_, err := service.RemovePerson(ctx, &pb.RemovePersonRequest{
				Id: personID.String(),
			})
			assert.Error(t, err)
			assert.Contains(t, err.Error(), expectedErr.Error())
		})
	})
}
