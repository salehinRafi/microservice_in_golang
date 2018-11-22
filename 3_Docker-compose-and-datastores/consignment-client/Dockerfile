FROM golang:latest as builder

WORKDIR /go/src/github.com/salehinRafi/microservice_in_golang/3_Docker-compose-and-datastores/consignment-client

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/github.com/salehinRafi/microservice_in_golang/3_Docker-compose-and-datastores/consignment-client/consignment-client .

CMD ["./consignment-client"]