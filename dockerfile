# Stage 1: Build the application
FROM golang:1.21 AS builder

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

# Install CA certificates and Go
RUN apk --no-cache add ca-certificates && \
    apk --no-cache add --virtual build-dependencies curl && \
    curl -LO https://golang.org/dl/go1.17.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz && \
    rm go1.17.3.linux-amd64.tar.gz && \
    apk del build-dependencies
# Stage 2: Create a minimal image to run the application
FROM alpine:latest

COPY --from=builder /app/web-service-collector .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./web-service-collector"]