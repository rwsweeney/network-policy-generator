package main

import (
	"github.com/rwsweeney/network-policy-generator/cmd"
)

var (
	// version is set during build
	version = "development"
	// commit is set during build
	commit = "n/a"
)

func main() {
	cmd.Execute(version, commit)
}
