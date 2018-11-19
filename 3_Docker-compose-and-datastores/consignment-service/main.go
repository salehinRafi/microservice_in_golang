package main

import (
	"fmt"
	"log"
	"os"

	// Import the generated protobuf code and go-micro library.
	"github.com/micro/go-micro"

	// In order to use `$go dep or $go get` in Dockerfile, we need to ensure that we have our code pushed up to Git, so that it can pull in the service we needed.
	// In local, we can `$go get github.com/<github-user>/<project-name>.
	pb "github.com/salehinRafi/microservice_in_golang/3_Docker-compose-and-datastores/consignment-service/proto/consignment"
	vesselProto "github.com/salehinRafi/microservice_in_golang/3_Docker-compose-and-datastores/vessel-service/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	// Database host from the environment variables
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	// Mgo creates a 'master' session, we need to end that session
	// before the main function closes.
	defer session.Close()

	if err != nil {

		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

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
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

/*
What happen here?
We separate the code to make it much more tidy.
*/
