FROM golang:1.25-alpine

# Install minimal tools
RUN apk add --no-cache \
    git \
    curl \
    ca-certificates \
    build-base

# Install air (hot reload)
RUN go install github.com/air-verse/air@latest

WORKDIR /src

# Pre-copy mod files for cache
COPY go.mod go.sum ./
RUN go mod download

# Copy source
# COPY . .

EXPOSE 8080

CMD ["air", "-c", "/src/.air.toml"]
