# Stage 1: Build the application
FROM golang:1.17 AS builder

# Set the working directory inside the container
WORKDIR ./cmd/api/


COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Stage 2: Create a minimal image to run the application
FROM alpine:latest

WORKDIR /app


COPY --from=builder /app/web-service-collector .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./web-service-collector"]