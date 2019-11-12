FROM docker.io/ndphu/ubuntu-go-dlib

WORKDIR /go/src/github.com/diogox/dom-face-recognizer

COPY . .

ARG DLIB_MODELS_DIRECTORY=/go/src/github.com/diogox/dom-face-recognizer/data/models

RUN mkdir -p $DLIB_MODELS_DIRECTORY

RUN wget -P $DLIB_MODELS_DIRECTORY http://dlib.net/files/shape_predictor_5_face_landmarks.dat.bz2 \
    && bzip2 -d $DLIB_MODELS_DIRECTORY/shape_predictor_5_face_landmarks.dat.bz2

RUN wget -P $DLIB_MODELS_DIRECTORY http://dlib.net/files/dlib_face_recognition_resnet_model_v1.dat.bz2 \
    && bzip2 -d $DLIB_MODELS_DIRECTORY/dlib_face_recognition_resnet_model_v1.dat.bz2

RUN wget -P $DLIB_MODELS_DIRECTORY http://dlib.net/files/mmod_human_face_detector.dat.bz2 \
    && bzip2 -d $DLIB_MODELS_DIRECTORY/mmod_human_face_detector.dat.bz2

CMD ["go", "run", "cmd/server/main.go"]
