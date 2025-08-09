FROM golang:1.23-alpine AS builder

# Optional: pass a changing value to force rebuilds when desired
ARG REBUILD=""

# Tooling for CGO builds
RUN apk add --no-cache build-base git

WORKDIR /app

# Copy go mod and sum files first for better layer caching
COPY go.mod go.sum ./

# Download dependencies with module cache (requires BuildKit)
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy the rest of the source
COPY . .

# Echo the arg to break cache when REBUILD changes
RUN echo "Rebuild token: ${REBUILD}"

# Build the application with build cache (requires BuildKit)
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"] 