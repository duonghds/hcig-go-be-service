# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod verify

COPY . .
RUN go get  -t -v ./...

RUN go build -o ./go-be-service ./cmd
EXPOSE 2411
CMD ["./go-be-service"]