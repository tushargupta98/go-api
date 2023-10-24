# Use a smaller base image
FROM golang:1.21-alpine as builder

WORKDIR /app

# Copy Go module and Go sum files for dependencies
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the rest of your application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o api

# Use a minimal base image for the final image
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder image
COPY --from=builder /app/api .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./api"]
