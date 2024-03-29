package main

import (
	"context"
	"fmt"
	"log"

	"github.com/diogox/dom-face-registry/pkg/client"
	pb "github.com/diogox/dom-face-registry/pkg/registry"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c, conn, err := client.NewClientBuilder().
		WithTarget(":8080").
		Build()
	if err != nil {
		return err
	}
	defer conn.Close()

	res, err := c.GetPeople(context.Background(), &pb.GetPeopleRequest{})
	if err != nil {
		return err
	}

	fmt.Println(res)
	return nil
}
