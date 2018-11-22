package main

import (
	"context"
	"fmt"
	"log"

	// Import the generated protobuf code and go-micro library.
	vesselProto "github.com/salehinRafi/microservice_in_golang/2_Docker-and-go-micro/vessel-service/proto/vessel"

	pb "github.com/salehinRafi/microservice_in_golang/2_Docker-and-go-micro/consignment-service/proto/consignment"

	"github.com/micro/go-micro"
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
	repo         Repository
	vesselClient vesselProto.VesselServiceClient
}

// CreateConsignment - we created just one method on our service, which is a create method, which takes a context and a request as an argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	// Here we call a client instance of our vessel service with our consignment weight, and the amount of containers as the capacity value.
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}
	// We set the VesselId as the vessel we got back from our vessel service
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment.
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return macthing the `Response` message we created in our protobuf definition.
	res.Created = true
	res.Consignment = consignment
	return nil
}

// GetConsignment - view all of our created consignments.
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &ConsignmentRepository{}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	// Init will parse the command line flags.
	srv.Init()

	// Register our service with the gRPC server using go-micro service handler, this will tie our implementation into the auto-generated interface code for our protobuf definition.
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

/*
What happen here?
1. Main changes on how we instantiate our gRPC server.
	i. We instantiate our gRPC server, which has been abstracted neatly behind a micro.NewService() method, which handles registering our service.
	ii. The srv.Run() function, which handles the connection itself.

2. Second changes on service methods
	i. the arguments and response types have changes slightly to take both the request and the response struct as arguments, and now only returning an error. Within our methods, we set the response, which is handled by go-micro.

3. No longger hard-coding port.
	i. Go-micro should be configured using environment variables, or command line arguments.
	ii. To set the address, use MICRO_SERVER_ADDRESS=:50051.
	iii. We also need to tell our service to use mdns (multicast dns) as our service broker for local use. You wouldn't typically use mdns for service discovery in production, but we want to avoid having to run something like Consul or etcd locally for the sakes of testing.

4. Created a client instance for our vessel service.
	i. this allows us to use the service name, i.e go.micro.srv.vessel to call the vessel service as a client and interact with its methods. In this case, just the one method (FindAvailable).
	ii. We send our consignment weight, along with the amount of containers we want to ship as a specification to the vessel-service. Which then returns an appropriate vessel.
*/
