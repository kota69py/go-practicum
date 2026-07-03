package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strings"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

var startForce bool

var validName = regexp.MustCompile(`^\d{2,3}-[a-z0-9](?:-?[a-z0-9])*$`)

func (r *Runner) newStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start <name>",
		Short: "演習を開始（カレントディレクトリに展開）",
		Args:  cobra.ExactArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			name := args[0]

			if !validName.MatchString(name) {
				return fmt.Errorf("演習名 %q の形式が正しくありません (例: 01-interface-design)", name)
			}

			if r.exercFS == nil {
				return fmt.Errorf("演習データが見つかりません")
			}

			ex, err := exercise.LoadFromFS(r.exercFS, name)
			if err != nil {
				if hints := exercise.SuggestNames(r.exercFS, name, 3); len(hints) > 0 {
					return fmt.Errorf("演習 %q が見つかりません\n  もしかして: %s", name, strings.Join(hints, ", "))
				}
				return fmt.Errorf("演習 %q が見つかりません", name)
			}

			prog, _ := progress.Load()
			if prog.InProgress != "" && !startForce {
				return fmt.Errorf("演習 %q が進行中です（--force で上書き）", prog.InProgress)
			}

			cwd, _ := os.Getwd()
			target := cwd

			if err := os.MkdirAll(target, 0755); err != nil {
				return fmt.Errorf("ディレクトリ作成失敗: %v", err)
			}

			if err := exercise.CopyFromFS(r.exercFS, name+"/starter", target); err != nil {
				return fmt.Errorf("展開失敗: %v", err)
			}
			// verify ディレクトリが存在する場合のみコピー
			if _, err := fs.Stat(r.exercFS, name+"/verify"); err == nil {
				if err := exercise.CopyFromFS(r.exercFS, name+"/verify", target); err != nil {
					return fmt.Errorf("テスト展開失敗: %v", err)
				}
			}

			prog.InProgress = name
			_ = progress.Save(prog)

			c.Printf("✅ %s\n", colorGreen("演習「"+ex.Title+"」を開始しました"))
			c.Println()
			c.Printf("  カテゴリ: %s\n", ex.Category)
			c.Printf("  難易度:   %s\n", stars(ex.Difficulty))
			c.Println()
			c.Println("  次のファイルを編集してください:")
			for _, f := range ex.Files {
				c.Printf("    - %s\n", strings.TrimSuffix(f, ".txt"))
			}
			c.Println()
			c.Printf("  編集後: %s\n", colorCyan("go-practicum verify"))
			c.Printf("  ヒント:  %s\n", colorCyan("go-practicum hint"))
			return nil
		},
	}
	cmd.Flags().BoolVarP(&startForce, "force", "f", false, "進行中の演習を上書きして開始")
	return cmd
}
