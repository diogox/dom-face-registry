package mongo

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/diogox/dom-face-registry/internal/person"
)

// TODO: Ensure indexes?

type Store struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewStore(ctx context.Context, client *mongo.Client, dbName string, collectionName string) (*Store, error) {
	err := client.Ping(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMongoFailedToConnect)
	}

	db := client.Database(dbName)
	return &Store{
		db:         db,
		collection: db.Collection(collectionName),
	}, nil
}

func (s *Store) GetPeople(ctx context.Context) ([]person.Person, error) {
	var people []Person

	cur, err := s.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, errMongoFailedToGetPeople)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var p Person

		err := cur.Decode(&p)
		if err != nil {
			return nil, errors.Wrap(err, errMongoFailedToDecodePerson)
		}

		people = append(people, p)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, errMongoFailedToGetPeople)
	}

	res, err := s.mapPeopleToExternalType(people)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert people to expected type")
	}

	return res, nil
}

func (s *Store) mapPeopleToExternalType(people []Person) ([]person.Person, error) {
	pp := make([]person.Person, 0, len(people))

	for _, p := range people {
		pID, err := uuid.Parse(p.ID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse person id")
		}

		pp = append(pp, person.Person{
			ID:        pID,
			FirstName: p.FirstName,
			LastName:  p.LastName,
			Roles:     p.Roles,
		})
	}

	return pp, nil
}

func (s *Store) FindPersonByID(ctx context.Context, id uuid.UUID) (person.Person, error) {
	const IDKey = "_id"

	filter := bson.M{
		IDKey: id.String(),
	}

	var p Person
	err := s.collection.FindOne(ctx, filter).Decode(&p)
	if err != nil {
		return person.Person{}, errors.Wrap(err, errMongoFailedToGetPersonByID)
	}

	res, err := s.mapPeopleToExternalType([]Person{p})
	if err != nil {
		return person.Person{}, errors.Wrap(err, "failed to convert people to expected type")
	}

	return res[0], nil
}

func (s *Store) CreatePerson(ctx context.Context, person person.Person) error {
	// TODO: Make sure there are no other people with the same info

	_, err := s.collection.InsertOne(ctx, Person{
		ID:        person.ID.String(),
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Roles:     person.Roles,
	})
	if err != nil {
		return errors.Wrap(err, errMongoFailedToInsertPerson)
	}

	return nil
}

func (s *Store) DeletePerson(ctx context.Context, id uuid.UUID) error {
	const ID = "_id"

	_, err := s.collection.DeleteOne(ctx, bson.M{
		ID: id.String(),
	})
	if err != nil {
		return errors.Wrap(err, errMongoFailedToDeletePerson)
	}

	return nil
}
