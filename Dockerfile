# Use Go base image
FROM golang:1.25-alpine

# Set the working directory
WORKDIR /app

# Install git for fetching Go modules
RUN apk add --no-cache git

COPY go.mod go.sum /app

# Download dependencies
RUN go mod tidy

# Copy all files into the container
COPY . /app

# Build the Go binary
RUN go build -o app .

# Expose the application port
EXPOSE 8000

# Run the application
CMD ["./app"]
