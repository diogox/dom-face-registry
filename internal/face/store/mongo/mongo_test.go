// +build integration

package mongo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/diogox/dom-face-registry/internal/face"
	"github.com/diogox/dom-face-registry/internal/face/recognizer"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupStore(t *testing.T) (*Store, *mongo.Collection, func()) {
	t.Helper()

	ctx, cancelCtx := context.WithDeadline(context.Background(), time.Now().Add(5 * time.Second))

	const (
		dbName         = "dom-face-registry"
		collectionName = "faces"
	)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb://localhost:27017"),
	))
	require.NoError(t, err)

	store, err := NewStore(ctx, client, dbName, collectionName)
	require.NoError(t, err)

	collection := client.Database(dbName).Collection(collectionName)

	cleanup := func() {
		_, err := collection.DeleteMany(ctx, bson.M{})
		require.NoError(t, err)

		cancelCtx()
	}

	return store, collection, cleanup
}

func TestStore_GetFaces(t *testing.T) {
	expectedFace := face.Face{
		PersonID:    uuid.New(),
		Encoding:    recognizer.Encoding{},
		ImageData:   []byte{0},
		ImageFormat: "jpg",
	}

	t.Run("should get faces", func(t *testing.T) {
		store, _, cleanup := setupStore(t)
		defer cleanup()

		ctx := context.Background()
		expected := expectedFace

		faceID, err := store.AddFace(
			ctx,
			expected.Encoding,
			expected.ImageData,
			expected.PersonID,
		)
		require.NoError(t, err)
		expected.ID = faceID

		faces, err := store.GetFaces(ctx)
		assert.NoError(t, err)
		assert.Equal(t, len(faces), 1)
		assert.Equal(t, faces[0], expected)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("finds document with invalid format", func(t *testing.T) {
			store, collection, cleanup := setupStore(t)
			defer cleanup()

			ctx := context.Background()

			_, err := collection.InsertOne(ctx, bson.M{})
			require.NoError(t, err)

			_, err = store.GetFaces(ctx)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), errMongoFailedToDecodeFace)
		})
	})
}

func TestStore_FindFacesByPersonID(t *testing.T) {
	expectedFace := face.Face{
		PersonID:    uuid.New(),
		Encoding:    recognizer.Encoding{},
		ImageData:   []byte{0},
		ImageFormat: "jpg",
	}

	t.Run("should find faces by person id", func(t *testing.T) {
		store, _, cleanup := setupStore(t)
		defer cleanup()

		ctx := context.Background()
		expected := expectedFace

		_, err := store.AddFace(
			ctx,
			expected.Encoding,
			expected.ImageData,
			expected.PersonID,
		)
		require.NoError(t, err)

		faces, err := store.FindFacesByPersonID(ctx, expected.PersonID)
		assert.NoError(t, err)
		assert.Equal(t, len(faces), 1)
		assert.Equal(t, faces[0].PersonID, expected.PersonID)
	})

	t.Run("should return an empty slice if no faces are found", func(t *testing.T) {
		store, _, cleanup := setupStore(t)
		defer cleanup()

		ctx := context.Background()

		faces, err := store.FindFacesByPersonID(ctx, uuid.New())
		assert.NoError(t, err)
		assert.IsType(t, faces, []face.Face{})
		assert.Equal(t, len(faces), 0)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("finds document with invalid format", func(t *testing.T) {
			store, collection, cleanup := setupStore(t)
			defer cleanup()

			ctx := context.Background()

			_, err := collection.InsertOne(ctx, bson.M{})
			require.NoError(t, err)

			_, err = store.GetFaces(ctx)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), errMongoFailedToDecodeFace)
		})
	})
}

func TestStore_AddFace(t *testing.T) {
	expectedFace := face.Face{
		PersonID:    uuid.New(),
		Encoding:    recognizer.Encoding{},
		ImageData:   []byte{0},
		ImageFormat: "jpg",
	}

	t.Run("should add face", func(t *testing.T) {
		store, _, cleanup := setupStore(t)
		defer cleanup()

		ctx := context.Background()

		faceID, err := store.AddFace(
			ctx,
			expectedFace.Encoding,
			expectedFace.ImageData,
			expectedFace.PersonID,
		)
		assert.NoError(t, err)
		assert.NotEqual(t, faceID, uuid.Nil)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("person id is nil", func(t *testing.T) {
			store, _, cleanup := setupStore(t)
			defer cleanup()

			ctx := context.Background()

			_, err := store.AddFace(
				ctx,
				expectedFace.Encoding,
				expectedFace.ImageData,
				uuid.Nil,
			)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), errMongoInvalidPersonID)
		})
	})
}

func TestStore_RemoveFace(t *testing.T) {
	t.Run("should remove face", func(t *testing.T) {
		store, _, cleanup := setupStore(t)
		defer cleanup()

		ctx := context.Background()

		faceID, err := store.AddFace(
			ctx,
			recognizer.Encoding{},
			[]byte{0},
			uuid.New(),
		)
		require.NoError(t, err)

		err = store.RemoveFace(ctx, faceID)
		assert.NoError(t, err)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("face id doesn't exist", func(t *testing.T) {
			store, _, cleanup := setupStore(t)
			defer cleanup()

			ctx := context.Background()

			err := store.RemoveFace(ctx, uuid.New())
			assert.Error(t, err)
			assert.Contains(t, err.Error(), errMongoFailedToRemoveFace)
		})
	})
}
