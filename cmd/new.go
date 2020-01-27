package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
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

		wd, err := os.Getwd()
		if err != nil {
			fmt.Printf("error finding current directory: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("opening git repository...")
		repo, err := git.PlainOpen(wd)
		if err != nil {
			fmt.Printf("error opening git repository: %v\n", err)
			os.Exit(1)
		}

		worktree, err := repo.Worktree()
		if err != nil {
			fmt.Printf("error getting git worktree: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("creating backing branch...")
		worktree.Checkout(&git.CheckoutOptions{
			Create: true,
			Force: false,
			Branch: plumbing.ReferenceName("refs/heads/post/new-post-test"),
		})
		if err != nil {
			fmt.Printf("error creating new branch for post: %v\n", err)
			os.Exit(1)
		}

		// create file
		fmt.Println("creating markdown file...")
		_, err = os.Create(fmt.Sprintf("%s/test-blog/blog/new-post-test.md", wd))
		if err != nil {
			fmt.Printf("error creating markdown file: %v\n", err)
			os.Exit(1)
		}
	},
}
