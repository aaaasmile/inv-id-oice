package util

import (
	"crypto/rand"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/kardianos/osext"
)

var rootPath string
var UseRelativeRoot = true

func GetFullPath(relPath string) string {
	if UseRelativeRoot {
		return relPath
	}

	if rootPath == "" {
		var err error
		rootPath, err = osext.ExecutableFolder()
		if err != nil {
			log.Fatalf("ExecutableFolder failed: %v", err)
		}
		log.Println("Executable folder (rootdir) is ", rootPath)
		//rootPath, _ = filepath.Split(os.Args[0]) // os.Args[0] can be "faked". (https://github.com/kardianos/osext)
	}
	r := filepath.Join(rootPath, relPath)
	return r
}

func PseudoUuid() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}

func GetLastPathInUri(uri string) (string, string) {
	arr := strings.Split(uri, "/")
	for i := len(arr) - 1; i >= 0; i-- {
		last := arr[i]
		rem_ix := i
		if last != "" {
			if !strings.HasPrefix(last, "?") {
				remPath := strings.Join(arr[0:rem_ix], "/")
				return last, remPath
			}
		}
	}
	return uri, ""
}
