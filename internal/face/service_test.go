package face

import (
	"context"
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
			ID: uuid.New(),
			Encoding: faceEncoding,
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
		peopleIDs := []uuid.UUID{ faces[0].ID }

		mockStore.EXPECT().GetFaces(ctx).Return(faces, nil)
		mockRecognizer.EXPECT().EncodeFace(reqImgBytes).Return(faceEncoding, nil)
		mockRecognizer.EXPECT().Recognize(faceEncoding, peopleIDs, faceEncoding).Return(personID, nil)

		res, err := service.RecognizeFace(ctx, reqImgBytes)
		assert.NoError(t, err)
		assert.Equal(t, res, personID)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("", func(t *testing.T) {

		})
	})
}
