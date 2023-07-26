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

// svcname : *M
var IdlMutexs = make(map[string]*sync.RWMutex)
var svcMutex sync.Mutex

func LoadIDLContents(svcPath string, idlDirPath string) {
	file, err := os.Open(svcPath)
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
		IdlMutexs[svcR[0]] = &rwlock

	}
	if err != nil {
		panic(err)
	}
	file.Close()

	err = filepath.Walk(idlDirPath,
		func(filename string, fi os.FileInfo, err error) error {

			if fi.IsDir() {
				return nil
			}
			if strings.HasSuffix(strings.ToUpper(fi.Name()), "THRIFT") {
				//TODO: adapt to dir structures
				tftPath := idlDirPath + "/" + fi.Name()
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
		})

	if err != nil {
		panic(err)
	}
}

func UpdateIDLContents(idlfilepath string) {

	if _, ok := IdlContents[idlfilepath]; !ok {
		klog.Infof("Unknown file " + idlfilepath)
		return
	}
	svcname := IdlNames[idlfilepath]

	IdlMutexs[svcname].RLock()
	file, err := os.Open(idlfilepath)
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	IdlMutexs[svcname].RUnlock()
	file.Close()

	IdlContents[idlfilepath] = string(content)
	err = IdlProviders[svcname].UpdateIDL(string(content), IdlContents)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated: " + idlfilepath)

}

func WatchIDLFiles(idlDirPath string) {
	idlWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer idlWatcher.Close()

	err = idlWatcher.Add(idlDirPath)
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

func GetIdlPath(svcName string) string {
	if path, ok := IdlPaths[svcName]; ok {
		return path
	} else {

		return ""
	}
}
