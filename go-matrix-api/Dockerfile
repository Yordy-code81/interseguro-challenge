# Build stage
FROM golang:1.26.4-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source code
COPY . .

# Generate Swagger documentation
RUN go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/api/main.go

# Run tests to ensure code quality before building
RUN go test ./... -v

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Run stage
FROM alpine:latest  

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
