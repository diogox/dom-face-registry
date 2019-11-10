package main

import (
	"context"
	"fmt"
	"github.com/diogox/dom-face-recognizer/pkg/client"
	pb "github.com/diogox/dom-face-recognizer/pkg/registry"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

const endpoint = ":8080"
const imgPath = "data/test1.jpg"

func run() error {
	personId := os.Args[1]

	c, conn, err := client.NewClientBuilder().
		WithTarget(endpoint).
		Build()
	if err != nil {
		return err
	}
	defer conn.Close()

	stream, err := c.AddFace(context.Background())
	if err != nil {
		return err
	}

	// Send person id
	err = stream.Send(&pb.AddFaceRequest{
		PersonId: personId,
	})
	if err != nil {
		return err
	}

	// Open img
	imgData, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return err
	}

	err = client.UploadImageInChunks(imgData, func(chunk []byte) error {
		return stream.Send(&pb.AddFaceRequest{
			FaceImage: &pb.FaceImage{
				ImageData: chunk,
			},
		})
	})
	if err != nil {
		return err
	}

	res, err := stream.CloseAndRecv()
	fmt.Println(res)

	return err
}
