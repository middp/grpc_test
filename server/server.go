package main

import (
	"context"
	"flag"
	"fmt"
	pb "grpc-demo/proto"
	"net"
	"os"

	"google.golang.org/grpc"
)

type Server struct{}

var port string

func init() {
	flag.StringVar(&port, "p", "9123", "开放端口号")
	flag.Parse()
}

func (s *Server) WriteSomething(ctx context.Context, w *pb.WriteRequest) (*pb.WriteResult, error) {
	// path := "/home/footman/goproj/go-programming-tour-book/grpc-demo/a.txt"
	path := "/root/test.txt"
	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	num, err := f.WriteString(w.Content)
	fmt.Printf("write %d bytes to file\n", num)
	if err != nil {
		return &pb.WriteResult{Result: false}, nil
	}
	return &pb.WriteResult{Result: true}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterWriteServer(server, &Server{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}
