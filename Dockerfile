# Step 1: Build the Go binary
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git for Go module fetching
RUN apk add --no-cache git

# Copy dependency files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the code
COPY . .

# Build the Go binary
RUN go build -o app .

# Step 2: Create a minimal runtime image
FROM alpine:3.20

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .

# Expose port
EXPOSE 8000

# Run the binary
CMD ["./app"]
