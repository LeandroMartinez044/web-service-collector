# Stage 1: Build the application
FROM golang:1.17 AS builder

# Set the working directory inside the container
WORKDIR ./cmd/api/


COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Stage 2: Create a minimal image to run the application
FROM alpine:latest


# Install youtube-dl
RUN apk --no-cache add youtube-dl

# Add youtube-dl to the system's PATH
ENV PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/root/bin"

# Build the application
RUN go build -o app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./web-service-collector"]