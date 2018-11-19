package main

import (
	"fmt"
	"log"
	"os"

	// Import the generated protobuf code and go-micro library.
	// In order to use `$go dep or $go get` in Dockerfile, we need to ensure that we have our code pushed up to Git, so that it can pull in the service we needed.
	// In local, we can `$go get github.com/<github-user>/<project-name>.
	pb "github.com/salehinRafi/microservice_in_golang/3_Docker-compose-and-datastores/vessel-service/proto/vessel"

	"github.com/micro/go-micro"
)

const (
	defaultHost = "localhost:27017"
)

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}
	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Fatalf("Error connecting to datastore: %v", err)
	}

	repo := &VesselRepository{session.Copy()}

	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

/*
What happen here?
We have a consignment service, this will deal with matching a consignment of containers to a vessel which is best suited to that consignment. In order to match our consignment, we need to send the weight and amount of containers to our new vessel service, which will then find a vessel capable of handling that consignment.
*/
