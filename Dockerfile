FROM golang:1.23.4-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main cmd/api/main.go
EXPOSE 8080
CMD ["./main"]