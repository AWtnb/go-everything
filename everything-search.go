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
	everything.Walk(query, func(path string, info everything.FileInfo, err error) error {
		name := filepath.Base(path)
		size := info.Size()
		tmod := info.ModTime()
		fmt.Println(name, size, tmod.Format("2006-01-02 15:04"))
		return nil
	})
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
