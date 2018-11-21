FROM golang:latest as builder

WORKDIR /go/src/microservice_in_golang/3_Docker-compose-and-datastores/user-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build R -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/microservice_in_golang/3_Docker-compose-and-datastores/user-service/user-service .

CMD ["./vessel-service"]