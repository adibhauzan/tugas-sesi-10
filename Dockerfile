# Stage 1: Build
FROM golang:alpine AS builder

# Set working directory for the build
WORKDIR /app

# Copy module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

# Stage 2: Runtime
FROM debian:bullseye-slim

# Install dependencies
RUN apt-get update && apt-get install -y \
    --no-install-recommends && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Set timezone ke Asia/Jakarta
ENV TZ=Asia/Jakarta

#  timezone lokal
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone

# Set working directory
WORKDIR /app

# Copy binary dari builder stage
COPY --from=builder /app/app /main

# Expose Port
EXPOSE 3291

# Run App
CMD ["/main"]