package main

import (
	"context"
	"fmt"
	"github.com/diogox/dom-face-registry/internal/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"net"
	"os"

	"github.com/diogox/dom-face-registry/internal/face"
	"github.com/diogox/dom-face-registry/internal/face/recognizer"
	faceMongo "github.com/diogox/dom-face-registry/internal/face/store/mongo"
	DomFaceRegistry "github.com/diogox/dom-face-registry/internal/pb"
	"github.com/diogox/dom-face-registry/internal/person"
	personMongo "github.com/diogox/dom-face-registry/internal/person/store/mongo"
	"github.com/diogox/dom-face-registry/internal/registry"
	grpcImpl "github.com/diogox/dom-face-registry/internal/transport/grpc"
)

func main() {
	logger := logrus.New()

	cfg, err := getConfig()
	if err != nil {
		logger.Fatal(err)
	}

	if cfg.Server.Debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	if err := run(cfg, logger); err != nil {
		logger.Fatal(err)
	}
}

func getConfig() (config.Config, error) {
	const configFileDefaultPath = "./config.yaml"

	configFilePath, ok := os.LookupEnv(configFileEnv)
	if !ok {
		configFilePath = configFileDefaultPath
	}

	cfg, err := config.Get(configFilePath)
	if err != nil {
		return cfg, errors.Wrap(err, "failed to get config")
	}

	return cfg, nil
}

const configFileEnv = "CONFIG_FILE"

func run(cfg config.Config, logger logrus.FieldLogger) error {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%d", cfg.Database.Host, cfg.Database.Port),
	))
	if err != nil {
		return errors.Wrap(err, "failed to connect to mongo")
	}
	defer client.Disconnect(ctx)

	peopleStore, err := personMongo.NewStore(ctx, client, cfg.Database.DBName, cfg.Database.PeopleCollection)
	if err != nil {
		return errors.Wrap(err, "failed to initialize people store")
	}

	faceStore, err := faceMongo.NewStore(ctx, client, cfg.Database.DBName, cfg.Database.FacesCollection)
	if err != nil {
		return errors.Wrap(err, "failed to initialize face store")
	}

	faceRecognizer, err := recognizer.NewRecognizer(
		cfg.Server.Recognition.DlibModelsDir,
		cfg.Server.Recognition.Threshold,
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize face recognizer")
	}
	defer faceRecognizer.Close()

	regService := registry.NewRegistryService(
		logger,
		person.NewService(peopleStore),
		face.NewService(faceStore, faceRecognizer),
	)

	grpcServer := grpc.NewServer()
	DomFaceRegistry.RegisterFaceRegistryServer(grpcServer, grpcImpl.NewServer(
		logger,
		regService,
		person.NewConverter(),
	))

	logger.Info("listening...")
	if err := grpcServer.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to start the grpc server")
	}

	return nil
}
