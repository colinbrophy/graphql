package graphql

import (
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	"github.com/colin/UserServer/userserver"
	"flag"
	"net/http"
	"google.golang.org/grpc"
)

type User struct {
	FirstName string
	LastName string
}
var serverAddr = flag.String("server_addr", "localhost:6666", "")


func main() {

	//conn, _ := grpc.Dial(*serverAddr, grpc.WithInsecure())
	//client := userserver.NewUserServiceClient(conn)

	resolver := Resolver{};

	schema := graphql.MustParseSchema(Schema, &resolver)

	http.Handle("/", &relay.Handler{schema})

	http.ListenAndServe(":8084", nil)

}

