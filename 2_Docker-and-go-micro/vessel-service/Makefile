build:
		protoc -I. --go_out=plugins=micro:$(GOPATH)\src\microservice_in_golang\2_Docker-and-go-micro\vessel-service \
    		proto\vessel\vessel.proto
		#docker build -t vessel-service .

# We can't run two docker containers on the same port, so we make use of Dockers port forwarding here to ensure this service forwards 50051 to 50052 on the host network.

run:
		docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service