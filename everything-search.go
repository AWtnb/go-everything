// https://github.com/jof4002/Everything/blob/master/_Example/walk/example.go

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AWtnb/go-win-everything/everything"
)

func main() {
	var (
		query string
	)
	flag.StringVar(&query, "query", "*go", "search query")
	flag.Parse()
	if err := checkDll("Everything64.dll"); err != nil {
		fmt.Println(err)
		return
	}
	paths := everything.Scan(query)
	for _, p := range paths {
		fmt.Println(p)
	}
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
