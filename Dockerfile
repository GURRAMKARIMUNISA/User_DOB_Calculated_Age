# Start with a Go base image that is Go 1.24 or newer
FROM golang:1.24-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the entire project directory into the container
# This is crucial for Go to recognize your internal modules
COPY . .

# Download Go module dependencies (external ones)
# This will now correctly resolve internal paths because the source is present
RUN go mod download

# Build the application
# CGO_ENABLED=0 disables CGO, useful for creating statically linked binaries
# -a ensures all packages are rebuilt
# -installsuffix nocgo prevents conflict with CGO-enabled builds
RUN CGO_ENABLED=0 go build -o /go-user-api ./cmd/server

# Use a minimal base image for the final stage
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built executable from the builder stage
COPY --from=builder /go-user-api .

# Copy the .env file (if you want to include it in the image, otherwise rely on env vars)
# COPY .env .

# Expose the port the app runs on
EXPOSE 3000

# Command to run the executable
CMD ["./go-user-api"]
