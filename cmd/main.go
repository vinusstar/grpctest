package main

import (
	pb "github.com/vinusstar/grpctest"
    "google.golang.org/grpc"
)

const (
	port = ":8888"
)

type server strunct{}


func (s *server) Transform(stream pb. UppercaseService_TransformServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF{
			return nil
		}

		if err !=nil{
			return err
		}

		resp := &pb.UppercaseResponse {
			Message: strings.ToUpperCase(in.Message),

		}

		err = stream.Send(resp)
		if err != nil{
			return err
		}
	}
}



func main(){

}