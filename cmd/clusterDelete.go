/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"errors"
	"os"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
	"github.com/spf13/cobra"
)

// clusterDeleteCmd represents the clusterDelete command
var clusterDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clusterDelete called")
		deletePrompt()
	},
}

func CheckErr(err error) {
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
}

func deletePrompt() {
	//val1, err := prompt.New().Ask("Name of cluster:").Input("Blah blah")
	//CheckErr(err)

	val2, err := prompt.New().Ask("Cluster name:").Input(
		"Blah blah",
		input.WithHelp(true),
	)
	CheckErr(err)

	//fmt.Printf("{ %s }, { %s }\n", val1, val2)
	fmt.Printf("{ %s }\n", val2)
}

func init() {
	clusterCmd.AddCommand(clusterDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
