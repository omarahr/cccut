package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cccut",
	Short: "cccut cuts out the selected portions from each line in a file.",
	Long:  `cccut cuts out the selected portions from each line in a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func init() {
	// TODO
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
