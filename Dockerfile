# Build stage
FROM golang:alpine AS builder

WORKDIR /app

# Copy go.mod
COPY go.mod ./

# Copy the source code
COPY . .

# Build the application
# CGO_ENABLED=0 ensures a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o blog-app .

# Final stage
FROM alpine:latest

# Add CA certificates and tzdata just in case the app makes external requests or uses timezones
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/blog-app .

# Expose the port the app runs on (from main.go)
EXPOSE 8080

# Run the application
CMD ["./blog-app"]
