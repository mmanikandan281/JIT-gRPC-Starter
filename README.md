# JIT gRPC Starter

A simple gRPC starter project in Go that demonstrates a basic access request service using Protocol Buffers.

## Description

This project provides a client-server architecture using gRPC for requesting temporary access. The server implements an `AccessService` that handles access requests, and the client can send requests to the server.

The service includes:
- **AccessService**: A gRPC service with a `RequestAccess` RPC method.
- **AccessRequest**: Message containing user ID, role, duration in minutes, and justification.
- **AccessResponse**: Message containing request ID and status.

## Prerequisites

- Go 1.25.0 or later
- Protocol Buffers compiler (protoc) if you need to regenerate the proto files

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd jit-grpc-starter
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Building

Build the server:
```bash
go build -o server.exe ./server
```

Build the client:
```bash
go build -o client.exe ./client
```

## Running

1. Start the server:
   ```bash
   ./server.exe
   ```
   The server will start listening on port 50051.

2. In a separate terminal, run the client:
   ```bash
   ./client.exe
   ```
   The client will connect to the server, send a sample access request, and print the response.

## Usage

### Server
The server logs incoming requests and responds with a fixed approval status. In a real-world scenario, you would implement business logic for access approval.

### Client
The client sends a sample request with:
- User ID: "user1"
- Role: "admin"
- Duration: 30 minutes
- Justification: "Need temporary access"

You can modify the client code to send different requests or integrate it into your application.

## Project Structure

- `client/`: Contains the gRPC client implementation
- `server/`: Contains the gRPC server implementation
- `proto/`: Contains the Protocol Buffer definitions and generated Go code
- `go.mod`: Go module file
- `go.sum`: Go module checksums

## Dependencies

- [gRPC](https://grpc.io/): High-performance RPC framework
- [Protocol Buffers](https://developers.google.com/protocol-buffers): Language-neutral data serialization

### Done By Manikandan M
