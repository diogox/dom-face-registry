package mongo

import "github.com/diogox/dom-face-registry/internal/face/recognizer"

type Face struct {
	ID          string              `bson:"_id"`
	Encoding    recognizer.Encoding `bson:"encoding"`
	ImageData   []byte              `bson:"image_data"`
	ImageFormat string              `bson:"image_format"` // TODO: Add a way to specify as JPG, JPEG, PNG, etc...
	PersonID    string              `bson:"person_id"`
}
