package client

import (
	"errors"
	pb "github.com/diogox/dom-face-recognizer/pkg/registry"
	"google.golang.org/grpc"
)

type builder struct {
	target string
}

func NewClientBuilder() *builder {
	return &builder{}
}

func (b *builder) WithTarget(target string) *builder {
	b.target = target
	return b
}

// Build returns the grpc clients, the connection (which you should `defer conn.Close`) and any errors that might have occurred
func (b *builder) Build() (pb.FaceRegistryClient, *grpc.ClientConn, error) {
	if b.target == "" {
		return nil, nil, errors.New("missing target")
	}

	conn, err := grpc.Dial(b.target, grpc.WithInsecure()) // TODO: Should I provide a method for the "secure" version?
	if err != nil {
		return nil, nil, err
	}

	return pb.NewFaceRegistryClient(conn), conn, nil
}