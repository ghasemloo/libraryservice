package library

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	lspb "github.com/ghasemloo/libraryservice/proto/library_service_proto"
)

var (
	bookID    = "Fake Book ID"
	bookTitle = "Fake Book Title"
	book      = &lspb.Book{
		Name:  "books/" + bookID,
		Title: bookTitle,
	}
)

func TestSimpleScenario(t *testing.T) {
	ctx := context.Background()

	// Create a Library.
	lib, err := New(ctx)
	if err != nil {
		t.Fatalf("New(_) failed: %v", err)
	}

	// Create a book.
	createReq := &lspb.CreateBookRequest{
		Parent: "",
		BookId: bookID,
		Book: &lspb.Book{
			Title: bookTitle,
		},
	}
	createResp, err := lib.CreateBook(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateBook(_, %+v) failed: %v", createReq, err)
	}
	if got, want := createResp, book; !proto.Equal(got, want) {
		t.Fatalf("CreateBook(_, %+v) = %+v, want %+v", createReq, got, want)
	}

	// Get the book.
	getReq := &lspb.GetBookRequest{
		Name: createResp.Name,
	}
	getResp, err := lib.GetBook(ctx, getReq)
	if err != nil {
		t.Fatalf("GetBook(_, %+v) failed: %v", getResp, err)
	}
	if got, want := getResp, book; !proto.Equal(got, want) {
		t.Fatalf("GetBook(_, %+v) = %+v, want %+v", getResp, got, want)
	}
}
