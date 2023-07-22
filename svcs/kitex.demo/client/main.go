package main

import (
	"context"
	"fmt"
	"math/rand"

	serverz "day3/kxS/client/kitex_gen/kitex/serverZ"
	"day3/kxS/client/kitex_gen/kitex/serverZ/studentservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type rPicker struct {
	instances []discovery.Instance
}

func (r *rPicker) Next(ctx context.Context, request interface{}) discovery.Instance {
	fmt.Printf("fnd: %v \n", len(r.instances))
	return r.instances[rand.Int()%len(r.instances)]
}

type rBalancer struct {
}

func (b *rBalancer) Name() string {
	return "rBalancer"
}

func (b *rBalancer) GetPicker(r discovery.Result) loadbalance.Picker {
	return &rPicker{
		instances: r.Instances,
	}
}

func main() {
	var opts []client.Option
	opts = append(opts, client.WithHostPorts("localhost:8888"))
	opts = append(opts, client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 10,
		MaxIdlePerAddress: 1000}))
	r, _ := etcd.NewEtcdResolver([]string{"localhost:2379"})
	opts = append(opts, client.WithResolver(r))

	opts = append(opts, client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()))

	greetCli := studentservice.MustNewClient("kitex.demo", opts...)

	for i := 0; i < 20; i++ {
		req := &serverz.QueryReq{
			Id: int32(i),
		}

		_, err := greetCli.Query(context.Background(), req)
		if err == nil {
			klog.Infof("OK")
		} else {
			klog.Infof("F")
		}

	}

}
