// Package library provides the Library object which implements the library_service
package library

import (
	"strings"
	"sync"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	lspb "github.com/ghasemloo/libraryservice/proto/api"
	lpb "github.com/ghasemloo/libraryservice/proto/storage"
)

// Library object.
type Library struct {
	// mu protects books.
	mu sync.Mutex
	// books maps book id to .
	books map[string]*lpb.Book

	// Embed the unimplemented server.
	lspb.UnimplementedLibraryServiceServer
}

// New creates a new Library object.
func New(context.Context) (*Library, error) {
	return &Library{
		books: make(map[string]*lpb.Book),
	}, nil
}

// CreateBook creates a book in the library. Standard REST method.
func (l *Library) CreateBook(ctx context.Context, req *lspb.CreateBookRequest) (*lspb.Book, error) {
	glog.Infof("CreateBookRequest: %+v", req)

	// Parse the request.
	if req.GetParent() != "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parent resource name: %v", req.GetParent())
	}

	// Create a book from the request.
	book := &lpb.Book{
		Id:    req.GetBookId(),
		Title: req.GetBook().GetTitle(),
	}

	// Check if the book already exists, if not store a copy of the book.
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.books[book.Id]; ok {
		return nil, status.Errorf(codes.AlreadyExists, "There is already a book with id: %v", book.Id)
	}
	l.books[book.Id] = proto.Clone(book).(*lpb.Book)

	// Return the response.
	return &lspb.Book{
		Name:  "books/" + book.GetId(),
		Title: book.GetTitle(),
	}, nil
}

// GetBook gets a book in the library. Standard REST method.
func (l *Library) GetBook(ctx context.Context, req *lspb.GetBookRequest) (*lspb.Book, error) {
	glog.Infof("GetBookRequest: %+v", req)

	// Parse the request.
	name := strings.Split(req.GetName(), "/")
	if len(name) != 2 || name[0] != "books" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid book resource name: %v", req.GetName())
	}
	bookID := name[1]

	// Check if the the book exits.
	l.mu.Lock()
	defer l.mu.Unlock()
	book, ok := l.books[bookID]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "No book with id: %v", bookID)
	}

	// Return the response.
	return &lspb.Book{
		Name:  "books/" + book.GetId(),
		Title: book.GetTitle(),
	}, nil
}

var _ lspb.LibraryServiceServer = (*Library)(nil)
