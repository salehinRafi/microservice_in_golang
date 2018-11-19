package main

import (
	"context"

	// In order to use `$go dep or $go get` in Dockerfile, we need to ensure that we have our code pushed up to Git, so that it can pull in the service we needed.
	// In local, we can `$go get github.com/<github-user>/<project-name>.
	pb "github.com/salehinRafi/microservice_in_golang/3_Docker-compose-and-datastores/user-service/proto/user"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (srv *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}
func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}
func (srv *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	_, err := srv.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	res.Token = "testingabc"
	return nil
}
func (srv *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := srv.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}
func (srv *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}

/*
What happen here?
1. we've created some code to interface our gRPC methods.
2. We just want to be able to create and fetch a user.
*/
