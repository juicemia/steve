package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new blog post",
	Long: `Create a new blog post.

TODO: update this documentation to give more technical details.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generating new post...")

		// create git branch

		// create file
	},
}
