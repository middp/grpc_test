package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	pb "grpc-demo/proto"
	"log"
	"os"

	"google.golang.org/grpc"
)

var port string

func init() {
	flag.StringVar(&port, "p", "9123", "开放端口号")
	flag.Parse()
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	conn, _ := grpc.Dial("120.79.31.199:"+port, grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewWriteClient(conn)
	fmt.Println("input(enter to stop):")
	content, _ := inputReader.ReadString('\n')
	err := writeSomething(content, client)
	if err != nil {
		log.Printf("write failed")
	}
}

func writeSomething(content string, client pb.WriteClient) error {
	resp, err := client.WriteSomething(context.Background(), &pb.WriteRequest{
		Content: content,
	})
	if err != nil {
		return err
	}
	if resp.Result {
		log.Println("write success")
	}
	return nil
}
