package main

import (
    "context"
    "fmt"
    "log"
    "time"

    pb "jit-grpc-starter/proto"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    c := pb.NewAccessServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    req := &pb.AccessRequest{
        UserId:          "user1",
        Role:            "admin",
        DurationMinutes: 30,
        Justification:   "Need temporary access",
    }

    res, err := c.RequestAccess(ctx, req)
    if err != nil {
        log.Fatalf("could not request access: %v", err)
    }

    fmt.Printf("Request ID: %s, Status: %s\n", res.RequestId, res.Status)
}
