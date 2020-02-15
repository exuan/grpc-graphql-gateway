package main

import (
	"log"
	"os"

	"io/ioutil"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql/generator"
	"github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql/spec"
)

func main() {
	var genError error

	resp := &plugin.CodeGeneratorResponse{}
	defer func() {
		// If some error has been occurred in generate process,
		// add error message to plugin response
		if genError != nil {
			message := genError.Error()
			resp.Error = &message
		}
		buf, err := proto.Marshal(resp)
		if err != nil {
			log.Fatalln(err)
		}
		os.Stdout.Write(buf)
	}()

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		genError = err
		return
	}

	var req plugin.CodeGeneratorRequest
	if err := proto.Unmarshal(buf, &req); err != nil {
		genError = err
		return
	}

	var args *spec.Params
	if req.Parameter != nil {
		args, err = spec.NewParams(req.GetParameter())
		if err != nil {
			genError = err
			return
		}
	}

	// We're dealing with each descriptors to out wrapper struct
	// in order to access easily plugin options, pakcage name, comment, etc...
	var files []*spec.File
	for _, f := range req.GetProtoFile() {
		files = append(files, spec.NewFile(f))
	}

	g := generator.New(generator.GenerationTypeGo)
	templates, err := g.Generate(files, args)
	if err != nil {
		genError = err
		return
	}
	for _, t := range templates {
		f, err := t.Execute(goTemplate)
		if err != nil {
			genError = err
			return
		}
		resp.File = append(resp.File, f)
	}
}

var goTemplate = `
// Code generated by proroc-gen-graphql, DO NOT EDIT.
package {{ .RootPackage.Name }}

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"

{{- range .Packages }}
	{{ .Name }} "{{ .Path }}"
{{ end }}
)

var _ = json.Marshal
var _ = json.Unmarshal

{{ range .Types -}}
var Gql__type_{{ .TypeName }} = graphql.NewObject(graphql.ObjectConfig{
	Name: "{{ .Name }}",
	{{- if .Comment }}
	Description: "{{ .Comment }}",
	{{- end }}
	Fields: graphql.Fields {
{{- range .Fields }}
		"{{ .Name }}": &graphql.Field{
			Type: {{ .FieldType $.RootPackage.Path }},
			{{- if .Comment }}
			Description: "{{ .Comment }}",
			{{- end }}
		},
{{- end }}
	},
}) // message {{ .Name }} in {{ .Filename }}

{{ end }}

{{ range .Enums -}}
var Gql__enum_{{ .Name }} = graphql.NewEnum(graphql.EnumConfig{
	Name: "{{ .Name }}",
	Values: graphql.EnumValueConfigMap{
{{- range .Values }}
		"{{ .Name }}": &graphql.EnumValueConfig{
			{{- if .Comment }}
			Description: "{{ .Comment }}",
			{{- end }}
			Value: {{ .Number }},
		},
{{- end }}
	},
}) // enum {{ .Name }} in {{ .Filename }}
{{ end }}

{{ range .Inputs -}}
var Gql__input_{{ .Name }} = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "{{ .Name }}",
	Fields: graphql.InputObjectConfigFieldMap{
{{- range .Fields }}
		"{{ .Name }}": &graphql.InputObjectFieldConfig{
			{{- if .Comment }}
			Description: "{{ .Comment }}",
			{{- end }}
			Type: {{ .FieldType $.RootPackage.Path }},
		},
{{- end }}
	},
}) // message {{ .Name }} in {{ .Filename }}
{{ end }}

// graphql__resolver_{{ .Service.Name }} is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.Resolver interface.
type graphql__resolver_{{ .Service.Name }} struct {
	conn *grpc.ClientConn
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *graphql__resolver_{{ .Service.Name }}) GetQueries() graphql.Fields {
	return graphql.Fields{
{{- range .Queries }}
		"{{ .QueryName }}": &graphql.Field{
			Type: {{ .QueryType }},
			{{- if .Comment }}
			Description: "{{ .Comment }}",
			{{- end }}
			{{- if .Args }}
			Args: graphql.FieldConfigArgument{
			{{- range .Args }}
				"{{ .Name }}": &graphql.ArgumentConfig{
					Type: {{ .FieldType $.RootPackage.Path }},
					{{- if .Comment }}
					Description: "{{ .Comment }}",
					{{- end }}
					{{- if .DefaultValue }}
					DefaultValue: {{ .DefaultValue }},
					{{- end }}
				},
			{{- end }}
			},
			{{- end }}
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req *{{ .RequestType }}
				if err := runtime.MarshalRequest(p.Args, req); err != nil {
					return nil, err
				}
				client := {{ .Package }}New{{ .Service.Name }}Client(x.conn)
				resp, err := client.{{ .Method.Name }}(p.Context, req)
				if err != nil {
					return nil, err
				}
				{{- if .Expose }}
				return resp.Get{{ .Expose }}(), nil
				{{- else }}
				return resp, nil
				{{- end }}
			},
		},
{{- end }}
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_{{ .Service.Name }}) GetMutations() graphql.Fields {
	return graphql.Fields{
{{- range .Mutations }}
		"{{ .MutationName }}": &graphql.Field{
			Type: {{ .MutationType }},
			{{- if .Comment }}
			Description: "{{ .Comment }}",
			{{ end }}
			Args: graphql.FieldConfigArgument{
				"{{ .InputName }}": &graphql.ArgumentConfig{
					Type: Gql__input_{{ .Input.Name }},
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req *{{ .RequestType }}
				if err := runtime.MarshalRequest(p.Args, req); err != nil {
					return nil, err
				}
				client := {{ .Package }}New{{ .Service.Name }}Client(x.conn)
				resp, err := client.{{ .Method.Name }}(p.Context, req)
				if err != nil {
					return nil, err
				}
				{{- if .Expose }}
				return resp.Get{{ .Expose }}(), nil
				{{- else }}
				return resp, nil
				{{- end }}
			},
		},
{{ end }}
	}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you worried about open/close performance for each handling graphql request,
// then you can call RegisterBookHandler with *grpc.ClientConn manually.
func Register{{ .Service.Name }}Graphql(mux *runtime.ServeMux) error {
	return Register{{ .Service.Name }}GraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when appication will terminate.
// Otherwise, the resolver opens connection automatically, but then you need to define host with ServiceOption like:
//
// service XXXService {
//    option (graphql.service) = {
//        host: "localhost:50051"
//    };
//
//    ...with RPC definitions
// }
func Register{{ .Service.Name }}GraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) (err error) {
	if conn == nil {
		conn, err = grpc.Dial("{{ .Service.Host }}"{{ if .Service.Insecure }}, grpc.WithInsecure(){{ end }})
		if err != nil {
			return
		}
	}
	mux.AddHandler(&graphql__resolver_{{ .Service.Name }}{conn})
	return
}`