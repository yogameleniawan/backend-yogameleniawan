package go_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println("Version:", version)
}

//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)

	if err != nil {
		t.Error(err)
	}

}