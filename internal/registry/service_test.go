package registry

import (
	"context"
	"errors"
	"testing"

	"github.com/diogox/dom-face-recognizer/internal/person"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestService_RecognizeFace(t *testing.T) {
	reqImgBytes := []byte{0}

	t.Run("should recognize face", func(t *testing.T) {
		ctx := context.Background()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockFaceService := NewMockFaceService(ctrl)
		mockPersonService := NewMockPersonService(ctrl)
		service := NewRegistryService(logrus.New(), mockPersonService, mockFaceService)

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
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), nil, mockFaceService)

			expectedErr := errors.New("error")

			mockFaceService.EXPECT().RecognizeFace(ctx, reqImgBytes).Return(uuid.Nil, expectedErr)

			_, err := service.RecognizeFace(ctx, reqImgBytes)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})

		t.Run("person service returns an error", func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			mockPersonService := NewMockPersonService(ctrl)
			service := NewRegistryService(logrus.New(), mockPersonService, mockFaceService)


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
		ctx := context.Background()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockFaceService := NewMockFaceService(ctrl)
		service := NewRegistryService(logrus.New(), nil, mockFaceService)

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
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), nil, mockFaceService)

			expectedErr := errors.New("error")

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
		ctx := context.Background()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockFaceService := NewMockFaceService(ctrl)
		service := NewRegistryService(logrus.New(), nil, mockFaceService)

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

			reqID := uuid.Nil
			mockFaceService.EXPECT().RemoveFace(gomock.Any(), gomock.Any()).Times(0)

			err := service.RemoveFace(context.Background(), reqID)
			assert.Error(t, err)
			assert.Equal(t, err.Error(), errMissingFaceID)

		})

		t.Run("face service returns an error", func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFaceService := NewMockFaceService(ctrl)
			service := NewRegistryService(logrus.New(), nil, mockFaceService)

			expectedErr := errors.New("error")

			reqID := uuid.New()
			mockFaceService.EXPECT().RemoveFace(ctx, reqID).Return(expectedErr)

			err := service.RemoveFace(ctx, reqID)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})
	})
}
