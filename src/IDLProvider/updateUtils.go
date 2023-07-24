package idlprovider

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/fsnotify/fsnotify"
)

// path : content
var IdlContents = make(map[string]string)

// svcname : path
var IdlPaths = make(map[string]string)

// path : svcname
var IdlNames = make(map[string]string)

// svcname : *p
var IdlProviders = make(map[string]*generic.ThriftContentProvider)

var idlDir = "./idl"
var svcRelation = idlDir + "/svcPath"

// svcname : *M
var idlMutexs = make(map[string]*sync.RWMutex)
var svcMutex sync.Mutex

func walkIdl(filename string, fi os.FileInfo, err error) error {

	if fi.IsDir() {
		return nil
	}
	if strings.HasSuffix(strings.ToUpper(fi.Name()), "THRIFT") {
		//TODO: adapt to dir structures
		tftPath := idlDir + "/" + fi.Name()
		file, err := os.Open(tftPath)
		if err != nil {
			panic(err)
		}
		content, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		fmt.Println(tftPath)

		IdlContents[tftPath] = string(content)

		file.Close()
	}

	return nil
}

func LoadIDLContents() {
	file, err := os.Open(svcRelation)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err != nil || io.EOF == err {
			break
		}

		svcR := strings.Split(string(line), ", ")
		if len(svcR) < 2 {
			break
		}
		IdlPaths[svcR[0]] = svcR[1]
		IdlNames[svcR[1]] = svcR[0]

		var rwlock sync.RWMutex
		idlMutexs[svcR[0]] = &rwlock

	}
	if err != nil {
		panic(err)
	}
	file.Close()

	err = filepath.Walk(idlDir, walkIdl)

	if err != nil {
		panic(err)
	}
}

func UpdateIDLContents(idlfilepath string) {

	if _, ok := IdlContents[idlfilepath]; !ok {
		klog.Infof("Unknown file (might be svcR) " + idlfilepath)
		return
	}
	svcname := IdlNames[idlfilepath]

	idlMutexs[svcname].RLock()
	file, err := os.Open(idlfilepath)
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	idlMutexs[svcname].RUnlock()
	file.Close()

	IdlContents[idlfilepath] = string(content)
	if IdlProviders[svcname] != nil {
		IdlProviders[svcname].UpdateIDL(string(content), IdlContents)
	}
	fmt.Println("Updated: " + idlfilepath)

}

func WatchIDLFiles() {
	idlWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer idlWatcher.Close()

	err = idlWatcher.Add(idlDir)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event := <-idlWatcher.Events:
			if event.Op.String() == "WRITE" {
				fmt.Println(event.Name)
				tftPath := "./" + event.Name
				go UpdateIDLContents(tftPath)
			}

		case err := <-idlWatcher.Errors:
			fmt.Println("error: ", err)
		}
	}
}
