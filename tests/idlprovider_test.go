package tests

import (
	idlprovider "apigateway/IDLProvider"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/stretchr/testify/assert"
)

func TestLoadIDLContents(t *testing.T) {
	idlprovider.LoadIDLContents("./idl/svcPath", "./idl")

	// name
	assert.Equal(t, "kitex.demo", idlprovider.IdlNames["./idl/kxServer.thrift"])
	// path
	assert.Equal(t, "./idl/kxServer.thrift", idlprovider.IdlPaths["kitex.demo"])
	// contents
	file, err := os.Open("./idl/kxServer.thrift")
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, string(content), idlprovider.IdlContents["./idl/kxServer.thrift"])
}

func TestUpdateIDLContents(t *testing.T) {
	idlprovider.IdlNames["./idl/test1.thrift"] = "test"
	idlprovider.IdlPaths["test"] = "./idl/test1.thrift"

	file, err := os.Open("./idl/test2.thrift")
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	file.Close()

	idlprovider.IdlContents["./idl/test1.thrift"] = string(content)

	p, err := generic.NewThriftContentProvider(idlprovider.IdlContents["./idl/test1.thrift"], idlprovider.IdlContents)
	if err != nil {
		panic(err)
	}
	idlprovider.IdlProviders["test"] = p
	var tMutex sync.RWMutex

	idlprovider.IdlMutexs["test"] = &tMutex

	idlprovider.UpdateIDLContents("./idl/test1.thrift")

	file, err = os.Open("./idl/test1.thrift")
	if err != nil {
		panic(err)
	}
	content, err = io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	file.Close()

	assert.Equal(t, string(content), idlprovider.IdlContents["./idl/test1.thrift"])
	fmt.Println(idlprovider.IdlContents["./idl/test1.thrift"])

}

func TestWatchIDLFiles(t *testing.T) {
	// use debug or '-v' to check the output
	idlprovider.IdlNames["./idl/test1.thrift"] = "test"
	idlprovider.IdlPaths["test"] = "./idl/test1.thrift"

	file, err := os.Open("./idl/test2.thrift")
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	file.Close()

	idlprovider.IdlContents["./idl/test1.thrift"] = string(content)

	p, err := generic.NewThriftContentProvider(idlprovider.IdlContents["./idl/test1.thrift"], idlprovider.IdlContents)
	if err != nil {
		panic(err)
	}
	idlprovider.IdlProviders["test"] = p
	var tMutex sync.RWMutex

	idlprovider.IdlMutexs["test"] = &tMutex

	go idlprovider.WatchIDLFiles("./idl")
	file, err = os.OpenFile("./idl/test1.thrift", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	file.Write([]byte("  "))
	file.Close()
}
