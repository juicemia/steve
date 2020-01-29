package cmd

import (
	"fmt"
	"os"
	"errors"

	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"github.com/juicemia/steve/print"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new FILENAME",
	Short: "Create a new blog post",
	Long: `Create a new blog post.

TODO: update this documentation to give more technical details.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing parameter for file name")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		print.Verboseln("generating new post...")

		wd, err := os.Getwd()
		if err != nil {
			print.Fatalf("error finding current directory: %v\n", err)
		}

		print.Verboseln("opening git repository...")
		repo, err := git.PlainOpen(wd)
		if err != nil {
			print.Fatalf("error opening git repository: %v\n", err)
		}

		worktree, err := repo.Worktree()
		if err != nil {
			print.Fatalf("error getting git worktree: %v\n", err)
		}

		print.Verboseln("creating backing branch...")
		worktree.Checkout(&git.CheckoutOptions{
			Create: true,
			Force: false,
			Branch: plumbing.ReferenceName("refs/heads/post/new-post-test"),
		})
		if err != nil {
			print.Fatalf("error creating new branch for post: %v\n", err)
		}

		print.Verboseln("creating markdown file...")
		path := fmt.Sprintf("%s/%s.md", wd, args[0])
		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			print.Fatalf("error creating markdown file: %v\n", err)
		}

		fmt.Fprintln(file, "# Title")
	},
}
