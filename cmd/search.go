// cmd/search.go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search through tagged responses",
	Long:  `Search through tagged responses in your local file or database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Search command is not yet implemented.")
	},
}
