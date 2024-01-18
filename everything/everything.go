// https://github.com/jof4002/Everything/blob/master/_Example/walk/example.go
package everything

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AWtnb/go-win-everything/everything/core"
)

func Scan(query string) []string {
	sl := []string{}
	if err := checkDll("Everything64.dll"); err != nil {
		fmt.Println(err)
		return sl
	}
	core.Walk(query, func(path string, info core.FileInfo, err error) error {
		sl = append(sl, path)
		return nil
	})
	return sl
}

func getExeDir() string {
	if exePath, err := os.Executable(); err != nil {
		return exePath
	}
	return ""
}

func checkDll(name string) error {
	exeDir := getExeDir()
	path := filepath.Join(exeDir, name)
	_, err := os.Stat(path)
	return err
}
