# Build stage
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install swag untuk generate swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy source code
COPY . .

# Generate swagger docs
RUN swag init -g cmd/main.go

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for SSL/TLS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"]