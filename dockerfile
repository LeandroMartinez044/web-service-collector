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

# Install necessary dependencies (including Python, youtube-dl, and other dependencies)
RUN apk --no-cache add \
    ca-certificates \
    ffmpeg \
    python3 \
    py3-pip \
    build-base \
    libffi-dev \
    openssl-dev

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
COPY --from=builder ./cmd/api/web-service-collector /usr/local/bin/web-service-collector

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./web-service-collector"]