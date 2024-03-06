/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize k3d configuration",
	Long: `Initialize k3d configuration file under .kube/k3d-config.

If the configuration file already exists at this location,
the command will show a message and stop.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		initConfig()
	},
}

func initConfig() (string, error) {
	// Check if file, exists, if yes fail with error message
	//Home, _ := os.UserHomeDir()
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path = filepath.Join(path, ".kube/k3d-config")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), err
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("File already exists, if you really want to overwrite, use --force")
		} else {
			color.Green("OK, we're good to go")
	}
}


func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
