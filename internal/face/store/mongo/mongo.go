package mongo

import (
	"context"
	"github.com/pkg/errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/diogox/dom-face-registry/internal/face"
	"github.com/diogox/dom-face-registry/internal/face/recognizer"
)

// TODO: Ensure indexes?

type Store struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewStore(ctx context.Context, client *mongo.Client, dbName, collectionName string) (*Store, error) {
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

func (s *Store) GetFaces(ctx context.Context) ([]face.Face, error) {
	var faces []face.Face

	cur, err := s.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, errMongoFailedToGetFaces)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var f Face

		err := cur.Decode(&f)
		if err != nil {
			return nil, errors.Wrap(err, errMongoFailedToDecodeFace)
		}

		fID, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse face id")
		}

		fPersonID, err := uuid.Parse(f.PersonID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse person id")
		}

		faces = append(faces, face.Face{
			ID:          fID,
			PersonID:    fPersonID,
			Encoding:    f.Encoding,
			ImageData:   f.ImageData,
			ImageFormat: f.ImageFormat,
		})
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, errMongoFailedToGetFaces)
	}

	return faces, nil
}

func (s *Store) FindFacesByPersonID(ctx context.Context, personID uuid.UUID) ([]face.Face, error) {
	const PersonIDKey = "person_id"

	cur, err := s.collection.Find(ctx, bson.M{
		PersonIDKey: personID,
	})
	if err != nil {
		return nil, errors.Wrap(err, errMongoFailedToGetFaces)
	}
	defer cur.Close(ctx)

	var faces []face.Face
	for cur.Next(ctx) {
		var f Face

		err := cur.Decode(&f)
		if err != nil {
			return nil, errors.Wrap(err, errMongoFailedToDecodeFace)
		}

		fID, err := uuid.Parse(f.ID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse face id")
		}

		fPersonID, err := uuid.Parse(f.PersonID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse person id")
		}

		faces = append(faces, face.Face{
			ID:          fID,
			PersonID:    fPersonID,
			Encoding:    f.Encoding,
			ImageData:   f.ImageData,
			ImageFormat: f.ImageFormat,
		})
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, errMongoFailedToGetFaces)
	}

	return faces, nil
}

func (s *Store) AddFace(
	ctx context.Context,
	encoding recognizer.Encoding,
	imgBytes []byte,
	personUID uuid.UUID,
) (uuid.UUID, error) {

	insertedResult, err := s.collection.InsertOne(ctx, Face{
		ID:          uuid.New().String(),
		Encoding:    encoding,
		ImageData:   imgBytes,
		ImageFormat: "jpg",
		PersonID:    personUID.String(),
	})
	if err != nil {
		return uuid.Nil, errors.Wrap(err, errMongoFailedToInsertFace)
	}

	objID, ok := insertedResult.InsertedID.(string)
	if !ok {
		return uuid.Nil, errors.New(errMongoInvalidInsertedFaceID)
	}

	faceID, err := uuid.Parse(objID)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, errMongoInvalidInsertedFaceID)
	}

	return faceID, nil
}

func (s *Store) RemoveFace(ctx context.Context, faceID uuid.UUID) error {
	const ID = "_id"

	_, err := s.collection.DeleteOne(ctx, bson.M{
		ID: faceID,
	})
	if err != nil {
		return errors.Wrap(err, errMongoFailedToRemoveFace)
	}

	return nil
}
