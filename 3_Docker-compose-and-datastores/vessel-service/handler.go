package main

import (
	"context"

	// In order to use `$go dep or $go get` in Dockerfile, we need to ensure that we have our code pushed up to Git, so that it can pull in the service we needed.
	// In local, we can `$go get github.com/<github-user>/<project-name>.
	pb "github.com/salehinRafi/microservice_in_golang/3_Docker-compose-and-datastores/vessel-service/proto/vessel"

	"gopkg.in/mgo.v2"
)

// Our gRPC service handler
type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	defer s.GetRepo().Close()
	// Find the next available vessel
	vessel, err := s.GetRepo().FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	defer s.GetRepo().Close()
	if err := s.GetRepo().Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Create = true
	return nil
}
