package gen

//go:generate mkdir -p ../pkg/registry/
//go:generate protoc -I ../api/proto registry.proto face_registry.proto person_registry.proto types.proto --go_out=plugins=grpc:../pkg/registry

//go:generate mkdir -p ../internal/pb/
//go:generate protoc -I ../api/proto registry.proto face_registry.proto person_registry.proto types.proto --go_out=plugins=grpc:../internal/pb
