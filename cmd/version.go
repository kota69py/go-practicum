package cmd

import (
	"runtime"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var version = "dev"

func (r *Runner) newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "バージョン情報を表示",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Printf("go-practicum %s\n", version)
			cmd.Printf("go version: %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
			if bi, ok := debug.ReadBuildInfo(); ok {
				for _, s := range bi.Settings {
					if s.Key == "vcs.revision" {
						cmd.Printf("commit: %s\n", s.Value)
						break
					}
				}
			}
			return nil
		},
	}
}
