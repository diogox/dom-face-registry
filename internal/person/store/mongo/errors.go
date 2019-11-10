package mongo

const (
	errMongoFailedToConnect = "could not connect to mongo"

	errMongoFailedToGetPeople = "could not get people from mongo"
	errMongoFailedToDecodePerson = "could not decode person from mongo"

	errMongoFailedToInsertPerson = "could not insert person in mongo"

	errMongoFailedToDeletePerson = "could not remove person in mongo"

	errMongoInvalidID = "invalid id, could not convert"
)
