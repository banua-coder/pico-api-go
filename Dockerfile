# Multi-stage build — pico-api-go
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build deps
RUN apk add --no-cache git

# Copy dependency files dulu (layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build production binary (no Swagger, smaller size)
RUN CGO_ENABLED=0 GOOS=linux go build -tags=production \
    -ldflags="-s -w" \
    -o bin/server \
    cmd/main_production.go

# ---- Runtime image (minimal) ----
FROM alpine:3.21

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/bin/server .

EXPOSE 8080

CMD ["./server"]
