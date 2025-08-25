package main

import (
	"log"
	"net"

	itemstore "gserver/proto"
	pb "gserver/proto" // Import the generated protobuf code

	"google.golang.org/grpc"
)

type server struct {
	// pb.UnimplementedGreeterServer is embedded for forward compatibility.
	itemstore.UnimplementedItemServiceServer
}

func (*server) InsertItem(stream itemstore.ItemService_InsertItemServer) error {

}

func ServerStart(addr string) (*grpc.Server, error) {

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("error starting service")
	}

	s := grpc.NewServer()
	// Register our server implementation with the gRPC server.
	pb.RegisterItemServiceServer(s, &server{})

	go func() {
		log.Printf("gRPC server starting to listen on %s", addr)
		if err := s.Serve(listener); err != nil {
			// This error is expected when GracefulStop() is called.
			if err == grpc.ErrServerStopped {
				log.Printf("gRPC server has been stopped.")
			} else {
				log.Fatalf("gRPC server failed to serve: %v", err)
			}
		}
	}()

	return s, err

}
