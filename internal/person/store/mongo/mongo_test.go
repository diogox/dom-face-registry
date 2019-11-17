// +build integration

package mongo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/diogox/dom-face-registry/internal/person"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupStore(t *testing.T) (*Store, *mongo.Collection, func()) {
	t.Helper()

	ctx, cancelCtx := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))

	const (
		dbName         = "dom-face-registry"
		collectionName = "test-people"
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

func TestStore_GetPeople(t *testing.T) {
	expectedPerson := person.Person{
		ID:        uuid.New(),
		FirstName: "First",
		LastName:  "Last",
		Roles:     []string{},
	}

	t.Run("should get people", func(t *testing.T) {
		store, _, cleanup := setupStore(t)
		defer cleanup()

		ctx := context.Background()

		err := store.CreatePerson(ctx, expectedPerson)
		require.NoError(t, err)

		people, err := store.GetPeople(ctx)
		assert.NoError(t, err)
		assert.Equal(t, len(people), 1)
		assert.Equal(t, people[0], expectedPerson)
	})

	t.Run("should return an empty slice if no people are found", func(t *testing.T) {
		store, _, cleanup := setupStore(t)
		defer cleanup()

		ctx := context.Background()

		people, err := store.GetPeople(ctx)
		assert.NoError(t, err)
		assert.IsType(t, people, []person.Person{})
		assert.Equal(t, len(people), 0)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("finds document with invalid format", func(t *testing.T) {
			store, collection, cleanup := setupStore(t)
			defer cleanup()

			ctx := context.Background()

			_, err := collection.InsertOne(ctx, bson.M{})
			require.NoError(t, err)

			_, err = store.GetPeople(ctx)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), errMongoFailedToDecodePerson)
		})
	})
}

func TestStore_FindPersonByID(t *testing.T) {
	expectedPerson := person.Person{
		ID:        uuid.New(),
		FirstName: "First",
		LastName:  "Last",
		Roles:     []string{},
	}

	t.Run("should find person by id", func(t *testing.T) {
		store, _, cleanup := setupStore(t)
		defer cleanup()

		ctx := context.Background()

		err := store.CreatePerson(ctx, expectedPerson)
		require.NoError(t, err)

		foundPerson, err := store.FindPersonByID(ctx, expectedPerson.ID)
		assert.NoError(t, err)
		assert.Equal(t, foundPerson, expectedPerson)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("finds document with invalid format", func(t *testing.T) {
			store, collection, cleanup := setupStore(t)
			defer cleanup()

			ctx := context.Background()

			_, err := collection.InsertOne(ctx, bson.M{})
			require.NoError(t, err)

			_, err = store.FindPersonByID(ctx, expectedPerson.ID)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), errMongoFailedToGetPersonByID)
		})
	})
}

func TestStore_CreatePerson(t *testing.T) {
	expectedPerson := person.Person{
		ID:        uuid.New(),
		FirstName: "First",
		LastName:  "Last",
		Roles:     []string{},
	}

	t.Run("should create person", func(t *testing.T) {
		store, _, cleanup := setupStore(t)
		defer cleanup()

		ctx := context.Background()

		err := store.CreatePerson(ctx, expectedPerson)
		assert.NoError(t, err)

		foundPerson, err := store.FindPersonByID(ctx, expectedPerson.ID)
		require.NoError(t, err)

		assert.Equal(t, foundPerson, expectedPerson)
	})
}

func TestStore_DeletePerson(t *testing.T) {
	expectedPerson := person.Person{
		ID:        uuid.New(),
		FirstName: "First",
		LastName:  "Last",
		Roles:     []string{},
	}

	t.Run("should delete person", func(t *testing.T) {
		store, _, cleanup := setupStore(t)
		defer cleanup()

		ctx := context.Background()

		err := store.CreatePerson(ctx, expectedPerson)
		require.NoError(t, err)

		err = store.DeletePerson(ctx, expectedPerson.ID)
		assert.NoError(t, err)
	})

	t.Run("should fail when", func(t *testing.T) {
		t.Run("person id doesn't exist", func(t *testing.T) {
			store, _, cleanup := setupStore(t)
			defer cleanup()

			ctx := context.Background()

			err := store.DeletePerson(ctx, uuid.New())
			assert.Error(t, err)
			assert.Contains(t, err.Error(), errMongoFailedToDeletePerson)
		})
	})
}
