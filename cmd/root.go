package cmd

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

type Runner struct {
	exercFS fs.FS
	rootCmd *cobra.Command
}

func NewRunner(fsys fs.FS) *Runner {
	r := &Runner{exercFS: fsys}
	r.rootCmd = &cobra.Command{
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
   go-practicum solution               解答例を表示
   go-practicum check                  コードを静的解析
   go-practicum export [json|html]     学習進捗をエクスポート`,
	}
	r.rootCmd.AddCommand(r.newListCmd())
	r.rootCmd.AddCommand(r.newStartCmd())
	r.rootCmd.AddCommand(r.newVerifyCmd())
	r.rootCmd.AddCommand(r.newHintCmd())
	r.rootCmd.AddCommand(r.newSolutionCmd())
	r.rootCmd.AddCommand(r.newInfoCmd())
	r.rootCmd.AddCommand(r.newSearchCmd())
	r.rootCmd.AddCommand(r.newGraphCmd())
	r.rootCmd.AddCommand(r.newStatusCmd())
	r.rootCmd.AddCommand(r.newCheckCmd())
	r.rootCmd.AddCommand(r.newExportCmd())
	r.rootCmd.AddCommand(r.newResetCmd())
	r.rootCmd.AddCommand(r.newVersionCmd())
	return r
}

func (r *Runner) Execute() {
	if err := r.rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
