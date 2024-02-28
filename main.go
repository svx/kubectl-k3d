/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"time"
	"kubectl-k3d/cmd"

	"github.com/carlmjohnson/versioninfo"
)

func main() {
	cmd.SetVersionInfo(versioninfo.Version, versioninfo.Revision, versioninfo.LastCommit.Format(time.RFC3339))
	cmd.Execute()
}
