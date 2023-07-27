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
	updateSvcPath(svcPath)

	err := filepath.Walk(idlDirPath,
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

	if provider, ok := IdlProviders[svcname]; ok {
		err := provider.UpdateIDL(string(content), IdlContents)
		if err != nil {
			fmt.Printf("%v", err)
			fmt.Println("update failed")
			return
		}
	} else {
		fmt.Println("no provider for svc " + svcname)
		return
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
				if strings.HasSuffix(strings.ToUpper(event.Name), "THRIFT") {
					tftPath := "./" + event.Name
					go UpdateIDLContents(tftPath)
				} else if strings.HasSuffix(event.Name, "svcPath") {
					path := "./" + event.Name
					updateSvcPath(path)
				}

			} else if event.Op == fsnotify.Create {
				fmt.Println(event.Name)
			} else if event.Op == fsnotify.Remove {
				fmt.Println(event.Name)
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

func updateSvcPath(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	visited := make(map[string]bool)

	for k, _ := range IdlPaths {
		visited[k] = false
	}

	for {
		line, _, err := reader.ReadLine()
		if err != nil || io.EOF == err {
			break
		}

		svcR := strings.Split(string(line), ", ")
		if len(svcR) < 2 {
			fmt.Println("not a valid line:" + string(line))
			continue
		}

		if p, ok := IdlPaths[svcR[0]]; ok {
			if p == svcR[1] {
			} else {
				IdlPaths[svcR[0]] = svcR[1]
				IdlNames[svcR[1]] = svcR[0]
				delete(IdlNames, p)
			}
		} else {
			IdlPaths[svcR[0]] = svcR[1]
			IdlNames[svcR[1]] = svcR[0]
		}

		visited[svcR[0]] = true

		var rwlock sync.RWMutex
		IdlMutexs[svcR[0]] = &rwlock
	}

	for k, v := range visited {
		if v == false {
			p := IdlPaths[k]
			delete(IdlPaths, k)
			delete(IdlNames, p)
		}
	}

	if err != nil {
		panic(err)
	}
	file.Close()

	fmt.Printf("IdlPaths: %v\n", IdlPaths)
	fmt.Printf("IdlNames: %v\n", IdlNames)
	fmt.Printf("\"-------------------------\": %v\n", "-------------------------")
}

func GetContentProvider(svcName string) (*generic.ThriftContentProvider, error) {
	idlPath, ok := IdlPaths[svcName]
	if !ok {
		return nil, fmt.Errorf("no such svc :" + svcName)
	}

	if _, ok := IdlContents[idlPath]; !ok {
		file, err := os.Open(idlPath)
		if err != nil {
			fmt.Println("no such file: "+idlPath)
		}
		content, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		IdlContents[idlPath] = string(content)

		file.Close()
	}

	return generic.NewThriftContentProvider(IdlContents[idlPath], nil)
}
