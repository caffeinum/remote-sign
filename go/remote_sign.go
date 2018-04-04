package sign

import (
	"flag"
	"io"
	"log"
	"math/rand"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
//	pb "google.golang.org/grpc/examples/route_guide/routeguide"
//	"google.golang.org/grpc/testdata"
)


var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:50051", "The server address in the format of host:port")
//	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func sign(msg []byte, ) {
}
