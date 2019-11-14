package person

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_GetPeople(t *testing.T) {
	person := Person{ ID: uuid.New() }
	people := []Person{ person }

	t.Run("should get people", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockStore := NewMockStore(ctrl)
		service := NewService(mockStore)

		ctx := context.Background()
		mockStore.EXPECT().GetPeople(ctx).Return(people, nil)

		res, err := service.GetPeople(ctx)
		assert.NoError(t, err)
		assert.Equal(t, res, people)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("store returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := NewMockStore(ctrl)
			service := NewService(mockStore)

			ctx := context.Background()
			expectedErr := errors.New("error")

			mockStore.EXPECT().GetPeople(ctx).Return([]Person{}, expectedErr)

			_, err := service.GetPeople(ctx)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})
	})
}

func TestService_FindPersonByID(t *testing.T) {
	reqID := uuid.New()
	person := Person{ ID: uuid.New() }

	t.Run("should find person by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockStore := NewMockStore(ctrl)
		service := NewService(mockStore)

		ctx := context.Background()

		mockStore.EXPECT().FindPersonByID(ctx, reqID).Return(person, nil)

		res, err := service.FindPersonByID(ctx, reqID)
		assert.NoError(t, err)
		assert.Equal(t, res, person)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("store returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := NewMockStore(ctrl)
			service := NewService(mockStore)

			ctx := context.Background()
			expectedErr := errors.New("error")

			mockStore.EXPECT().FindPersonByID(ctx, reqID).Return(Person{}, expectedErr)

			_, err := service.FindPersonByID(ctx, reqID)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})
	})
}

func TestService_AddPerson(t *testing.T) {
	reqPerson := Person{ ID: uuid.New() }

	t.Run("should add person", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockStore := NewMockStore(ctrl)
		service := NewService(mockStore)

		ctx := context.Background()

		mockStore.EXPECT().CreatePerson(ctx, reqPerson).Return(nil)

		err := service.AddPerson(ctx, reqPerson)
		assert.NoError(t, err)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("store returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := NewMockStore(ctrl)
			service := NewService(mockStore)

			ctx := context.Background()
			expectedErr := errors.New("error")

			mockStore.EXPECT().CreatePerson(ctx, reqPerson).Return(expectedErr)

			err := service.AddPerson(ctx, reqPerson)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})
	})
}



func TestService_RemovePerson(t *testing.T) {
	reqID := uuid.New()

	t.Run("should remove person", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockStore := NewMockStore(ctrl)
		service := NewService(mockStore)

		ctx := context.Background()
		mockStore.EXPECT().DeletePerson(ctx, reqID).Return(nil)

		err := service.RemovePerson(ctx, reqID)
		assert.NoError(t, err)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("store returns an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := NewMockStore(ctrl)
			service := NewService(mockStore)

			ctx := context.Background()
			expectedErr := errors.New("error")

			mockStore.EXPECT().DeletePerson(ctx, reqID).Return(expectedErr)

			err := service.RemovePerson(ctx, reqID)
			assert.Error(t, err)
			assert.Equal(t, err, expectedErr)
		})
	})
}
