# Use an official Go runtime as a parent image
FROM golang:1.21

WORKDIR /app1

COPY ./envoy.yaml envoy.yaml

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Install any needed dependencies
RUN go mod download

# Generate the gRPC code
# RUN protoc -I proto/ proto/myservice.proto --go_out=plugins=grpc:.

# Build the Go application
RUN go build -o mygrpcapp .

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the executable
CMD ["./mygrpcapp"]
