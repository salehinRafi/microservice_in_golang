FROM golang:latest as builder

WORKDIR /go/src/microservice_in_golang/2_Docker-and-go-micro/consignment-client

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/microservice_in_golang/2_Docker-and-go-micro/consignment-client/consignment-client .

CMD ["./consignment-client"]