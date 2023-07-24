package kxcliprovider

import (
	"context"

	idlprovider "apigateway/IDLProvider"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/generic"
	etcd "github.com/kitex-contrib/registry-etcd"
	ruleBasedResolver "github.com/kitex-contrib/resolver-rule-based"
)

// svcname : *cli
var gClis = make(map[string]*genericclient.Client)

func GetGenericCli(svcName string) genericclient.Client {

	idlPath := idlprovider.IdlPaths[svcName]

	if gClis[svcName] == nil {
		var opts []client.Option

		opts = append(opts, client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 10,
			MaxIdlePerAddress: 1000}))

		r, _ := etcd.NewEtcdResolver([]string{"localhost:2379"})

		filterFunc := func(ctx context.Context, instance []discovery.Instance) []discovery.Instance {
			var res []discovery.Instance
			for _, ins := range instance {
				if v, ok := ins.Tag("Cluster"); ok && v == "A" {
					res = append(res, ins)
				}
			}
			return res
		}

		filterRule := &ruleBasedResolver.FilterRule{
			Name:  "coded-cluster",
			Funcs: []ruleBasedResolver.FilterFunc{filterFunc},
		}
		r = ruleBasedResolver.NewRuleBasedResolver(r, filterRule)

		opts = append(opts, client.WithResolver(r))

		p, err := generic.NewThriftContentProvider(idlprovider.IdlContents[idlPath], idlprovider.IdlContents)

		idlprovider.IdlProviders[svcName] = p

		if err != nil {
			panic(err)
		}

		g, err := generic.JSONThriftGeneric(p)
		if err != nil {
			panic(err)
		}

		cli, err := genericclient.NewClient(svcName, g, opts...)

		if err != nil {
			panic(err)
		}

		gClis[svcName] = &cli
	}

	return *gClis[svcName]

}
