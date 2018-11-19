build:
		protoc -I. --go_out=plugins=grpc:$(GOPATH)\src\microservice_in_golang\1_Basics-of-writing-a-gRPC-based-microservice\consignment-service \
			proto\consignment\consignment.proto

# 1. Call the protoc library, which is responsible for compiling our protobuf definition into code. Also specify the use of the grpc plugin, as well as the build context and the output path.