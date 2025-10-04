# Stage 1: Build
# Use the official Go image as the builder
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy dependency management files
COPY go.mod go.sum ./

# Download dependencies
# This will be cached by Docker if go.mod/go.sum do not change
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application
# -o /app/main will create a binary named 'main' in the /app directory
# CGO_ENABLED=0 and -ldflags="-s -w" to produce a smaller static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main -ldflags="-s -w" ./main.go

# Stage 2: Final Image
# Use a minimal base image, since we only need the compiled binary
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the build result (the 'main' binary) from the builder stage
COPY --from=builder /app/main .

# Expose the port the app will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]