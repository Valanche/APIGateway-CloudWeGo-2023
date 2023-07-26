package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"

	"net/http"
	_ "net/http/pprof" // injecting routing into http

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"

	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	go http.ListenAndServe("localhost:8082", nil)

	handler := new(StudentServiceImpl)
	handler.InitDB()

	r, err := etcd.NewEtcdRegistry([]string{"localhost:2379"})
	if err != nil {
		log.Fatal(err)
	}

	p, err := generic.NewThriftFileProvider("./kxServer.thrift")
	if err != nil {
		panic(err)
	}
	// 构造 JSON 请求和返回类型的泛化调用
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	porty := fmt.Sprint(9000 + rand.Uint32()%10)

	addr, _ := net.ResolveTCPAddr("tcp", ":"+porty)
	svr := genericserver.NewServer(handler,
		g,
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "kitex.demo",
		}),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{
			Tags: map[string]string{
				"Cluster": "A",
			},
		}))

	err = svr.Run()
	if err != nil {
		log.Fatal(err)
	}

}
