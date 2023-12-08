FROM golang:1.21

WORKDIR /app
ADD . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /movieapp cmd/main.go

EXPOSE 8080

CMD ["/movieapp"]
