// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/example/library/v1/library.proto
// DO NOT EDIT!

/*
Package google_example_library_v1 is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/example/library/v1/library.proto

It has these top-level messages:
	Book
	Shelf
	CreateShelfRequest
	GetShelfRequest
	ListShelvesRequest
	ListShelvesResponse
	DeleteShelfRequest
	MergeShelvesRequest
	CreateBookRequest
	GetBookRequest
	ListBooksRequest
	ListBooksResponse
	UpdateBookRequest
	DeleteBookRequest
	MoveBookRequest
*/
package google_example_library_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A single book in the library.
type Book struct {
	// The resource name of the book.
	// Book names have the form `shelves/{shelf_id}/books/{book_id}`.
	// The name is ignored when creating a book.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The name of the book author.
	Author string `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	// The title of the book.
	Title string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	// Value indicating whether the book has been read.
	Read bool `protobuf:"varint,4,opt,name=read" json:"read,omitempty"`
}

func (m *Book) Reset()                    { *m = Book{} }
func (m *Book) String() string            { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()               {}
func (*Book) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// A Shelf contains a collection of books with a theme.
type Shelf struct {
	// The resource name of the shelf.
	// Shelf names have the form `shelves/{shelf_id}`.
	// The name is ignored when creating a shelf.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The theme of the shelf
	Theme string `protobuf:"bytes,2,opt,name=theme" json:"theme,omitempty"`
}

func (m *Shelf) Reset()                    { *m = Shelf{} }
func (m *Shelf) String() string            { return proto.CompactTextString(m) }
func (*Shelf) ProtoMessage()               {}
func (*Shelf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Request message for LibraryService.CreateShelf.
type CreateShelfRequest struct {
	// The shelf to create.
	Shelf *Shelf `protobuf:"bytes,1,opt,name=shelf" json:"shelf,omitempty"`
}

func (m *CreateShelfRequest) Reset()                    { *m = CreateShelfRequest{} }
func (m *CreateShelfRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateShelfRequest) ProtoMessage()               {}
func (*CreateShelfRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateShelfRequest) GetShelf() *Shelf {
	if m != nil {
		return m.Shelf
	}
	return nil
}

// Request message for LibraryService.GetShelf.
type GetShelfRequest struct {
	// The name of the shelf to retrieve.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetShelfRequest) Reset()                    { *m = GetShelfRequest{} }
func (m *GetShelfRequest) String() string            { return proto.CompactTextString(m) }
func (*GetShelfRequest) ProtoMessage()               {}
func (*GetShelfRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Request message for LibraryService.ListShelves.
type ListShelvesRequest struct {
	// Requested page size. Server may return fewer shelves than requested.
	// If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// A token identifying a page of results the server should return.
	// Typically, this is the value of
	// [ListShelvesResponse.next_page_token][google.example.library.v1.ListShelvesResponse.next_page_token]
	// returned from the previous call to `ListShelves` method.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListShelvesRequest) Reset()                    { *m = ListShelvesRequest{} }
func (m *ListShelvesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListShelvesRequest) ProtoMessage()               {}
func (*ListShelvesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Response message for LibraryService.ListShelves.
type ListShelvesResponse struct {
	// The list of shelves.
	Shelves []*Shelf `protobuf:"bytes,1,rep,name=shelves" json:"shelves,omitempty"`
	// A token to retrieve next page of results.
	// Pass this value in the
	// [ListShelvesRequest.page_token][google.example.library.v1.ListShelvesRequest.page_token]
	// field in the subsequent call to `ListShelves` method to retrieve the next
	// page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListShelvesResponse) Reset()                    { *m = ListShelvesResponse{} }
func (m *ListShelvesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListShelvesResponse) ProtoMessage()               {}
func (*ListShelvesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListShelvesResponse) GetShelves() []*Shelf {
	if m != nil {
		return m.Shelves
	}
	return nil
}

// Request message for LibraryService.DeleteShelf.
type DeleteShelfRequest struct {
	// The name of the shelf to delete.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteShelfRequest) Reset()                    { *m = DeleteShelfRequest{} }
func (m *DeleteShelfRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteShelfRequest) ProtoMessage()               {}
func (*DeleteShelfRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// Describes the shelf being removed (other_shelf_name) and updated
// (name) in this merge.
type MergeShelvesRequest struct {
	// The name of the shelf we're adding books to.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The name of the shelf we're removing books from and deleting.
	OtherShelfName string `protobuf:"bytes,2,opt,name=other_shelf_name,json=otherShelfName" json:"other_shelf_name,omitempty"`
}

func (m *MergeShelvesRequest) Reset()                    { *m = MergeShelvesRequest{} }
func (m *MergeShelvesRequest) String() string            { return proto.CompactTextString(m) }
func (*MergeShelvesRequest) ProtoMessage()               {}
func (*MergeShelvesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// Request message for LibraryService.CreateBook.
type CreateBookRequest struct {
	// The name of the shelf in which the book is created.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The book to create.
	Book *Book `protobuf:"bytes,2,opt,name=book" json:"book,omitempty"`
}

func (m *CreateBookRequest) Reset()                    { *m = CreateBookRequest{} }
func (m *CreateBookRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateBookRequest) ProtoMessage()               {}
func (*CreateBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CreateBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

// Request message for LibraryService.GetBook.
type GetBookRequest struct {
	// The name of the book to retrieve.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetBookRequest) Reset()                    { *m = GetBookRequest{} }
func (m *GetBookRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBookRequest) ProtoMessage()               {}
func (*GetBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// Request message for LibraryService.ListBooks.
type ListBooksRequest struct {
	// The name of the shelf whose books we'd like to list.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Requested page size. Server may return fewer books than requested.
	// If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// A token identifying a page of results the server should return.
	// Typically, this is the value of
	// [ListBooksResponse.next_page_token][google.example.library.v1.ListBooksResponse.next_page_token].
	// returned from the previous call to `ListBooks` method.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListBooksRequest) Reset()                    { *m = ListBooksRequest{} }
func (m *ListBooksRequest) String() string            { return proto.CompactTextString(m) }
func (*ListBooksRequest) ProtoMessage()               {}
func (*ListBooksRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// Response message for LibraryService.ListBooks.
type ListBooksResponse struct {
	// The list of books.
	Books []*Book `protobuf:"bytes,1,rep,name=books" json:"books,omitempty"`
	// A token to retrieve next page of results.
	// Pass this value in the
	// [ListBooksRequest.page_token][google.example.library.v1.ListBooksRequest.page_token]
	// field in the subsequent call to `ListBooks` method to retrieve the next
	// page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListBooksResponse) Reset()                    { *m = ListBooksResponse{} }
func (m *ListBooksResponse) String() string            { return proto.CompactTextString(m) }
func (*ListBooksResponse) ProtoMessage()               {}
func (*ListBooksResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ListBooksResponse) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

// Request message for LibraryService.UpdateBook.
type UpdateBookRequest struct {
	// The name of the book to update.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The book to update with. The name must match or be empty.
	Book *Book `protobuf:"bytes,2,opt,name=book" json:"book,omitempty"`
}

func (m *UpdateBookRequest) Reset()                    { *m = UpdateBookRequest{} }
func (m *UpdateBookRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateBookRequest) ProtoMessage()               {}
func (*UpdateBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *UpdateBookRequest) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

// Request message for LibraryService.DeleteBook.
type DeleteBookRequest struct {
	// The name of the book to delete.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteBookRequest) Reset()                    { *m = DeleteBookRequest{} }
func (m *DeleteBookRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteBookRequest) ProtoMessage()               {}
func (*DeleteBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

// Describes what book to move (name) and what shelf we're moving it
// to (other_shelf_name).
type MoveBookRequest struct {
	// The name of the book to move.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The name of the destination shelf.
	OtherShelfName string `protobuf:"bytes,2,opt,name=other_shelf_name,json=otherShelfName" json:"other_shelf_name,omitempty"`
}

func (m *MoveBookRequest) Reset()                    { *m = MoveBookRequest{} }
func (m *MoveBookRequest) String() string            { return proto.CompactTextString(m) }
func (*MoveBookRequest) ProtoMessage()               {}
func (*MoveBookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func init() {
	proto.RegisterType((*Book)(nil), "google.example.library.v1.Book")
	proto.RegisterType((*Shelf)(nil), "google.example.library.v1.Shelf")
	proto.RegisterType((*CreateShelfRequest)(nil), "google.example.library.v1.CreateShelfRequest")
	proto.RegisterType((*GetShelfRequest)(nil), "google.example.library.v1.GetShelfRequest")
	proto.RegisterType((*ListShelvesRequest)(nil), "google.example.library.v1.ListShelvesRequest")
	proto.RegisterType((*ListShelvesResponse)(nil), "google.example.library.v1.ListShelvesResponse")
	proto.RegisterType((*DeleteShelfRequest)(nil), "google.example.library.v1.DeleteShelfRequest")
	proto.RegisterType((*MergeShelvesRequest)(nil), "google.example.library.v1.MergeShelvesRequest")
	proto.RegisterType((*CreateBookRequest)(nil), "google.example.library.v1.CreateBookRequest")
	proto.RegisterType((*GetBookRequest)(nil), "google.example.library.v1.GetBookRequest")
	proto.RegisterType((*ListBooksRequest)(nil), "google.example.library.v1.ListBooksRequest")
	proto.RegisterType((*ListBooksResponse)(nil), "google.example.library.v1.ListBooksResponse")
	proto.RegisterType((*UpdateBookRequest)(nil), "google.example.library.v1.UpdateBookRequest")
	proto.RegisterType((*DeleteBookRequest)(nil), "google.example.library.v1.DeleteBookRequest")
	proto.RegisterType((*MoveBookRequest)(nil), "google.example.library.v1.MoveBookRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for LibraryService service

type LibraryServiceClient interface {
	// Creates a shelf, and returns the new Shelf.
	CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*Shelf, error)
	// Gets a shelf. Returns NOT_FOUND if the shelf does not exist.
	GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error)
	// Lists shelves. The order is unspecified but deterministic. Newly created
	// shelves will not necessarily be added to the end of this list.
	ListShelves(ctx context.Context, in *ListShelvesRequest, opts ...grpc.CallOption) (*ListShelvesResponse, error)
	// Deletes a shelf. Returns NOT_FOUND if the shelf does not exist.
	DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Merges two shelves by adding all books from the shelf named
	// `other_shelf_name` to shelf `name`, and deletes
	// `other_shelf_name`. Returns the updated shelf.
	// The book ids of the moved books may not be the same as the original books.
	//
	// Returns NOT_FOUND if either shelf does not exist.
	// This call is a no-op if the specified shelves are the same.
	MergeShelves(ctx context.Context, in *MergeShelvesRequest, opts ...grpc.CallOption) (*Shelf, error)
	// Creates a book, and returns the new Book.
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Gets a book. Returns NOT_FOUND if the book does not exist.
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Lists books in a shelf. The order is unspecified but deterministic. Newly
	// created books will not necessarily be added to the end of this list.
	// Returns NOT_FOUND if the shelf does not exist.
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	// Deletes a book. Returns NOT_FOUND if the book does not exist.
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Updates a book. Returns INVALID_ARGUMENT if the name of the book
	// is non-empty and does equal the previous name.
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error)
	// Moves a book to another shelf, and returns the new book. The book
	// id of the new book may not be the same as the original book.
	MoveBook(ctx context.Context, in *MoveBookRequest, opts ...grpc.CallOption) (*Book, error)
}

type libraryServiceClient struct {
	cc *grpc.ClientConn
}

func NewLibraryServiceClient(cc *grpc.ClientConn) LibraryServiceClient {
	return &libraryServiceClient{cc}
}

func (c *libraryServiceClient) CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := grpc.Invoke(ctx, "/google.example.library.v1.LibraryService/CreateShelf", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := grpc.Invoke(ctx, "/google.example.library.v1.LibraryService/GetShelf", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListShelves(ctx context.Context, in *ListShelvesRequest, opts ...grpc.CallOption) (*ListShelvesResponse, error) {
	out := new(ListShelvesResponse)
	err := grpc.Invoke(ctx, "/google.example.library.v1.LibraryService/ListShelves", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/google.example.library.v1.LibraryService/DeleteShelf", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) MergeShelves(ctx context.Context, in *MergeShelvesRequest, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := grpc.Invoke(ctx, "/google.example.library.v1.LibraryService/MergeShelves", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/google.example.library.v1.LibraryService/CreateBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/google.example.library.v1.LibraryService/GetBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := grpc.Invoke(ctx, "/google.example.library.v1.LibraryService/ListBooks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/google.example.library.v1.LibraryService/DeleteBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/google.example.library.v1.LibraryService/UpdateBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) MoveBook(ctx context.Context, in *MoveBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := grpc.Invoke(ctx, "/google.example.library.v1.LibraryService/MoveBook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LibraryService service

type LibraryServiceServer interface {
	// Creates a shelf, and returns the new Shelf.
	CreateShelf(context.Context, *CreateShelfRequest) (*Shelf, error)
	// Gets a shelf. Returns NOT_FOUND if the shelf does not exist.
	GetShelf(context.Context, *GetShelfRequest) (*Shelf, error)
	// Lists shelves. The order is unspecified but deterministic. Newly created
	// shelves will not necessarily be added to the end of this list.
	ListShelves(context.Context, *ListShelvesRequest) (*ListShelvesResponse, error)
	// Deletes a shelf. Returns NOT_FOUND if the shelf does not exist.
	DeleteShelf(context.Context, *DeleteShelfRequest) (*google_protobuf1.Empty, error)
	// Merges two shelves by adding all books from the shelf named
	// `other_shelf_name` to shelf `name`, and deletes
	// `other_shelf_name`. Returns the updated shelf.
	// The book ids of the moved books may not be the same as the original books.
	//
	// Returns NOT_FOUND if either shelf does not exist.
	// This call is a no-op if the specified shelves are the same.
	MergeShelves(context.Context, *MergeShelvesRequest) (*Shelf, error)
	// Creates a book, and returns the new Book.
	CreateBook(context.Context, *CreateBookRequest) (*Book, error)
	// Gets a book. Returns NOT_FOUND if the book does not exist.
	GetBook(context.Context, *GetBookRequest) (*Book, error)
	// Lists books in a shelf. The order is unspecified but deterministic. Newly
	// created books will not necessarily be added to the end of this list.
	// Returns NOT_FOUND if the shelf does not exist.
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	// Deletes a book. Returns NOT_FOUND if the book does not exist.
	DeleteBook(context.Context, *DeleteBookRequest) (*google_protobuf1.Empty, error)
	// Updates a book. Returns INVALID_ARGUMENT if the name of the book
	// is non-empty and does equal the previous name.
	UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
	// Moves a book to another shelf, and returns the new book. The book
	// id of the new book may not be the same as the original book.
	MoveBook(context.Context, *MoveBookRequest) (*Book, error)
}

func RegisterLibraryServiceServer(s *grpc.Server, srv LibraryServiceServer) {
	s.RegisterService(&_LibraryService_serviceDesc, srv)
}

func _LibraryService_CreateShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).CreateShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.example.library.v1.LibraryService/CreateShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).CreateShelf(ctx, req.(*CreateShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.example.library.v1.LibraryService/GetShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetShelf(ctx, req.(*GetShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListShelves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShelvesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListShelves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.example.library.v1.LibraryService/ListShelves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListShelves(ctx, req.(*ListShelvesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_DeleteShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).DeleteShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.example.library.v1.LibraryService/DeleteShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).DeleteShelf(ctx, req.(*DeleteShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_MergeShelves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeShelvesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).MergeShelves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.example.library.v1.LibraryService/MergeShelves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).MergeShelves(ctx, req.(*MergeShelvesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.example.library.v1.LibraryService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.example.library.v1.LibraryService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.example.library.v1.LibraryService/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.example.library.v1.LibraryService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.example.library.v1.LibraryService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_MoveBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).MoveBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.example.library.v1.LibraryService/MoveBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).MoveBook(ctx, req.(*MoveBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LibraryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.example.library.v1.LibraryService",
	HandlerType: (*LibraryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShelf",
			Handler:    _LibraryService_CreateShelf_Handler,
		},
		{
			MethodName: "GetShelf",
			Handler:    _LibraryService_GetShelf_Handler,
		},
		{
			MethodName: "ListShelves",
			Handler:    _LibraryService_ListShelves_Handler,
		},
		{
			MethodName: "DeleteShelf",
			Handler:    _LibraryService_DeleteShelf_Handler,
		},
		{
			MethodName: "MergeShelves",
			Handler:    _LibraryService_MergeShelves_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _LibraryService_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _LibraryService_GetBook_Handler,
		},
		{
			MethodName: "ListBooks",
			Handler:    _LibraryService_ListBooks_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _LibraryService_DeleteBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _LibraryService_UpdateBook_Handler,
		},
		{
			MethodName: "MoveBook",
			Handler:    _LibraryService_MoveBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/example/library/v1/library.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 854 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0x96, 0x43, 0x02, 0xe1, 0x84, 0xdf, 0x01, 0xa1, 0xac, 0x17, 0x16, 0x76, 0x04, 0x6c, 0x36,
	0xcb, 0xda, 0x02, 0xb4, 0x7b, 0x11, 0x69, 0x6f, 0xd8, 0xdd, 0x56, 0x95, 0xa0, 0x8d, 0x92, 0xf6,
	0x0e, 0x29, 0x72, 0x60, 0x70, 0x2c, 0x12, 0x8f, 0x6b, 0x3b, 0x11, 0x3f, 0xea, 0x45, 0x2b, 0x55,
	0x95, 0x7a, 0xdb, 0x57, 0xe8, 0x1b, 0xf5, 0x15, 0xfa, 0x20, 0x9d, 0x1f, 0x1b, 0x4c, 0x62, 0xc6,
	0xee, 0x45, 0x6f, 0xa2, 0x99, 0x33, 0x67, 0xbe, 0xf3, 0x9d, 0x9f, 0xf9, 0x1c, 0x78, 0x62, 0x53,
	0x6a, 0xf7, 0x89, 0x61, 0xd3, 0xbe, 0xe5, 0xda, 0x06, 0xf5, 0x6d, 0xd3, 0x26, 0xae, 0xe7, 0xd3,
	0x90, 0x9a, 0xf2, 0xc8, 0xf2, 0x9c, 0xc0, 0x24, 0x57, 0xd6, 0xc0, 0xeb, 0x13, 0xb3, 0xef, 0x74,
	0x7d, 0xcb, 0xbf, 0x36, 0x47, 0xfb, 0xf1, 0xd2, 0x10, 0xbe, 0xe8, 0xa7, 0x08, 0x27, 0x72, 0x34,
	0xe2, 0xd3, 0xd1, 0xbe, 0xfe, 0x2c, 0x5f, 0x08, 0xf6, 0x63, 0x06, 0xc4, 0x1f, 0x39, 0x67, 0xe4,
	0x8c, 0xba, 0x17, 0x8e, 0x6d, 0x5a, 0xae, 0x4b, 0x43, 0x2b, 0x74, 0xa8, 0x1b, 0xc8, 0x28, 0xfa,
	0xa1, 0xed, 0x84, 0xbd, 0x61, 0xd7, 0x38, 0xa3, 0x03, 0x53, 0xc2, 0x99, 0xe2, 0xa0, 0x3b, 0xbc,
	0x30, 0xbd, 0xf0, 0xda, 0x23, 0x8c, 0xe9, 0x80, 0x2d, 0xe4, 0xaf, 0xbc, 0x84, 0x4f, 0xa1, 0x78,
	0x44, 0xe9, 0x25, 0x42, 0x50, 0x74, 0xad, 0x01, 0xa9, 0x6a, 0x5b, 0x5a, 0x6d, 0xb6, 0x25, 0xd6,
	0x68, 0x0d, 0xa6, 0xad, 0x61, 0xd8, 0xa3, 0x7e, 0xb5, 0x20, 0xac, 0xd1, 0x0e, 0xad, 0x42, 0x29,
	0x74, 0xc2, 0x3e, 0xa9, 0x4e, 0x09, 0xb3, 0xdc, 0x70, 0x04, 0x9f, 0x58, 0xe7, 0xd5, 0x22, 0x33,
	0x96, 0x5b, 0x62, 0x8d, 0xf7, 0xa1, 0xd4, 0xee, 0x91, 0xfe, 0x45, 0x2a, 0x3c, 0x87, 0xe9, 0x11,
	0x66, 0x2c, 0x44, 0x30, 0x7c, 0x83, 0x8f, 0x01, 0xfd, 0xcb, 0xee, 0x86, 0x44, 0x5c, 0x6c, 0x91,
	0xd7, 0x43, 0x12, 0x84, 0xe8, 0x6f, 0x28, 0x05, 0x7c, 0x2f, 0x00, 0x2a, 0x07, 0x5b, 0xc6, 0xa3,
	0x15, 0x35, 0xe4, 0x3d, 0xe9, 0x8e, 0x77, 0x60, 0xf1, 0x29, 0x09, 0x1f, 0x40, 0xa5, 0x50, 0xc1,
	0x4d, 0x40, 0xc7, 0x4e, 0x20, 0xfc, 0x46, 0x24, 0x88, 0x3d, 0x7f, 0x86, 0x59, 0xcf, 0xb2, 0x49,
	0x27, 0x70, 0x6e, 0xa4, 0x7b, 0xa9, 0x55, 0xe6, 0x86, 0x36, 0xdb, 0xa3, 0x0d, 0x00, 0x71, 0x18,
	0xd2, 0x4b, 0xe2, 0x46, 0x29, 0x08, 0xf7, 0x97, 0xdc, 0x80, 0xaf, 0x61, 0xe5, 0x01, 0x62, 0xe0,
	0xb1, 0x46, 0x11, 0xd4, 0x80, 0x99, 0x40, 0x9a, 0x18, 0xe0, 0x54, 0xae, 0x4c, 0xe2, 0x0b, 0x68,
	0x17, 0x16, 0x5d, 0x72, 0x15, 0x76, 0x26, 0xc2, 0xce, 0x73, 0x73, 0xf3, 0x2e, 0x74, 0x0d, 0xd0,
	0x7f, 0xa4, 0x4f, 0xc6, 0x2a, 0x98, 0x96, 0x76, 0x1b, 0x56, 0x4e, 0x88, 0x6f, 0x93, 0xb1, 0xbc,
	0xd3, 0x9a, 0x55, 0x83, 0x25, 0xca, 0x1a, 0xe4, 0x77, 0x44, 0x5d, 0x3b, 0xe2, 0x5c, 0x46, 0x5f,
	0x10, 0x76, 0x11, 0xeb, 0x39, 0x07, 0x3d, 0x85, 0x65, 0xd9, 0x40, 0x3e, 0x57, 0x2a, 0xc8, 0x43,
	0x28, 0x76, 0x99, 0x8b, 0x80, 0xa9, 0x1c, 0x6c, 0x2a, 0x0a, 0x21, 0x90, 0x84, 0x33, 0xde, 0x86,
	0x05, 0xd6, 0xd0, 0x0c, 0x68, 0xdc, 0x85, 0x25, 0x5e, 0x7d, 0xee, 0xa6, 0xcc, 0xea, 0x41, 0x87,
	0x0b, 0xca, 0x0e, 0x4f, 0x8d, 0x77, 0xd8, 0x87, 0xe5, 0x44, 0x8c, 0xa8, 0xbf, 0x7f, 0x41, 0x89,
	0xd3, 0x8c, 0xbb, 0x9b, 0x99, 0x94, 0xf4, 0xce, 0xdd, 0x5a, 0x56, 0xdb, 0x57, 0xde, 0xf9, 0x8f,
	0xaa, 0xed, 0x6f, 0xb0, 0x2c, 0x07, 0x27, 0xab, 0xbc, 0x2f, 0x60, 0xf1, 0x84, 0x8e, 0x32, 0x49,
	0xe4, 0x9e, 0x99, 0x83, 0xcf, 0x15, 0x58, 0x38, 0x96, 0x9c, 0xda, 0x52, 0xe5, 0xd0, 0x0d, 0x54,
	0x12, 0x3a, 0x80, 0xfe, 0x54, 0xa4, 0x30, 0xa9, 0x17, 0x7a, 0xe6, 0xb3, 0xc2, 0xfa, 0xbb, 0x2f,
	0x5f, 0x3f, 0x15, 0x56, 0x71, 0x85, 0xcb, 0x75, 0xf4, 0xc4, 0x1a, 0x52, 0x35, 0xd0, 0x08, 0xca,
	0xb1, 0x6a, 0xa0, 0xba, 0x02, 0x69, 0x4c, 0x5a, 0x72, 0x44, 0x5d, 0x17, 0x51, 0xd7, 0xd0, 0x2a,
	0x8f, 0x7a, 0xcb, 0x2b, 0xf2, 0x4f, 0x14, 0xdb, 0xac, 0xbf, 0x41, 0x6f, 0x35, 0xa8, 0x24, 0x54,
	0x43, 0x99, 0xf4, 0xa4, 0x5e, 0xe9, 0x46, 0x5e, 0x77, 0x39, 0xac, 0x78, 0x45, 0x90, 0x99, 0x47,
	0xc9, 0x12, 0x20, 0x1f, 0x2a, 0x09, 0xf5, 0x50, 0x52, 0x98, 0x54, 0x19, 0x7d, 0x2d, 0x76, 0x8f,
	0xbf, 0x3c, 0xc6, 0xff, 0xfc, 0x63, 0x13, 0xe7, 0x5d, 0x4f, 0xcf, 0xfb, 0x83, 0x06, 0x73, 0x49,
	0x21, 0x42, 0xaa, 0x4c, 0x52, 0x14, 0x2b, 0x47, 0xe1, 0x77, 0x04, 0x81, 0x4d, 0xac, 0xa7, 0x11,
	0x68, 0x0c, 0x38, 0x66, 0x43, 0xab, 0xa3, 0xf7, 0x1a, 0xc0, 0xbd, 0x7a, 0xa1, 0xbd, 0xcc, 0xa9,
	0x4b, 0xbc, 0x01, 0x3d, 0xeb, 0x99, 0xe1, 0x9a, 0x20, 0x81, 0xd3, 0x49, 0x98, 0x42, 0x0a, 0x1a,
	0xe2, 0x29, 0xa2, 0x5b, 0x98, 0x89, 0x64, 0x0e, 0xfd, 0xae, 0x1e, 0xc0, 0xef, 0x22, 0xb0, 0x2d,
	0x08, 0xfc, 0x82, 0xd6, 0x53, 0x08, 0xc8, 0xf8, 0xbc, 0x1d, 0x1f, 0x35, 0x98, 0xbd, 0x93, 0x36,
	0xf4, 0x47, 0xc6, 0x54, 0x25, 0x45, 0x56, 0xdf, 0xcb, 0xe7, 0x1c, 0x0d, 0x20, 0x16, 0x74, 0xd6,
	0x91, 0xa2, 0x1e, 0xe8, 0x0a, 0xe0, 0x5e, 0x94, 0x94, 0x0d, 0x99, 0xd0, 0xae, 0x47, 0xa7, 0x31,
	0x2a, 0x43, 0x5d, 0x5d, 0x06, 0x36, 0x95, 0x70, 0xaf, 0xb6, 0xca, 0xd0, 0x13, 0xa2, 0x9c, 0xdd,
	0x8a, 0xba, 0xe0, 0xb0, 0xad, 0x2b, 0x39, 0x44, 0xd3, 0xc0, 0xa6, 0xb2, 0x1c, 0x0b, 0xae, 0x52,
	0x90, 0xc6, 0x54, 0x39, 0x9b, 0xc5, 0x9e, 0x60, 0xb1, 0x8b, 0x7f, 0x55, 0xb2, 0x18, 0x30, 0x58,
	0xf6, 0x3a, 0x8e, 0x4c, 0xd8, 0x60, 0x7f, 0x2e, 0x1f, 0xc7, 0x3c, 0x9a, 0x8b, 0x44, 0xbc, 0xc9,
	0xeb, 0xdd, 0xd4, 0xba, 0xd3, 0xa2, 0xf0, 0x87, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x39,
	0x37, 0x73, 0x49, 0x0b, 0x00, 0x00,
}
