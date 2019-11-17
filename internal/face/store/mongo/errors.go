package mongo

const (
	errMongoFailedToConnect = "could not connect to mongo"

	errMongoFailedToGetFaces   = "could not get faces from mongo"
	errMongoFailedToDecodeFace = "could not decode face from mongo"

	errMongoFailedToInsertFace = "could not insert face in mongo"

	errMongoFailedToRemoveFace =  "could not remove face in mongo"

	errMongoInvalidPersonID       = "person id is nil"
	errMongoInvalidInsertedFaceID = "invalid inserted face id type"
)
