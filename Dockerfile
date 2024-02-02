FROM golang:latest AS builder

WORKDIR /app


COPY go.mod .
COPY go.sum .
RUN go mod download


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


COPY --from=builder /app/server/main ./server/
COPY --from=builder /app/client/main ./client/


EXPOSE 8080

# Command to run the server and client
CMD ./server/main & ./client/main 
