package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// Import the generated protobuf code.
	pb "microservice_in_golang/1_Basics-of-writing-a-gRPC-based-microservice/consignment-service/proto/consignment"
)

const (
	port = ":50051"
)

type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// ConsignmentRepository - Dummy repository, this simulates the use of a datastore of some kind. We'll replace this with a real implementation later on.
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// service - Service should implement all of the methods to satisfy the service we defined in our protobuf definition. You can check the interface in the generated code itself (consignment.pb.go) for the exact method signatures etc to give you a better idea.
type service struct {
	repo Repository
}

// CreateConsignment - we created just one method on our service, which is a create method, which takes a context and a request as an argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	// Save our consignment.
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Return macthing the `Response` message we created in our protobuf definition.
	return &pb.Response{
		Created:     true,
		Consignment: consignment,
	}, nil
}

// GetConsignment - view all of our created consignments.
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	repo := &ConsignmentRepository{}

	// Set up our gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our implementation into the auto-generated interface code for our protobuf definition.
	pb.RegisterShippingServiceServer(s, &service{repo})

	// Register reflection service on gRPC server
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

/*
 Here we are creating the implementation logic which our gRPC methods interface with, using the generated formats, creating a new gRPC server on port 50051.
*/
