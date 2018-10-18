package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bft",
	Short: "bft - binary file toolbox",
	Long:  `bft allows you to manipulate binary files in a convenient way`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bft - binary file toolbox")
	},
}

func init() {
	rootCmd.AddCommand(sliceCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
