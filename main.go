package main

import (
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	"flag"
	"net/http"
)

var serverAddr = flag.String("server_addr", "localhost:6666", "")


func main() {

	//conn, _ := grpc.Dial(*serverAddr, grpc.WithInsecure())
	//client := userserver.NewUserServiceClient(conn)

	resolver := Resolver{}

	schema := graphql.MustParseSchema(Schema, &resolver)

	http.Handle("/", &relay.Handler{schema})

	http.ListenAndServe(":8084", nil)

}

