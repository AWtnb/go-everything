// https://github.com/jof4002/Everything/blob/master/_Example/walk/example.go

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/AWtnb/go-win-everything/everything"
)

func main() {
	var (
		query    string
		skipFile bool
		exclude  string
	)
	flag.StringVar(&query, "query", "*go", "search query")
	flag.BoolVar(&skipFile, "skipFile", false, "switch to skip file")
	flag.StringVar(&exclude, "exclude", "", "search exception (comma-separated)")
	flag.Parse()
	var res EverythingResult
	found, err := everything.Scan(query, skipFile)
	if err != nil {
		return
	}
	res.paths = found
	res.exclude = toSlice(exclude, ",")
	for _, p := range res.filtered() {
		fmt.Println(p)
	}
}

type EverythingResult struct {
	paths   []string
	exclude []string
}

func (er EverythingResult) isException(i int) bool {
	p := er.paths[i]
	sep := string(os.PathSeparator)
	for _, s := range er.exclude {
		if strings.Contains(p, sep+s+sep) || strings.HasSuffix(p, s) {
			return true
		}
	}
	return false
}

func (er EverythingResult) filtered() []string {
	if len(er.exclude) < 1 {
		return er.paths
	}
	sl := []string{}
	for i := 0; i < len(er.paths); i++ {
		p := er.paths[i]
		if er.isException(i) {
			continue
		}
		sl = append(sl, p)
	}
	return sl
}

func toSlice(s string, sep string) []string {
	var ss []string
	if len(s) < 1 {
		return ss
	}
	for _, elem := range strings.Split(s, sep) {
		ss = append(ss, strings.TrimSpace(elem))
	}
	return ss
}
