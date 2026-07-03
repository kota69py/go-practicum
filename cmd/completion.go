package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func (r *Runner) newCompletionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "シェルの補完スクリプトを生成",
		Long: `指定されたシェルの補完スクリプトを標準出力に出力します。

使い方:
  go-practicum completion bash > /etc/bash_completion.d/go-practicum
  go-practicum completion zsh > /usr/local/share/zsh/site-functions/_go-practicum
  go-practicum completion fish > ~/.config/fish/completions/go-practicum.fish
  go-practicum completion powershell > go-practicum.ps1`,
		Args:      cobra.MaximumNArgs(1),
		ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
		RunE: func(cmd *cobra.Command, args []string) error {
			shell := "bash"
			if len(args) > 0 {
				shell = args[0]
			}
			switch shell {
			case "bash":
				return r.rootCmd.GenBashCompletion(os.Stdout)
			case "zsh":
				return r.rootCmd.GenZshCompletion(os.Stdout)
			case "fish":
				return r.rootCmd.GenFishCompletion(os.Stdout, true)
			case "powershell":
				return r.rootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
			default:
				return fmt.Errorf("未対応のシェル %q (bash / zsh / fish / powershell)", shell)
			}
		},
	}
}
