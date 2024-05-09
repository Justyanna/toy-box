FROM golang:bookworm
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN ls
RUN go build -o start ./cmd/app/main.go
EXPOSE 8080
CMD ["/app/start"]
