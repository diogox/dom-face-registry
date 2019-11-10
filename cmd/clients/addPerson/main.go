package main

import (
	"context"
	"fmt"
	"log"

	"github.com/diogox/dom-face-recognizer/pkg/client"
	pb "github.com/diogox/dom-face-recognizer/pkg/registry"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

const endpoint = ":8080"

func run() error {
	c, conn, err := client.NewClientBuilder().
		WithTarget(endpoint).
		Build()
	if err != nil {
		return err
	}
	defer conn.Close()

	res, err := c.AddPerson(context.Background(), &pb.AddPersonRequest{
		FirstName: "Nicholas",
		LastName:  "Cage",
		Roles:     nil,
	})
	if err != nil {
		return err
	}

	fmt.Println(res)
	return nil
}
