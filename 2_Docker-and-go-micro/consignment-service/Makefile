build:
		protoc -I. --go_out=plugins=micro:$(GOPATH)\src\microservice_in_golang\2_Docker-and-go-micro\consignment-service \
			proto\consignment\consignment.proto

		docker build -t consignment-service .

run:
		docker run -p 50051:50051 \
			-e MICRO_SERVER_ADDRESS=:50051 \
        		-e MICRO_REGISTRY=mdns consignment-service

# 1. Call the protoc library, which is responsible for compiling our protobuf definition into code. Also specify the use of the go-micro plugin, as well as the build context and the output path.

# 2. Building Go Binary
#	i. GOOS and GOARCH allow you to cross-compile your go binary for another operating system.
#	ii. docker build process - read our Dockerfile & build an image by the name `consignment-service`.

# 3. Run docker image.
#	i. exposing the port 50051 - Because Docker runs on a separate networking layer, you need to forward the port used within your Docker container, to your host. You can forward the internal port to a new port on the host by changing the first segment.
#	ii. -e is envinronment flag. This allow us to pass in envinronement variables into your Docker container