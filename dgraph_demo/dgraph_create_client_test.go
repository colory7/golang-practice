package dgraph_demo

import (
	"context"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
	"log"
	"testing"
)

func newClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

func newClientGrpc() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	dialOpts := append([]grpc.DialOption{},
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	d, err := grpc.Dial("localhost:9080", dialOpts...)

	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

func TestMultiTenancy(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		glog.Error("While trying to dial gRPC, got error", err)
	}
	dc := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	ctx := context.Background()
	// Login to namespace 123
	if err := dc.LoginIntoNamespace(ctx, "groot", "password", 123); err != nil {
		glog.Error("Failed to login: ", err)
	}

}

func TestCloud(t *testing.T) {
	// This example uses dgo
	conn, err := dgo.DialCloud("https://frozen-mango.eu-central-1.aws.cloud.dgraph.io/graphql", "<api-key>")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

}
