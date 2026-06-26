package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "進捗データをすべてリセット",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("すべての進捗データを削除します。よろしいですか？ (y/N): ")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if strings.ToLower(line) != "y" && strings.ToLower(line) != "yes" {
			fmt.Println("キャンセルしました。")
			return
		}

		prog, err := progress.Load()
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
			os.Exit(1)
		}
		prog.Completed = nil
		prog.InProgress = ""
		if err := prog.Save(); err != nil {
			fmt.Fprintf(os.Stderr, "エラー: 保存失敗: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ " + colorGreen("進捗データをリセットしました"))
	},
}
