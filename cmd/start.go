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

func init() {
	startCmd.Flags().BoolVarP(&startForce, "force", "f", false, "進行中の演習を上書きして開始")
	rootCmd.AddCommand(startCmd)
}

var validName = regexp.MustCompile(`^\d{2}-[a-z0-9](?:-?[a-z0-9])*$`)

var startCmd = &cobra.Command{
	Use:   "start <name>",
	Short: "演習を開始（カレントディレクトリに展開）",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		if !validName.MatchString(name) {
			fmt.Fprintf(os.Stderr, "エラー: 演習名 %q の形式が正しくありません (例: 01-interface-design)\n", name)
			os.Exit(1)
		}

		if exercFS == nil {
			fmt.Fprintln(os.Stderr, "エラー: 演習データが見つかりません")
			os.Exit(1)
		}

		ex, err := exercise.LoadFromFS(exercFS, name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: 演習 %q が見つかりません\n", name)
			os.Exit(1)
		}

		prog, _ := progress.Load()
		if prog.InProgress != "" && !startForce {
			fmt.Fprintf(os.Stderr, "エラー: 演習 %q が進行中です（--force で上書き）\n", prog.InProgress)
			os.Exit(1)
		}

		cwd, _ := os.Getwd()
		target := cwd

		if err := os.MkdirAll(target, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "エラー: ディレクトリ作成失敗: %v\n", err)
			os.Exit(1)
		}

		if err := exercise.CopyFromFS(exercFS, name+"/starter", target); err != nil {
			fmt.Fprintf(os.Stderr, "エラー: 展開失敗: %v\n", err)
			os.Exit(1)
		}
		// verify ディレクトリが存在する場合のみコピー
		if _, err := fs.Stat(exercFS, name+"/verify"); err == nil {
			if err := exercise.CopyFromFS(exercFS, name+"/verify", target); err != nil {
				fmt.Fprintf(os.Stderr, "エラー: テスト展開失敗: %v\n", err)
				os.Exit(1)
			}
		}

		prog.InProgress = name
		prog.Save()

		fmt.Printf("✅ %s\n", colorGreen("演習「"+ex.Title+"」を開始しました"))
		fmt.Println()
		fmt.Printf("  カテゴリ: %s\n", ex.Category)
		fmt.Printf("  難易度:   %s\n", stars(ex.Difficulty))
		fmt.Println()
		fmt.Println("  次のファイルを編集してください:")
		for _, f := range ex.Files {
			fmt.Printf("    - %s\n", strings.TrimSuffix(f, ".txt"))
		}
		fmt.Println()
		fmt.Printf("  編集後: %s\n", colorCyan("go-practicum verify"))
		fmt.Printf("  ヒント:  %s\n", colorCyan("go-practicum hint"))
	},
}
