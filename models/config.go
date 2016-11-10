package models

import (
	"os"
	"path/filepath"
)

var (
	defaultRoleDir string
	Gopath         = os.Getenv("GOPATH")
)

func init() {
	// good enough
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	defaultRoleDir = dir
}
