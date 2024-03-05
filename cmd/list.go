/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all running k3d clusters",
	Long:  `List all running k3d cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("list called")
		color.Green("List of available k3d cluster")
		listK3d()
	},
}

func listK3d() {
	out, err := exec.Command("k3d", "cluster", "list").Output()
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println(string(out))
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
