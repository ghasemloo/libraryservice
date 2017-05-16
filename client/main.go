// Binary client is a LibraryService client.
package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	lspb "github.com/ghasemloo/libraryservice/proto/library_service_proto"
)

var (
	addr      = flag.String("addr", "localhost:9000", "LibraryService server address")
	bookID    = flag.String("book_id", "0134190564", "The ISBN of the book")
	bookTitle = flag.String("title", "The Go Programming Language", "The title of the book")
)

func main() {
	flag.Parse()
	glog.Info("LibraryService client")
	ctx := context.Background()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		glog.Exitf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a LibraryService client.
	c := lspb.NewLibraryServiceClient(conn)

	// Contact the server and create a book.
	createReq := &lspb.CreateBookRequest{
		BookId: *bookID,
		Book: &lspb.Book{
			Title: *bookTitle,
		},
	}
	book, err := c.CreateBook(ctx, createReq)
	if err != nil {
		glog.Exitf("CreateBook(%+v) failed: %v", createReq, err)
	}
	fmt.Printf("Created book: %+v\n", book)

	// Contact the server to check that the book is created.
	getReq := &lspb.GetBookRequest{
		Name: book.GetName(),
	}
	book, err = c.GetBook(ctx, getReq)
	if err != nil {
		glog.Exitf("GetBook(%+v) failed: %v", createReq, err)
	}
	fmt.Printf("Got book: %+v\n", book)
}
