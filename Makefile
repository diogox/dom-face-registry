.PHONY: lint

lint: ensure-linter-exists
	golangci-lint run

LINTER_VERSION=v1.18.0
ensure-linter-exists:
	command -v golangci-lint || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ${GOPATH}/bin $(LINTER_VERSION)

DLIB_MODELS_DIRECTORY=./data/models
download-models:
	wget -P ${DLIB_MODELS_DIRECTORY} http://dlib.net/files/shape_predictor_5_face_landmarks.dat.bz2 \
 	    && bzip2 -d ${DLIB_MODELS_DIRECTORY}/shape_predictor_5_face_landmarks.dat.bz2 &&\
	wget -P ${DLIB_MODELS_DIRECTORY} http://dlib.net/files/dlib_face_recognition_resnet_model_v1.dat.bz2 \
    	    && bzip2 -d ${DLIB_MODELS_DIRECTORY}/dlib_face_recognition_resnet_model_v1.dat.bz2 &&\
	wget -P ${DLIB_MODELS_DIRECTORY} http://dlib.net/files/mmod_human_face_detector.dat.bz2 \
    	    && bzip2 -d ${DLIB_MODELS_DIRECTORY}/mmod_human_face_detector.dat.bz2
