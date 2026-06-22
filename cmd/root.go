package cmd

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

var exercFS fs.FS

func SetExercFS(fsys fs.FS) {
	exercFS = fsys
}

var rootCmd = &cobra.Command{
	Use:   "go-practicum",
	Short: "Go実戦演習 — 設計・テスト・実装力を鍛えるCLIツール",
	Long: `go-practicum は、Goエンジニアとしての実戦スキルを鍛えるCLI演習ツールです。

写経ではなく、インターフェース設計・テスト実装・エラーハンドリング・
並行処理パターンなどを「自分で考えてコードを書く」ことに焦点を当てています。

使い方:
  go-practicum list                   演習一覧を表示
  go-practicum start <name>           演習を開始（カレントディレクトリに展開）
  go-practicum verify                 演習を検証（go test を実行）
  go-practicum hint                   ヒントを表示
  go-practicum solution               解答例を表示`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
