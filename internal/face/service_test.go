package face

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/diogox/dom-face-registry/internal/face/recognizer"
)

func TestService_RecognizeFace(t *testing.T) {
	reqImgBytes := []byte{0}
	faceEncoding := recognizer.Encoding{}
	faces := []Face{
		{
			ID:       uuid.New(),
			Encoding: faceEncoding,
			PersonID: uuid.New(),
		},
	}
	personID := uuid.New()

	t.Run("should recognize face", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockStore := NewMockStore(ctrl)
		mockRecognizer := recognizer.NewMockRecognizer(ctrl)
		service := NewService(mockStore, mockRecognizer)

		ctx := context.Background()
		peopleIDs := []uuid.UUID{faces[0].PersonID}
		allFaceEncodings := []recognizer.Encoding{faces[0].Encoding}

		mockStore.EXPECT().GetFaces(ctx).Return(faces, nil)
		mockRecognizer.EXPECT().EncodeFace(reqImgBytes).Return(faceEncoding, nil)
		mockRecognizer.EXPECT().Recognize(faceEncoding, peopleIDs, allFaceEncodings).Return(personID, nil)

		res, err := service.RecognizeFace(ctx, reqImgBytes)
		assert.NoError(t, err)
		assert.Equal(t, res, personID)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("no faces exist in store", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := NewMockStore(ctrl)
			mockRecognizer := recognizer.NewMockRecognizer(ctrl)
			service := NewService(mockStore, mockRecognizer)

			ctx := context.Background()

			mockStore.EXPECT().GetFaces(ctx).Return([]Face{}, nil)
			mockRecognizer.EXPECT().EncodeFace(gomock.Any()).Times(0)
			mockRecognizer.EXPECT().Recognize(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

			_, err := service.RecognizeFace(ctx, reqImgBytes)
			assert.Error(t, err)
			assert.Equal(t, err.Error(), errNoFacesFound)
		})

		t.Run("store returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := NewMockStore(ctrl)
			mockRecognizer := recognizer.NewMockRecognizer(ctrl)
			service := NewService(mockStore, mockRecognizer)

			ctx := context.Background()
			expectedErr := errors.New("error")

			mockStore.EXPECT().GetFaces(ctx).Return(nil, expectedErr)
			mockRecognizer.EXPECT().Recognize(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

			_, err := service.RecognizeFace(ctx, reqImgBytes)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})

		t.Run("recognizer returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := NewMockStore(ctrl)
			mockRecognizer := recognizer.NewMockRecognizer(ctrl)
			service := NewService(mockStore, mockRecognizer)

			ctx := context.Background()
			expectedErr := errors.New("error")

			peopleIDs := []uuid.UUID{faces[0].PersonID}
			allFaceEncodings := []recognizer.Encoding{faces[0].Encoding}

			mockStore.EXPECT().GetFaces(ctx).Return(faces, nil)
			mockRecognizer.EXPECT().EncodeFace(reqImgBytes).Return(faceEncoding, nil)
			mockRecognizer.EXPECT().Recognize(faceEncoding, peopleIDs, allFaceEncodings).Return(uuid.Nil, expectedErr)

			_, err := service.RecognizeFace(ctx, reqImgBytes)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), expectedErr.Error())
		})
	})
}

func TestService_FindFacesByPersonID(t *testing.T) {
	reqID := uuid.New()
	face := Face{ID: uuid.New()}
	faces := []Face{face}

	t.Run("should find faces", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockStore := NewMockStore(ctrl)
		service := NewService(mockStore, nil)

		ctx := context.Background()
		mockStore.EXPECT().FindFacesByPersonID(ctx, reqID).Return(faces, nil)

		res, err := service.FindFacesByPersonID(ctx, reqID)
		assert.NoError(t, err)
		assert.Equal(t, res, faces)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("store returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := NewMockStore(ctrl)
			service := NewService(mockStore, nil)

			ctx := context.Background()
			expectedErr := errors.New("error")

			mockStore.EXPECT().FindFacesByPersonID(ctx, reqID).Return(nil, expectedErr)

			_, err := service.FindFacesByPersonID(ctx, reqID)
			assert.Error(t, err, expectedErr)
		})
	})
}

func TestService_AddFace(t *testing.T) {
	reqImgBytes := []byte{0}
	faceEncoding := recognizer.Encoding{}
	reqPersonID := uuid.New()

	t.Run("should add face", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockStore := NewMockStore(ctrl)
		mockRecognizer := recognizer.NewMockRecognizer(ctrl)
		service := NewService(mockStore, mockRecognizer)

		ctx := context.Background()
		expectedFaceID := uuid.New()

		mockRecognizer.EXPECT().EncodeFace(reqImgBytes).Return(faceEncoding, nil)
		mockStore.EXPECT().AddFace(ctx, faceEncoding, reqImgBytes, reqPersonID).Return(expectedFaceID, nil)

		res, err := service.AddFace(ctx, reqImgBytes, reqPersonID)
		assert.NoError(t, err)
		assert.Equal(t, res, expectedFaceID)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("store returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := NewMockStore(ctrl)
			mockRecognizer := recognizer.NewMockRecognizer(ctrl)
			service := NewService(mockStore, mockRecognizer)

			ctx := context.Background()
			expectedErr := errors.New("error")

			mockRecognizer.EXPECT().EncodeFace(reqImgBytes).Return(recognizer.Encoding{}, expectedErr)
			mockStore.EXPECT().AddFace(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

			_, err := service.AddFace(ctx, reqImgBytes, reqPersonID)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), expectedErr.Error())
		})

		t.Run("recognizer returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := NewMockStore(ctrl)
			mockRecognizer := recognizer.NewMockRecognizer(ctrl)
			service := NewService(mockStore, mockRecognizer)

			ctx := context.Background()
			expectedErr := errors.New("error")

			mockRecognizer.EXPECT().EncodeFace(reqImgBytes).Return(faceEncoding, nil)
			mockStore.EXPECT().AddFace(ctx, faceEncoding, reqImgBytes, reqPersonID).Return(uuid.Nil, expectedErr)

			_, err := service.AddFace(ctx, reqImgBytes, reqPersonID)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})
	})
}

func TestService_RemoveFace(t *testing.T) {
	reqFaceID := uuid.New()

	t.Run("should remove face", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockStore := NewMockStore(ctrl)
		service := NewService(mockStore, nil)

		ctx := context.Background()

		mockStore.EXPECT().RemoveFace(ctx, reqFaceID).Return(nil)

		err := service.RemoveFace(ctx, reqFaceID)
		assert.NoError(t, err)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("store returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := NewMockStore(ctrl)
			service := NewService(mockStore, nil)

			ctx := context.Background()
			expectedErr := errors.New("error")

			mockStore.EXPECT().RemoveFace(ctx, reqFaceID).Return(expectedErr)

			err := service.RemoveFace(ctx, reqFaceID)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})
	})
}
