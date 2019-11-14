package face

import (
	"github.com/diogox/dom-face-registry/internal/face/recognizer"
	"github.com/google/uuid"
)

type Face struct {
	ID          uuid.UUID
	PersonID    uuid.UUID
	Encoding    recognizer.Encoding
	ImageData   []byte
	ImageFormat string
}
