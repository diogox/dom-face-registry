package face

import (
	"github.com/diogox/dom-face-recognizer/internal/face/recognizer"
	"github.com/google/uuid"
)

type Face struct {
	ID          uuid.UUID
	PersonID    uuid.UUID
	Encoding    recognizer.FaceEncoding
	ImageData   []byte
	ImageFormat string
}
