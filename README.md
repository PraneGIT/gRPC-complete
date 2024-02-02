# gRPC Complete

This is a simple gRPC project that demonstrates unary,client-server,bidirectional communication between a server and a client using Go programming language,
while attempting a simple implementation of containerization and creating a pod/deployment.

## Project Structure


- `client`: Contains the client-side code.
- `server`: Contains the server-side code.
- `proto`: Contains the Protocol Buffers definition file (`greet.proto`) and generated gRPC code.
- `go.mod` and `go.sum`: Go module files.
- `Dockerfile`: Dockerfile for building Docker image containing both server and client.
- `deployment.yaml`: Kubernetes Deployment file.
- `service.yaml`: Kubernetes Service file.
- `README.md`: Instructions and information about the project.

## Getting Started

### Prerequisites

- Go programming language (https://golang.org/dl/)
- Docker (https://docs.docker.com/get-docker/)
- Kubernetes (https://kubernetes.io/docs/setup/)

### Running Locally

1. Clone this repository:

    ```bash
    git clone https://github.com/PraneGIT/gRPC-complete.git
    ```

2. Navigate to the project directory:

    ```bash
    cd GRPC
    ```

3. Build and run the server:

    ```bash
    cd server
    go run *.go
    ```

4. In another terminal window, build and run the client:

    ```bash
    cd client
    go run *.go
    ```

### Docker Deployment

1. Build the Docker images:

    ```bash
    docker build -t grpc-app .
    ```

2. Tag the Docker image:

    ```bash
    docker tag grpc-app <your-docker-username>/grpc-app:latest
    ```

3. Push the Docker image to Docker Hub:

    ```bash
    docker push <your-docker-username>/grpc-app:latest
    ```

4. Run the Docker container:

    ```bash
    docker run -p 8080:8080 grpc-app
    ```

### Kubernetes Deployment

1. Apply the Kubernetes manifests:

    ```bash
    kubectl apply -f deployment.yaml
    kubectl apply -f service.yaml
    ```

2. Verify the deployment:

    ```bash
    kubectl get pods
    kubectl get services
    ```

