# Contributing

## 開発の流れ

1. Issue で提案・議論
2. Fork → Feature Branch
3. 実装 + テスト
4. Pull Request

## 動作確認

```bash
make all          # vet → test → build
make lint         # golangci-lint
go run . doctor   # 演習の整合性チェック
```

## 演習の追加

1. `internal/exercdata/exercdata/NN-name/` を作成
2. `exercise.json` にメタデータを記述
3. `starter/` にひな形、`solution/` に解答例、`verify/` にテストを配置
4. カテゴリは `internal/categories/categories.go` の Known 一覧から選ぶ
5. 新しいカテゴリを追加する場合は Known に追加する

## コミットメッセージ

- `feat:` 新機能・新演習
- `fix:` バグ修正
- `chore:` 依存関係・CI・ドキュメント
- `refactor:` リファクタリング
