# Stage 1: Build the Golang application
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy module definition
COPY go.mod ./
# Copy source code
COPY *.go ./

# Compile a statically linked Go binary (no CGO, stripped debug symbols for minimal size)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /mock-api main.go

# Stage 2: Minimal lightweight runtime image
FROM alpine:3.20

# Install ca-certificates and timezone data
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /mock-api /app/mock-api

# Default HTTP port
EXPOSE 8080
ENV PORT=8080

# Run the server
ENTRYPOINT ["/app/mock-api"]
