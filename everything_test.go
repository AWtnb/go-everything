package everything_test

import (
	"testing"

	"github.com/AWtnb/go-everything"
)

func TestEverything(t *testing.T) {
	found, err := everything.Scan(`C:\Personal\gotemp`, false)
	if err != nil {
		t.Error(err)
		return
	}
	for _, s := range found {
		t.Logf("'%s' was found", s)
	}
}

func TestEverythingWithWhitespace(t *testing.T) {
	found, err := everything.Scan(`C:\Personal\go temp`, false)
	if err != nil {
		t.Error(err)
		return
	}
	for _, s := range found {
		t.Logf("'%s' was found", s)
	}
}
