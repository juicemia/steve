package cmd

import (
	"fmt"
	"os"

	"github.com/juicemia/steve/print"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "steve",
	Short: "Steve is an opinionated site generator",
	Long: `
An opinionated Static Site Generator built with hate by juice and nobody else in Go.

No complete documentation is available,`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&print.VerboseEnabled,
		"verbose",
		"v",
		false,
		"enable verbose output",
	)
}
