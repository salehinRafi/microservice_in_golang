package main

import (
	"context"
	"log"

	"gopkg.in/mgo.v2"

	// In order to use `$go dep or $go get` in Dockerfile, we need to ensure that we have our code pushed up to Git, so that it can pull in the service we needed.
	// In local, we can `$go get github.com/<github-user>/<project-name>.
	pb "github.com/salehinRafi/microservice_in_golang/3_Docker-compose-and-datastores/consignment-service/proto/consignment"
	vesselProto "github.com/salehinRafi/microservice_in_golang/3_Docker-compose-and-datastores/vessel-service/proto/vessel"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	session      *mgo.Session
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	defer s.GetRepo().Close()
	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	err = s.GetRepo().Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	defer s.GetRepo().Close()
	consignments, err := s.GetRepo().GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
