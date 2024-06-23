FROM golang:1.21-bookworm

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-rest-api ./cmd/main

EXPOSE 8080

CMD ["/docker-rest-api"]