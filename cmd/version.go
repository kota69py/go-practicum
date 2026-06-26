package cmd

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var version = "dev"

func (r *Runner) newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "バージョン情報を表示",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("go-practicum %s\n", version)
			fmt.Printf("go version: %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
			if bi, ok := debug.ReadBuildInfo(); ok {
				for _, s := range bi.Settings {
					if s.Key == "vcs.revision" {
						fmt.Printf("commit: %s\n", s.Value)
						break
					}
				}
			}
		},
	}
}
