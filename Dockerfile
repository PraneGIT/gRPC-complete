# Use a multi-stage build
FROM golang:latest AS builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the entire project
COPY . .

# Build the server
WORKDIR /app/server
RUN go build -o /app/server/main

# Build the client
WORKDIR /app/client
RUN go build -o /app/client/main

# Final stage: create a lightweight container for running the server and client
FROM golang:latest

WORKDIR /app

# Copy the built binaries from the previous stage
COPY --from=builder /app/server/main ./server/
COPY --from=builder /app/client/main ./client/

# Expose necessary ports if needed
EXPOSE 8080

# Command to run the server and client
CMD ["./server/main"]  
