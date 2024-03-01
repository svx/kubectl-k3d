//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

func Tidy() error {
	return sh.Run("go", "mod", "tidy")
}

// mage fmt
func Fmt() error {
	fmt.Println("Formatting code ...")
	return sh.Run("gofmt", "-l", "-w", ".")
}
