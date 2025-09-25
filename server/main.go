package main

import (
    "context"
    "fmt"
    "log"
    "net"

    pb "jit-grpc-starter/proto"
 // Import the generated proto package
    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedAccessServiceServer
}

func (s *server) RequestAccess(ctx context.Context, req *pb.AccessRequest) (*pb.AccessResponse, error) {
    fmt.Printf("Received request from user: %s\n", req.UserId)
    return &pb.AccessResponse{
        RequestId: "REQ123",
        Status:    "APPROVED",
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterAccessServiceServer(s, &server{})

    fmt.Println("Server running at :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
