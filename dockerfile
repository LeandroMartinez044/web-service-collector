# Use an official Golang runtime as a base image
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /go/src/app/cmd/api/

# Copy the local package files to the container's workspace
COPY . .

# Build the Golang application
RUN go build -o web-service-collector

# Use a smaller Alpine Linux image for the final image
FROM alpine:latest

# Install necessary dependencies for youtube-dl
RUN apk --no-cache add \
    ca-certificates \
    ffmpeg \
    python3 \
    py3-pip

# Update the package index
RUN apk update

# Check Python version
RUN python3 --version

# Install youtube-dl using pip within a virtual environment
RUN python3 -m venv /venv && \
    source /venv/bin/activate && \
    pip install --upgrade youtube-dl && \
    deactivate

# Copy the built Golang executable from the previous stage
COPY --from=builder /go/src/app/cmd/api/web-service-collector /usr/local/bin/web-service-collector

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the Golang executable
CMD ["web-service-collector"]
