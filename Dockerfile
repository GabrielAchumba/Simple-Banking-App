# Use the official Golang image as a build stage
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
Run go build -o /bankingapp main.go

# Use a minimal image for the final stage
FROM alpine:latest

# COPY the binary from the builder stage
COPY --from=builder /bankingapp /bankingapp

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the binary
ENTRYPOINT ["/bankingapp"]