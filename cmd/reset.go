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

func init() {
	rootCmd.AddCommand(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:   "reset [name]",
	Short: "進捗をリセット",
	Long: `進捗データをリセットします。
   name を指定すると該当演習のみ、省略すると全演習をリセットします。`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prog, err := progress.Load()
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: 進捗読み込み失敗: %v\n", err)
			os.Exit(1)
		}

		if len(args) == 0 {
			if len(prog.Completed) == 0 && prog.InProgress == "" {
				fmt.Println("リセットするデータがありません。")
				return
			}
			if !confirmStdin("全ての進捗データをリセットしますか？") {
				fmt.Println("キャンセルしました。")
				return
			}
			prog.Completed = nil
			prog.InProgress = ""
			if err := prog.Save(); err != nil {
				fmt.Fprintf(os.Stderr, "エラー: 保存失敗: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("✅ " + colorGreen("全ての進捗をリセットしました"))
			return
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
			fmt.Fprintf(os.Stderr, "エラー: 演習 %q の進捗データが見つかりません\n", name)
			os.Exit(1)
		}
		if !confirmStdin(fmt.Sprintf("演習 %q の進捗をリセットしますか？", name)) {
			fmt.Println("キャンセルしました。")
			return
		}
		prog.Completed = kept
		if prog.InProgress == name {
			prog.InProgress = ""
		}
		if err := prog.Save(); err != nil {
			fmt.Fprintf(os.Stderr, "エラー: 保存失敗: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✅ " + colorGreen(fmt.Sprintf("演習 %q の進捗をリセットしました\n", name)))
	},
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
