# Use an official Golang runtime as a base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR ./cmd/api/

# Copy the local package files to the container's workspace
COPY . .

# Build the application
RUN go build -o web-service-collector ./cmd/api

# Use a smaller Alpine Linux image for the final image
FROM alpine:latest

# Install necessary dependencies (including Python and youtube-dl)
RUN apk --no-cache add \
    ca-certificates \
    ffmpeg \
    python3 \
    py3-pip

# Upgrade pip
RUN pip3 install --upgrade pip

# Install youtube-dl using pip
RUN pip3 install --upgrade youtube-dl

# Copy the built Golang executable from the previous stage
COPY --from=0 /go/src/app/cmd/api/web-service-collector /usr/local/bin/web-service-collector

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./web-service-collector"]