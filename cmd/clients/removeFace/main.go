package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/diogox/dom-face-registry/pkg/client"
	pb "github.com/diogox/dom-face-registry/pkg/registry"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

const endpoint = ":8080"

func run() error {
	faceId := os.Args[1]

	c, conn, err := client.NewClientBuilder().
		WithTarget(endpoint).
		Build()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = c.RemoveFace(context.Background(), &pb.RemoveFaceRequest{
		FaceId: faceId,
	})
	if err != nil {
		return err
	}

	fmt.Println("Successfully Removed!")
	return err
}
