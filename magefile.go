//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

// mage tidy
func Tidy() error {
	return sh.RunV("go", "mod", "tidy")
}

// mage fmt
func Fmt() error {
	fmt.Println("Formatting code ...")
	return sh.RunV("gofmt", "-l", "-w", ".")
}

// mage TestBuild (create test binary with goreleaser)
func TestBuild() error {
	fmt.Println("Building test binary ...")
	return sh.RunV("goreleaser", "release", "--clean", "--snapshot", "--skip=sign")
}

// mage lint WIP
/* func Lint() error {
	fmt.Println("Running GolangCI ...")
	return sh.RunV("golangci-lint", "run", "./...")
} */