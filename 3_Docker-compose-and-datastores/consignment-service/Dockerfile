FROM golang:latest as builder

WORKDIR /go/src/github.com/salehinRafi/microservice_in_golang/3_Docker-compose-and-datastores/consignment-service

COPY . .


RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/salehinRafi/microservice_in_golang/3_Docker-compose-and-datastores/consignment-service/consignment-service .

CMD ["./consignment-service"]


#This is multi-stage build