package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/dollarkillerx/grpc_discover"
	"github.com/dollarkillerx/zim/api/manager"
	"github.com/dollarkillerx/zim/api/protocol"
	"github.com/dollarkillerx/zim/pkg/enums"
	"github.com/dollarkillerx/zim/pkg/utils"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"
)

var (
	etcdAddress string
)

func init() {
	flag.StringVar(&etcdAddress, "etcdAddress", "127.0.0.1:2379", "etcd address")
	flag.Parse()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	plugin, err := grpc_discover.NewETCDPlugin(clientv3.Config{
		Endpoints:   []string{etcdAddress},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	resolver.Register(plugin)

	// auth

	// Set up a connection to the server.
	conn, err := grpc.Dial(
		fmt.Sprintf("etcd:///%s", enums.DiscoverManager),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(utils.NewAuthCredential("8e24a40a-7f13-469d-8141-e5a81df7b629")))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := manager.NewManagerClient(conn)
	sa, err := client.SuperAdminCreate(metadata.AppendToOutgoingContext(context.TODO(), "k1", "v1"), &protocol.Empty{})
	if err != nil {
		panic(err)
	}
	utils.PrintObject(sa)
}
