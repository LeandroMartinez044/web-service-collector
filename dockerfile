# Use an official Golang runtime as a base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR ./cmd/api/


# Copy go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download and install dependencies
RUN go mod download

# Install youtube-dl
RUN apt-get clean && \
    apt-get update && \
    apt-get install -y --no-install-recommends apt-utils && \
    apt-get install -y youtube-dl && \
    rm -rf /var/lib/apt/lists/*

# Set the PATH to include the directory where youtube-dl is installed
ENV PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/root/bin"

# Copy the local package files to the container's workspace
COPY . .

# Build the application
RUN go build -o web-service-collector ./cmd/api

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./web-service-collector"]