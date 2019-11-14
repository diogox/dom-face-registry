package main

import (
	"context"
	"fmt"
	"github.com/diogox/dom-face-registry/pkg/client"
	pb "github.com/diogox/dom-face-registry/pkg/registry"
	"io/ioutil"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

const endpoint = ":8080"
const imgPath = "data/photo1.jpg"

func run() error {
	c, conn, err := client.NewClientBuilder().
		WithTarget(endpoint).
		Build()
	if err != nil {
		return err
	}
	defer conn.Close()

	stream, err := c.RecognizeFace(context.Background())
	if err != nil {
		return err
	}

	// Open img
	imgData, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return err
	}

	err = client.UploadImageInChunks(imgData, func(chunk []byte) error {
		return stream.Send(&pb.FaceImage{
			ImageData: chunk,
		})
	})
	if err != nil {
		return err
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	fmt.Println(res.PersonInfo)
	return nil
}
