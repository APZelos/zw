package cli

import (
	"fmt"

	"github.com/apzelos/zw/internal/shell"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <shell>",
	Short: "Output shell integration code",
	Long: `Output shell integration code for the specified shell.

Supported shells: bash, zsh, fish

Add the following to your shell configuration file:

Bash (~/.bashrc):
  eval "$(zw init bash)"

Zsh (~/.zshrc):
  eval "$(zw init zsh)"

Fish (~/.config/fish/config.fish):
  zw init fish | source`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		shellName := args[0]

		script, err := shell.GetIntegration(shellName)
		if err != nil {
			return err
		}

		fmt.Fprint(cmd.OutOrStdout(), script)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
