# Use an official Golang runtime as a base image
FROM golang:1.20 AS builder


# Set the working directory inside the container
WORKDIR /web-service-collector


# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o web-service-collector ./cmd/api

# Stage 2: Create a minimal image to run the application
FROM alpine:latest


# Set the working directory inside the container
WORKDIR /web-service-collector

# Copy the built Go application from the builder stage
COPY --from=builder /web-service-collector .

# Install youtube-dl dependencies
RUN apk --no-cache add curl python3 py3-pip 

# Install youtube-dl
RUN curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o /usr/local/bin/youtube-dl && \
    chmod a+rx /usr/local/bin/youtube-dl

# Install AWS CLI using apk package manager
RUN apk --no-cache add aws-cli

# Set AWS CLI default configuration during image build
ARG AWS_ACCESS_KEY_ID
ARG AWS_SECRET_ACCESS_KEY
ARG AWS_REGION

RUN aws configure set aws_access_key_id $AWS_ACCESS_KEY_ID && \
    aws configure set aws_secret_access_key $AWS_SECRET_ACCESS_KEY && \
    aws configure set default.region $AWS_REGION && \
    aws configure set default.output json

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./web-service-collector"]

