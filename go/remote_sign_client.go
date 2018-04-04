package main

import (
	"flag"
//	"io"
	"log"
//	"math/rand"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials"

	pb "./remote_sign"
//	pb "google.golang.org/grpc/examples/route_guide/routeguide"
//	"google.golang.org/grpc/testdata"
)


var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:50051", "The server address in the format of host:port")
//	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func main() {
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewBrowserSignClient(conn)

	hash := []byte("asd")
	data := &pb.SignMessageRequest{Msg: hash}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message, err := client.SignMessage( ctx, data )
	if err != nil {
		log.Fatalf("fail to call: %v", err)
	}

	log.Printf("Message received: %v", message)
}


