# Use official Golang image as builder
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o viddler-server ./cmd/viddler-server

FROM builder AS dev
RUN apk add --no-cache yt-dlp
RUN go install github.com/air-verse/air@latest

# Run the binary
CMD ["air"]

# Use minimal alpine image for runtime
FROM alpine:latest AS prod

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates yt-dlp
WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/viddler-server .
COPY --from=builder /app/templates ./templates

CMD ["./viddler-server"]
