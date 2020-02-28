// Code generated by proroc-gen-graphql, DO NOT EDIT.
package book

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
	author "github.com/ysugimoto/grpc-graphql-gateway/examples/basic/app/author"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

var _ = json.Marshal
var _ = json.Unmarshal

var Gql__type_GetBookRequest = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetBookRequest",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "this is example comment for id field",
		},
	},
}) // message GetBookRequest in book/book.proto

var Gql__type_ListBooksRequest = graphql.NewObject(graphql.ObjectConfig{
	Name:   "ListBooksRequest",
	Fields: graphql.Fields{},
}) // message ListBooksRequest in book/book.proto

var Gql__type_ListBooksResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "ListBooksResponse",
	Fields: graphql.Fields{
		"books": &graphql.Field{
			Type: graphql.NewList(Gql__type_Book),
		},
	},
}) // message ListBooksResponse in book/book.proto

var Gql__type_Book_Description = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book_Description",
	Fields: graphql.Fields{
		"value": &graphql.Field{
			Type: graphql.String,
		},
	},
}) // message Book.Description in book/book.proto

var Gql__type_Timestamp = graphql.NewObject(graphql.ObjectConfig{
	Name: "Timestamp",
	Fields: graphql.Fields{
		"seconds": &graphql.Field{
			Type: graphql.Int,
		},
		"nanos": &graphql.Field{
			Type: graphql.Int,
		},
	},
}) // message Timestamp in google/protobuf/timestamp.proto

var Gql__type_CreateBookRequest = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreateBookRequest",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: Gql__enum_BookType,
		},
		"author": &graphql.Field{
			Type: author.Gql__type_Author,
		},
		"author_type": &graphql.Field{
			Type: author.Gql__enum_AuthorType,
		},
		"nested": &graphql.Field{
			Type: Gql__type_CreateBookRequest_NestedOne,
		},
		"created_at": &graphql.Field{
			Type: Gql__type_Timestamp,
		},
	},
}) // message CreateBookRequest in book/book.proto

var Gql__type_CreateBookRequest_NestedOne = graphql.NewObject(graphql.ObjectConfig{
	Name: "CreateBookRequest_NestedOne",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
	},
}) // message CreateBookRequest.NestedOne in book/book.proto

var Gql__type_Book = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: Gql__enum_BookType,
		},
		"author": &graphql.Field{
			Type: author.Gql__type_Author,
		},
		"description": &graphql.Field{
			Type: Gql__type_Book_Description,
		},
	},
}) // message Book in book/book.proto

var Gql__enum_BookType = graphql.NewEnum(graphql.EnumConfig{
	Name: "BookType",
	Values: graphql.EnumValueConfigMap{
		"JAVASCRIPT": &graphql.EnumValueConfig{
			Value: 0,
		},
		"ECMASCRIPT": &graphql.EnumValueConfig{
			Value: 1,
		},
		"GIT": &graphql.EnumValueConfig{
			Value: 2,
		},
		"ASP_DOT_NET": &graphql.EnumValueConfig{
			Value: 3,
		},
	},
}) // enum BookType in book/book.proto

var Gql__input_Author = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "Author",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
}) // message Author in author/author.proto

var Gql__input_Timestamp = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "Timestamp",
	Fields: graphql.InputObjectConfigFieldMap{
		"seconds": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"nanos": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
}) // message Timestamp in google/protobuf/timestamp.proto

var Gql__input_CreateBookRequest = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateBookRequest",
	Fields: graphql.InputObjectConfigFieldMap{
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"type": &graphql.InputObjectFieldConfig{
			Type: Gql__enum_BookType,
		},
		"author": &graphql.InputObjectFieldConfig{
			Type: Gql__input_Author,
		},
		"author_type": &graphql.InputObjectFieldConfig{
			Type: author.Gql__enum_AuthorType,
		},
		"nested": &graphql.InputObjectFieldConfig{
			Type: Gql__input_CreateBookRequest_NestedOne,
		},
		"created_at": &graphql.InputObjectFieldConfig{
			Type: Gql__input_Timestamp,
		},
	},
}) // message CreateBookRequest in book/book.proto

var Gql__input_CreateBookRequest_NestedOne = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateBookRequest_NestedOne",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
}) // message CreateBookRequest.NestedOne in book/book.proto

// graphql__resolver_BookService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_BookService struct {
	// grpc client connection.
	// this connection may be provided by user, then isAutoConnection should be false
	conn *grpc.ClientConn

	// isAutoConnection indicates that the grpc connection is opened by this handler.
	// If true, this handler opens connection automatically, and it should be closed on Close() method.
	isAutoConnection bool
}

// Close() closes grpc connection if it is opened automatically.
func (x *graphql__resolver_BookService) Close() error {
	// nothing to do because the connection is supplied by user, and it should be closed user themselves.
	if !x.isAutoConnection {
		return nil
	}
	return x.conn.Close()
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *graphql__resolver_BookService) GetQueries() graphql.Fields {
	return graphql.Fields{
		"books": &graphql.Field{
			Type: graphql.NewList(Gql__type_Book),
			Args: graphql.FieldConfigArgument{},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req *ListBooksRequest
				if err := runtime.MarshalRequest(p.Args, &req); err != nil {
					return nil, err
				}
				client := NewBookServiceClient(x.conn)
				resp, err := client.ListBooks(p.Context, req)
				if err != nil {
					return nil, err
				}
				return resp.GetBooks(), nil
			},
		},
		"book": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					Description:  "this is example comment for id field",
					DefaultValue: 10,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req *GetBookRequest
				if err := runtime.MarshalRequest(p.Args, &req); err != nil {
					return nil, err
				}
				client := NewBookServiceClient(x.conn)
				resp, err := client.GetBook(p.Context, req)
				if err != nil {
					return nil, err
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_BookService) GetMutations() graphql.Fields {
	return graphql.Fields{
		"createbook": &graphql.Field{
			Type: Gql__type_Book,
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"type": &graphql.ArgumentConfig{
					Type: Gql__enum_BookType,
				},
				"author": &graphql.ArgumentConfig{
					Type: Gql__input_Author,
				},
				"author_type": &graphql.ArgumentConfig{
					Type: author.Gql__enum_AuthorType,
				},
				"nested": &graphql.ArgumentConfig{
					Type: Gql__input_CreateBookRequest_NestedOne,
				},
				"created_at": &graphql.ArgumentConfig{
					Type: Gql__input_Timestamp,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req *CreateBookRequest
				if err := runtime.MarshalRequest(p.Args, &req); err != nil {
					return nil, err
				}
				client := NewBookServiceClient(x.conn)
				resp, err := client.CreateBook(p.Context, req)
				if err != nil {
					return nil, err
				}
				return resp, nil
			},
		},
	}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterBookServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterBookServiceGraphql(mux *runtime.ServeMux) error {
	return RegisterBookServiceGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, the resolver opens connection automatically and then you need to define host with FileOption like:
//
// service BookService {
//    option (graphql.service) = {
//        host: "your default host like localhost:50051";
//        insecure: true or false;
//    };
//
//    ...with RPC definitions
// }
//
func RegisterBookServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) (err error) {
	var isAutoConnection bool
	if conn == nil {
		isAutoConnection = true
		conn, err = grpc.Dial("localhost:8080", grpc.WithInsecure())
		if err != nil {
			return
		}
	}
	mux.AddHandler(&graphql__resolver_BookService{conn, isAutoConnection})
	return
}