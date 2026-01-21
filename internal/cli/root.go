package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	queryFlag string
)

var rootCmd = &cobra.Command{
	Use:   "zw",
	Short: "A git worktree navigation tool",
	Long: `zw is a CLI tool for navigating between git worktrees.

It provides fuzzy matching and interactive selection to quickly
switch between worktrees in your repository.

To enable shell integration, add the following to your shell config:
  eval "$(zw init bash)"   # for bash
  eval "$(zw init zsh)"    # for zsh
  zw init fish | source    # for fish`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if queryFlag != "" {
			return runQuery(cmd, queryFlag)
		}

		// If no arguments and no flags, show help
		if len(args) == 0 {
			return cmd.Help()
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().StringVar(&queryFlag, "query", "", "Query for worktree and output path (for shell integration)")
}

func Execute() error {
	return rootCmd.Execute()
}

func runQuery(cmd *cobra.Command, pattern string) error {
	// TODO: Implement worktree lookup and fuzzy matching
	// For now, just print a placeholder message to stderr
	// and nothing to stdout (so the shell wrapper won't cd)
	_, _ = fmt.Fprintln(cmd.ErrOrStderr(), "zw: worktree lookup not yet implemented")
	return nil
}
