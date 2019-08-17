package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/vinusstar/grpctest"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8888"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	stream, err := pb.NewUppercaseServiceClient(conn).Transform(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go receive(done, stream)

	data := []string{"tokyo", "001", "Japan"}

	for i := 0; i < 10; i++ {
		err = send(data, stream)
		if err != nil {
			log.Fatalf("send error : %v", err)
		}

		log.Println("wait")
		time.Sleep(1 * time.Second)

	}

	err = stream.CloseSend()
	if err != nil {
		log.Fatalln("close error ", err)
	}

	<-done
}

func send(data []string, stream pb.UppercaseService_TransformClient) error {
	for _, v := range data {
		log.Println("send message : ", v)
		err := stream.Send(&pb.UppercaseRequest{
			Message: v,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func receive(done chan struct{}, stream pb.UppercaseService_TransformClient) {
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			close(done)
			return
		}

		if err != nil {
			log.Fatalf("receive error : %v", err)
		}

		log.Println("receive message : ", res.Message)
	}
}
