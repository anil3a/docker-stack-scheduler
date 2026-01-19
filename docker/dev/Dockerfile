FROM golang:1.25-alpine

# Minimal tools
RUN apk add --no-cache git curl ca-certificates build-base bash

# Install air for hot reload
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Expose port for dev server
EXPOSE 8080

# Default command
CMD ["bash"]
