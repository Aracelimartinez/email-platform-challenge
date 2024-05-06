# Development stage
FROM golang:1.22 AS development

#Create the /app directory within the image to hold the app code
WORKDIR /app

# Install live reloading tool
RUN go install github.com/cosmtrek/air@latest
COPY .air.toml .

# Copy all code
COPY . .

EXPOSE 3000

# Command to start the development environment with live reloading
CMD ["air", "-c", ".air.toml"]


# Builder stage
FROM golang:1.22 AS builder

#Create the /app directory within the image to hold the app code
WORKDIR /app

# Set environment variables for building the binary
ENV CGO_ENABLED=0 GOOS=linux

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download  && go mod verify

# Copy app files
COPY . .

# Build the application's binary
RUN go build -v -o api .


# Production stage
FROM alpine:latest AS production

WORKDIR /root/

# Install ca-certificates for HTTPS requests.
RUN apk add --no-cache ca-certificates

# Copy the binary from the builder stage.
COPY --from=builder /app/api .

# Command to run the api binary
CMD ["./api"]
