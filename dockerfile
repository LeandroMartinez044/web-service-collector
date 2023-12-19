# Use an official Golang runtime as a base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR ./cmd/api/

# Copy the local package files to the container's workspace
COPY . .

# Install youtube-dl
RUN apt-get update && \
    apt-get install -y youtube-dl && \
    rm -rf /var/lib/apt/lists/*

# Build the application
RUN go build -o web-service-collector ./cmd/api

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./web-service-collector"]