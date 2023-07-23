package idlprovider

import (
	"fmt"
	"os"
)

func AddIDL(svcname string, filecontent []byte) (err error) {
	//TODO: impl
	return
}
func ChangeIDL(svcname string, filecontent []byte) (err error) {
	idlfilepath := IdlPaths[svcname]
	if idlfilepath == "" {
		err = fmt.Errorf("no matching service file")
		return
	}

	//TODO: update file name (version or sth)

	file, err := os.OpenFile(idlfilepath, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	file.WriteAt(filecontent, 0)
	if err != nil {
		panic(err)
	}
	file.Truncate(int64(len(filecontent)))
	fileString := string(filecontent)

	IdlContents[idlfilepath] = fileString

	//TODO: should provider be inited at start or right before creating client?
	if IdlProviders[idlfilepath] != nil {
		IdlProviders[idlfilepath].UpdateIDL(fileString, IdlContents)
	}

	file.Close()

	return

}
func DeleteIDL(svcname string) (err error) {
	//TODO: impl
	return
}
