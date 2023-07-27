package kxcliprovider

import (
	idlprovider "apigateway/IDLProvider"
	"sync"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/loadbalance/lbcache"
	etcd "github.com/kitex-contrib/registry-etcd"
	//ruleBasedResolver "github.com/kitex-contrib/resolver-rule-based"
)

var etcdEndPoints = []string{
	// configure your etcd address here
	// "YOUR_ETCD_ADDRESS_HERE",
	"localhost:2379",
}

// svcname : *cli
var gClis = make(map[string]*genericclient.Client)

func GetGenericCli(svcName string) genericclient.Client {

	if gClis[svcName] == nil {
		var opts []client.Option

		opts = append(opts, client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 10,
			MaxIdlePerAddress: 1000}))

		r, _ := etcd.NewEtcdResolver(etcdEndPoints)

		// filterFunc := func(ctx context.Context, instance []discovery.Instance) []discovery.Instance {
		// 	var res []discovery.Instance
		// 	for _, ins := range instance {
		// 		if v, ok := ins.Tag("Cluster"); ok && v == "A" {
		// 			res = append(res, ins)
		// 		}
		// 	}
		// 	return res
		// }

		// filterRule := &ruleBasedResolver.FilterRule{
		// 	Name:  "coded-cluster",
		// 	Funcs: []ruleBasedResolver.FilterFunc{filterFunc},
		// }
		// r = ruleBasedResolver.NewRuleBasedResolver(r, filterRule)

		opts = append(opts, client.WithResolver(r))

		// IMPROVEMENTZ: longer time
		opts = append(opts, client.WithLoadBalancer(
			loadbalance.NewWeightedRandomBalancer(),
			&lbcache.Options{
				RefreshInterval: 30 * time.Second,
				ExpireInterval:  60 * time.Second,
			}))

		p, err := idlprovider.GetContentProvider(svcName)

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

var gCliCount = make(map[string]int)
var gCliPool = make(map[string][10]*genericclient.Client)
var mutex sync.Mutex

func GetGenericCliFromCliPool(svcName string) genericclient.Client {


	if _, ok := gCliPool[svcName]; !ok {
		mutex.Lock()
		defer mutex.Unlock()
		var arr [10]*genericclient.Client
		gCliPool[svcName] = arr
		gCliCount[svcName] = 1
		var cli genericclient.Client

		for i := 0; i < 10; i++ {
			var opts []client.Option

			opts = append(opts, client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 10,
				MaxIdlePerAddress: 1000}))

			r, _ := etcd.NewEtcdResolver(etcdEndPoints)

			opts = append(opts, client.WithResolver(r))

			opts = append(opts, client.WithLoadBalancer(
				loadbalance.NewWeightedRandomBalancer(),
				&lbcache.Options{
					RefreshInterval: 30 * time.Second,
					ExpireInterval:  40 * time.Second,
				}))

			p, err := idlprovider.GetContentProvider(svcName)

			idlprovider.IdlProviders[svcName] = p

			if err != nil {
				panic(err)
			}

			g, err := generic.JSONThriftGeneric(p)
			if err != nil {
				panic(err)
			}

			cli, err = genericclient.NewClient(svcName, g, opts...)

			if err != nil {
				panic(err)
			}
			pool := gCliPool[svcName]
			pool[i] = &cli
			gCliPool[svcName] = pool

		}
		return cli
	}

	mutex.Lock()
	defer mutex.Unlock()
	ret := gCliPool[svcName][gCliCount[svcName]]
	gCliCount[svcName] = (gCliCount[svcName] + 1) % 10
	return *ret
}
