package idlprovider

import (
	idlmanager "apigateway/biz/model/idlmanager"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func AddIDL(req idlmanager.AddServiceReq) (err error) {

	idlfilepath := idlDir + "/" + req.FileName
	var newLock sync.RWMutex
	idlMutexs[idlfilepath] = &newLock

	idlMutexs[idlfilepath].Lock()
	file, err := os.Create(idlfilepath)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(req.FileContent)
	if err != nil {
		panic(err)
	}
	file.Close()

	IdlContents[idlfilepath] = ""
	idlMutexs[idlfilepath].Unlock()

	// add to svcR
	svcMutex.Lock()
	file, err = os.OpenFile(svcRelation, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	svcRelationItem := req.Name + ", " + idlfilepath + "\n"

	file.Write([]byte(svcRelationItem))
	file.Sync()

	if err != nil {
		panic(err)
	}
	file.Close()
	svcMutex.Unlock()

	IdlPaths[req.Name] = idlfilepath
	IdlNames[idlfilepath] = req.Name
	return
}
func ChangeIDL(req idlmanager.ChangeServiceReq) (err error) {

	idlfilepath := IdlPaths[req.Name]
	if idlfilepath == "" {
		err = fmt.Errorf("no matching service file")
		return
	}

	idlMutexs[req.Name].Lock()
	file, err := os.OpenFile(idlfilepath, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	file.WriteAt(req.FileContent, 0)
	if err != nil {
		panic(err)
	}
	file.Truncate(int64(len(req.FileContent)))

	file.Close()
	idlMutexs[req.Name].Unlock()

	return
}
func DeleteIDL(req idlmanager.DeleteServiceReq) (err error) {

	svcMutex.Lock()
	file, err := os.OpenFile(svcRelation, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	pos := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil || io.EOF == err {
			break
		}

		svcR := strings.Split(string(line), ", ")

		if svcR[0] == req.Name {
			// replace with empty bytes
			_, err = file.WriteAt(make([]byte, len(line)), int64(pos))
			if err != nil {
				panic(err)
			}

			err = file.Sync()
			pos = -1
			break
		}

		pos += 1 + len(line)
	}

	if err != nil {
		panic(err)
	}
	file.Close()
	svcMutex.Unlock()

	if pos != -1 {
		err = fmt.Errorf("relation not present")
		return
	}

	delete(idlMutexs, req.Name)
	delete(IdlProviders, req.Name)
	delete(IdlContents, IdlPaths[req.Name])
	delete(IdlNames, IdlPaths[req.Name])
	delete(IdlPaths, req.Name)

	return
}
