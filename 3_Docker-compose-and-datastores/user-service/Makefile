build:
		protoc -I. --go_out=plugins=micro:$(GOPATH)\src\microservice_in_golang\3_Docker-compose-and-datastores\user-service \
			proto\user\user.proto
		docker build -t user-service .

run:
	docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns user-service