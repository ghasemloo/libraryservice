package tests

import (
	"context"
	"testing"

	"github.com/ghasemloo/libraryservice/server/library"
	"github.com/ghasemloo/libraryservice/tests/fakegrpc"
	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"

	lspb "github.com/ghasemloo/libraryservice/proto"
)

const (
	id    = "0134190564"
	name  = "books/0134190564"
	title = "The Go Programming Language"
)

func TestCreateBook(t *testing.T) {
	ctx := context.Background()

	rpc, cleanup := fakegrpc.New(t)
	defer cleanup()

	// Create a Library server.
	s, err := library.New(ctx)
	if err != nil {
		t.Fatalf("library.New(ctx) failed: %v", err)
	}

	// Register LibraryService server on the gRPC server.
	lspb.RegisterLibraryServiceServer(rpc.Server, s)

	stop := rpc.Start(t)
	defer stop()

	// Create a LibraryService client using the gRPC clinet.
	c := lspb.NewLibraryServiceClient(rpc.Client)

	// Contact the server and create a book.
	req := &lspb.CreateBookRequest{
		BookId: id,
		Book:   &lspb.Book{Title: title},
	}
	got, err := c.CreateBook(ctx, req)
	if err != nil {
		t.Fatalf("CreateBook(%+v) failed: %v", req, err)
	}
	want := &lspb.Book{
		Name:  name,
		Title: title,
	}
	if diff := cmp.Diff(want, got, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("CreateBook(_,%+v) returned diff (-want +got):\n%s", req, diff)
	}
}
