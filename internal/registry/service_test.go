package registry

import (
	"context"
	"errors"
	"github.com/diogox/dom-face-registry/internal/face"
	"testing"

	"github.com/diogox/dom-face-registry/internal/person"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestService_RecognizeFace(t *testing.T) {
	reqImgBytes := []byte{0}

	t.Run("should recognize face", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockFaceService := NewMockFaceService(ctrl)
		mockPersonService := NewMockPersonService(ctrl)
		service := NewRegistryService(logrus.New(), mockPersonService, mockFaceService)

		ctx := context.Background()
		expectedPersonID := uuid.New()

		mockFaceService.EXPECT().RecognizeFace(ctx, reqImgBytes).Return(expectedPersonID, nil)

		mockPersonService.EXPECT().FindPersonByID(ctx, gomock.Any()).
			Return(person.Person{
				ID: expectedPersonID,
			}, nil)

		res, err := service.RecognizeFace(ctx, reqImgBytes)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.ID, expectedPersonID)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("image provided is nil", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), nil, mockFaceService)

			mockFaceService.EXPECT().RecognizeFace(gomock.Any(), gomock.Any()).Times(0)

			_, err := service.RecognizeFace(context.Background(), nil)
			assert.Error(t, err)
			assert.Equal(t, err.Error(), errMissingImage)
		})

		t.Run("image provided is empty slice of bytes", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), nil, mockFaceService)

			mockFaceService.EXPECT().RecognizeFace(gomock.Any(), gomock.Any()).Times(0)

			_, err := service.RecognizeFace(context.Background(), []byte{})
			assert.Error(t, err)
			assert.Equal(t, err.Error(), errMissingImage)
		})

		t.Run("face service returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), nil, mockFaceService)

			ctx := context.Background()
			expectedErr := errors.New("error")

			mockFaceService.EXPECT().RecognizeFace(ctx, reqImgBytes).Return(uuid.Nil, expectedErr)

			_, err := service.RecognizeFace(ctx, reqImgBytes)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})

		t.Run("person service returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			mockPersonService := NewMockPersonService(ctrl)
			service := NewRegistryService(logrus.New(), mockPersonService, mockFaceService)

			ctx := context.Background()
			expectedPersonID := uuid.New()

			mockFaceService.EXPECT().RecognizeFace(ctx, reqImgBytes).Return(expectedPersonID, nil)

			mockPersonService.EXPECT().FindPersonByID(ctx, expectedPersonID).Return(person.Person{}, errors.New("error"))

			_, err := service.RecognizeFace(ctx, reqImgBytes)
			assert.Error(t, err)
		})
	})
}

func TestService_AddFace(t *testing.T) {
	reqImgBytes := []byte{0}
	reqPersonID := uuid.New()

	t.Run("should add face", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockFaceService := NewMockFaceService(ctrl)
		service := NewRegistryService(logrus.New(), nil, mockFaceService)

		ctx := context.Background()
		resFaceID := uuid.New()

		mockFaceService.EXPECT().AddFace(ctx, reqImgBytes, reqPersonID).Return(resFaceID, nil)

		res, err := service.AddFace(ctx, reqImgBytes, reqPersonID)
		assert.NoError(t, err)
		assert.Equal(t, res, resFaceID)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("face image provided is nil", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), nil, mockFaceService)

			mockFaceService.EXPECT().AddFace(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

			_, err := service.AddFace(context.Background(), nil, reqPersonID)
			assert.Error(t, err)
			assert.Equal(t, err.Error(), errMissingImage)
		})

		t.Run("face image provided is empty slice of bytes", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), nil, mockFaceService)

			mockFaceService.EXPECT().AddFace(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

			_, err := service.AddFace(context.Background(), []byte{}, reqPersonID)
			assert.Error(t, err)
			assert.Equal(t, err.Error(), errMissingImage)
		})

		t.Run("face service returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), nil, mockFaceService)

			expectedErr := errors.New("error")

			ctx := context.Background()
			reqID := uuid.New()

			mockFaceService.EXPECT().RemoveFace(ctx, reqID).Return(expectedErr)

			err := service.RemoveFace(ctx, reqID)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})
	})
}

func TestService_RemoveFace(t *testing.T) {
	t.Run("should remove face", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockFaceService := NewMockFaceService(ctrl)
		service := NewRegistryService(logrus.New(), nil, mockFaceService)

		ctx := context.Background()
		reqFaceID := uuid.New()

		mockFaceService.EXPECT().RemoveFace(ctx, reqFaceID).Return(nil)

		err := service.RemoveFace(ctx, reqFaceID)
		assert.NoError(t, err)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("face id provided is nil", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), nil, mockFaceService)

			mockFaceService.EXPECT().RemoveFace(gomock.Any(), gomock.Any()).Times(0)

			err := service.RemoveFace(context.Background(), uuid.Nil)
			assert.Error(t, err)
			assert.Equal(t, err.Error(), errMissingFaceID)

		})

		t.Run("face service returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), nil, mockFaceService)

			expectedErr := errors.New("error")

			ctx := context.Background()
			reqID := uuid.New()

			mockFaceService.EXPECT().RemoveFace(ctx, reqID).Return(expectedErr)

			err := service.RemoveFace(ctx, reqID)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})
	})
}

func TestService_GetPeople(t *testing.T) {
	expectedPerson := person.Person{ID: uuid.New()}

	t.Run("should get people", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockPersonService := NewMockPersonService(ctrl)
		service := NewRegistryService(logrus.New(), mockPersonService, nil)

		ctx := context.Background()
		expectedPeople := []person.Person{expectedPerson}

		mockPersonService.EXPECT().GetPeople(ctx).Return(expectedPeople, nil)

		res, err := service.GetPeople(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, res[0].ID, expectedPerson.ID)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("person service returns an error", func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPersonService := NewMockPersonService(ctrl)
			service := NewRegistryService(logrus.New(), mockPersonService, nil)

			expectedErr := errors.New("error")
			mockPersonService.EXPECT().GetPeople(ctx).Return(nil, expectedErr)

			_, err := service.GetPeople(ctx)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})
	})
}

func TestService_AddPerson(t *testing.T) {
	reqPerson := person.Person{ID: uuid.New()}

	t.Run("should add person", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockPersonService := NewMockPersonService(ctrl)
		service := NewRegistryService(logrus.New(), mockPersonService, nil)

		ctx := context.Background()
		mockPersonService.EXPECT().AddPerson(ctx, reqPerson).Return(nil)

		err := service.AddPerson(ctx, reqPerson)
		assert.NoError(t, err)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("empty request body", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPersonService := NewMockPersonService(ctrl)
			service := NewRegistryService(logrus.New(), mockPersonService, nil)

			mockPersonService.EXPECT().AddPerson(gomock.Any(), gomock.Any()).Times(0)

			err := service.AddPerson(context.Background(), person.Person{})
			assert.Error(t, err)
			assert.Equal(t, err.Error(), errMissingPersonInfo)
		})

		t.Run("person service returns an error", func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPersonService := NewMockPersonService(ctrl)
			service := NewRegistryService(logrus.New(), mockPersonService, nil)

			expectedErr := errors.New("error")
			mockPersonService.EXPECT().AddPerson(ctx, reqPerson).Return(expectedErr)

			err := service.AddPerson(ctx, reqPerson)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})
	})
}

func TestService_RemovePerson(t *testing.T) {
	personFace := face.Face{
		ID: uuid.New(),
	}

	t.Run("should remove person", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockPersonService := NewMockPersonService(ctrl)
		mockFaceService := NewMockFaceService(ctrl)
		service := NewRegistryService(logrus.New(), mockPersonService, mockFaceService)

		ctx := context.Background()
		reqID := uuid.New()

		mockFaceService.EXPECT().FindFacesByPersonID(ctx, reqID).Return([]face.Face{personFace}, nil)
		mockFaceService.EXPECT().RemoveFace(ctx, personFace.ID).Return(nil)
		mockPersonService.EXPECT().RemovePerson(ctx, reqID).Return(nil)

		err := service.RemovePerson(ctx, reqID)
		assert.NoError(t, err)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("face service returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), nil, mockFaceService)

			ctx := context.Background()
			reqID := uuid.New()

			mockFaceService.EXPECT().FindFacesByPersonID(ctx, reqID).Return(nil, errors.New("error"))

			err := service.RemovePerson(ctx, reqID)
			assert.Error(t, err)
		})

		t.Run("person service returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPersonService := NewMockPersonService(ctrl)
			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), mockPersonService, mockFaceService)

			expectedErr := errors.New("error")

			ctx := context.Background()
			reqID := uuid.New()

			mockFaceService.EXPECT().FindFacesByPersonID(ctx, reqID).Return([]face.Face{personFace}, nil)
			mockFaceService.EXPECT().RemoveFace(ctx, gomock.Any()).Return(nil)
			mockPersonService.EXPECT().RemovePerson(ctx, reqID).Return(expectedErr)

			err := service.RemovePerson(ctx, reqID)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})

		t.Run("person id is nil", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPersonService := NewMockPersonService(ctrl)
			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), mockPersonService, mockFaceService)

			mockFaceService.EXPECT().FindFacesByPersonID(gomock.Any(), gomock.Any()).Times(0)
			mockFaceService.EXPECT().RemoveFace(gomock.Any(), gomock.Any()).Times(0)
			mockPersonService.EXPECT().RemovePerson(gomock.Any(), gomock.Any()).Times(0)

			err := service.RemovePerson(context.Background(), uuid.Nil)
			assert.Error(t, err)
		})
	})
}
