package fakegrpc

import (
	"net"
	"testing"

	"google.golang.org/grpc"
)

// Fake contains the server and client for a GRPC connection.
type Fake struct {
	Listener net.Listener
	Server   *grpc.Server
	Client   *grpc.ClientConn
}

// New creates a gRPC client and server connected to eachother.
func New(t *testing.T) (*Fake, func() error) {
	t.Helper()

	f := &Fake{}

	// Create a gRPC server.
	f.Server = grpc.NewServer()

	// ":0" means pick a random free port on the local host.
	port := ":0"

	// Listen on the given TCP port.
	var err error
	f.Listener, err = net.Listen("tcp", port)
	if err != nil {
		t.Fatalf("net.Listen(\"tcp\", %v) failed: %v", port, err)
	}

	// Set up a client connection to the server.
	addr := f.Listener.Addr().String()
	f.Client, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial(%v, _) failed: %v", addr, err)
	}

	return f, f.Client.Close
}

// Start starts the server.
// Must be called after registering the services on the server.
func (f *Fake) Start(t *testing.T) func() error {
	t.Helper()

	// gRPC server serves on the port.
	go func() {
		if err := f.Server.Serve(f.Listener); err != nil {
			t.Fatalf("server.Serve(_) failed: %v", err)
		}
	}()

	return func() error {
		f.Server.Stop()
		return nil
	}
}
