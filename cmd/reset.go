package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func (r *Runner) newResetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reset [name]",
		Short: "進捗をリセット",
		Long: `進捗データをリセットします。
   name を指定すると該当演習のみ、省略すると全演習をリセットします。`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			prog, err := progress.Load()
			if err != nil {
				return fmt.Errorf("進捗読み込み失敗: %v", err)
			}

			if len(args) == 0 {
				if len(prog.Completed) == 0 && prog.InProgress == "" {
					cmd.Println("リセットするデータがありません。")
					return nil
				}
				if !confirmStdin("全ての進捗データをリセットしますか？") {
					cmd.Println("キャンセルしました。")
					return nil
				}
				prog.Completed = nil
				prog.InProgress = ""
				if err := progress.Save(prog); err != nil {
					return fmt.Errorf("保存失敗: %v", err)
				}
				cmd.Println("✅ " + colorGreen("全ての進捗をリセットしました"))
				return nil
			}

			name := args[0]
			found := false
			var kept []string
			for _, c := range prog.Completed {
				if c == name {
					found = true
				} else {
					kept = append(kept, c)
				}
			}
			if !found && prog.InProgress != name {
				return fmt.Errorf("演習 %q の進捗データが見つかりません", name)
			}
			if !confirmStdin(fmt.Sprintf("演習 %q の進捗をリセットしますか？", name)) {
				cmd.Println("キャンセルしました。")
				return nil
			}
			prog.Completed = kept
			if prog.InProgress == name {
				prog.InProgress = ""
			}
			if err := progress.Save(prog); err != nil {
				return fmt.Errorf("保存失敗: %v", err)
			}
			cmd.Println("✅ " + colorGreen(fmt.Sprintf("演習 %q の進捗をリセットしました", name)))
			return nil
		},
	}
}

func confirmStdin(prompt string) bool {
	return confirm(prompt, os.Stdin)
}

func confirm(prompt string, r io.Reader) bool {
	fmt.Printf("%s [y/N]: ", prompt)
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		return false
	}
	return strings.EqualFold(strings.TrimSpace(scanner.Text()), "y")
}
