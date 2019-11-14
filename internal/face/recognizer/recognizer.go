//go:generate mockgen -package=recognizer -source=./recognizer.go -destination=./recognizer_mock.go -self_package=github.com/diogox/dom-face-registry/internal/face/recognizer

package recognizer

import (
	goface "github.com/Kagami/go-face"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Encoding goface.Descriptor

type Recognizer interface {
	Recognize(targetFaceEncoding Encoding, peopleIDs []uuid.UUID, allFaceEncondings []Encoding) (uuid.UUID, error)
	EncodeFace(imgBytes []byte) (Encoding, error)
}

type recognizer struct {
	recognizer        *goface.Recognizer
	classifyThreshold float32
}

func NewRecognizer(dlibModelsPath string, classifyThreshold float32) (*recognizer, error) {
	rec, err := goface.NewRecognizer(dlibModelsPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed at initializing face recognizer")
	}

	return &recognizer{
		recognizer:        rec,
		classifyThreshold: classifyThreshold,
	}, nil
}

func (r *recognizer) Recognize(
	targetFaceEnding Encoding,
	peopleIDs []uuid.UUID,
	allFaceEncodings []Encoding,
) (uuid.UUID, error) {

	// Set Samples
	var samples []goface.Descriptor
	for _, fe := range allFaceEncodings {
		samples = append(samples, goface.Descriptor(fe))
	}

	categories := make([]int32, 0, len(peopleIDs))
	categoriesLabelsMap := make(map[int]uuid.UUID, len(peopleIDs))
	for i, label := range peopleIDs {
		categoriesLabelsMap[i] = label
		categories = append(categories, int32(i))
	}

	r.recognizer.SetSamples(samples, categories)

	categoryID := r.recognizer.ClassifyThreshold(goface.Descriptor(targetFaceEnding), r.classifyThreshold)
	if categoryID < 0 {
		return uuid.Nil, errors.New("unable to classify image")
	}

	return categoriesLabelsMap[categoryID], nil
}

func (r *recognizer) EncodeFace(imgBytes []byte) (Encoding, error) {
	f, err := r.recognizer.RecognizeSingle(imgBytes)
	if err != nil {
		return Encoding{}, errors.Wrap(err, "failed to encode face")
	}

	faceEncoding := Encoding(f.Descriptor)
	return faceEncoding, nil
}

func (r *recognizer) Close() {
	r.recognizer.Close()
}
