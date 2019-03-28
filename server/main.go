// Binary server is a LibraryService server.
package main

import (
	"flag"
	"net"

	"github.com/ghasemloo/libraryservice/server/library"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	lspb "github.com/ghasemloo/libraryservice/proto"
)

var (
	addr = flag.String("addr", ":9000", "server port to listen on")
)

func main() {
	flag.Parse()
	glog.Infof("Starting LibraryService server at %v", *addr)
	ctx := context.Background()

	// Create a Library.
	l, err := library.New(ctx)
	if err != nil {
		glog.Exitf("Failed to create a Library: %v", err)
	}

	// Create a gRPC server.
	s := grpc.NewServer()

	// Register Library on the gRPC server.
	lspb.RegisterLibraryServiceServer(s, l)

	// Listen on the given TCP port.
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		glog.Fatalf("Failed to listen: %v", err)
	}

	// gRPC server serves on the port.
	if err := s.Serve(lis); err != nil {
		glog.Fatalf("Failed to serve: %v", err)
	}
}
