package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/micro/go-micro/cmd"

	// In order to use `$go dep or $go get` in Dockerfile, we need to ensure that we have our code pushed up to Git, so that it can pull in the service we needed.
	// In local, we can `$go get github.com/<github-user>/<project-name>
	pb "github.com/salehinRafi/microservice_in_golang/3_Docker-compose-and-datastores/consignment-service/proto/consignment"

	microclient "github.com/micro/go-micro/client"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {

	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err

}

func main() {

	cmd.Init()

	// Create new greeter client
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	resp, err := client.CreateConsignment(context.TODO(), consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", resp.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}

/*
What happen here?

1. This is a client to see consignment-service action
2. Here we create a command line interface, which will take a JSON consignment file and interact with our gRPC service.
3. Here we've imported the go-micro libraries for creating clients, and replaced our existing connection code, with the go-micro client code, which uses service resolution instead of connecting directly to an address.
*/
