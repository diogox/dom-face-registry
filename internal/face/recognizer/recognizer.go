package recognizer

import (
	"github.com/Kagami/go-face"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type FaceEncoding face.Descriptor

type Recognizer interface {
	Recognize(targetFaceEncoding FaceEncoding, peopleIDs []uuid.UUID, allFaceEncondings []FaceEncoding) (uuid.UUID, error)
	EncodeFace(imgBytes []byte) (FaceEncoding, error)
}

type recognizer struct {
	recognizer        *face.Recognizer
	classifyThreshold float32
}

func NewRecognizer(dlibModelsPath string, classifyThreshold float32) (*recognizer, error) {
	rec, err := face.NewRecognizer(dlibModelsPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed at initializing face recognizer")
	}

	return &recognizer{
		recognizer:        rec,
		classifyThreshold: classifyThreshold,
	}, nil
}

func (r *recognizer) Recognize(
	targetFaceEnding FaceEncoding,
	peopleIDs []uuid.UUID,
	allFaceEncodings []FaceEncoding,
) (uuid.UUID, error) {

	// Set Samples
	var samples []face.Descriptor
	for _, fe := range allFaceEncodings {
		samples = append(samples, face.Descriptor(fe))
	}

	categories := make([]int32, 0, len(peopleIDs))
	categoriesLabelsMap := make(map[int]uuid.UUID, len(peopleIDs))
	for i, label := range peopleIDs {
		categoriesLabelsMap[i] = label
		categories = append(categories, int32(i))
	}

	r.recognizer.SetSamples(samples, categories)

	categoryID := r.recognizer.ClassifyThreshold(face.Descriptor(targetFaceEnding), r.classifyThreshold)
	if categoryID < 0 {
		return uuid.Nil, errors.New("unable to classify image")
	}

	return categoriesLabelsMap[categoryID], nil
}

func (r *recognizer) EncodeFace(imgBytes []byte) (FaceEncoding, error) {
	f, err := r.recognizer.RecognizeSingle(imgBytes)
	if err != nil {
		return FaceEncoding{}, errors.Wrap(err, "failed to encode face")
	}

	faceEncoding := FaceEncoding(f.Descriptor)
	return faceEncoding, nil
}

func (r *recognizer) Close() {
	r.recognizer.Close()
}
