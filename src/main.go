// Code generated by hertz generator.

package main

import (
	idlprovider "apigateway/IDLProvider"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	register(h)
	idlprovider.LoadIDLContents()
	go idlprovider.WatchIDLFiles()
	h.Spin()
}
