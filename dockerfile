# Use an official Golang runtime as a base image
FROM golang:1.16 AS builder

# Set the working directory inside the container
WORKDIR /go/src/web-service-collector

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download and install any required dependencies
RUN go get -d -v ./...

# Install the application
RUN go install -v ./...

# Use a smaller base image for the final image
FROM alpine:latest

# Install necessary dependencies
RUN apk --no-cache add \
    ca-certificates \
    ffmpeg \
    python3

# Install pip
RUN apk add --update py3-pip

# Upgrade pip
RUN pip3 install --upgrade pip

# Install youtube-dl using pip
RUN pip3 install --upgrade youtube-dl

# Copy the built executable from the previous stage
COPY --from=builder /go/bin/web-service-collector /usr/local/bin/web-service-collector

# Set the PORT environment variable
ENV PORT 8080

# Expose the application's port
EXPOSE 8080

# Command to run the executable
CMD ["./web-service-collector"]




