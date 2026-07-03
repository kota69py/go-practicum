# go-practicum

Go エンジニアとしての実戦スキルを鍛える CLI 演習ツール。

## Overview

**写経ではなく、自分で考えてコードを書く** ことに焦点を当てた演習集です。
各演習は「設問コード＋テスト」で構成され、自力で実装してテスト通過を目指します。

- インターフェース設計・テスト実装・エラーハンドリング・並行処理パターンなど、実務で頻出するトピックをカバー
- 各行にヒントと解答例付き（段階的に参照可能）
- 進捗は自動保存。途中で止めても再開可能

## Quick Start

```bash
# 演習一覧を表示
go-practicum list

# 演習を開始（カレントディレクトリに展開）
go-practicum start 01-interface-design

# 編集後、テストで検証
cd 01-interface-design
go-practicum verify

# ヒントを表示
go-practicum hint

# 解答例を表示
go-practicum solution
```

## コマンド一覧

| コマンド | 説明 |
|----------|------|
| `list` | 全演習一覧を表示（`--category`, `--difficulty` フィルタ対応） |
| `start <name>` | 演習を開始（カレントディレクトリにスターターコードを展開） |
| `verify` | 進行中の演習を `go test` で検証（全テスト通過で完了マーク） |
| `hint` | 現在の演習のヒントを表示 |
| `solution` | 現在の演習の解答例を表示 |
| `info <name>` | 演習の詳細（カテゴリ・難易度・トピック・ヒント）を表示 |
| `search <query>` | 演習を名前・タイトル・トピックで検索 |
| `graph` | カテゴリ別の学習マップを表示 |
| `status` | 学習進捗を表示（カテゴリ別プログレスバー） |
| `check` | 現在の演習コードを `go vet` / `gofmt` で静的解析 |
| `export [format]` | 進捗をエクスポート（`json` / `html` / `csv` / `md`、`-o` で出力先指定） |
| `reset [name]` | 進捗をリセット（name指定で個別、省略で全削除） |
| `version` | バージョン情報を表示 |
| `doctor` | 全演習データ（JSON・ファイル構成）を検証 |
| `completion [bash\|zsh\|fish\|powershell]` | シェルの補完スクリプトを生成 |

## 演習一覧（全126演習）

### architecture（4演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 103 | clean-arch | ★★★ |
| 104 | cqrs | ★★★ |
| 105 | event-driven | ★★★ |
| 113 | hexagonal-arch | ★★★ |

### concurrency（16演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 05 | concurrency | ★★★★ |
| 07 | context | ★★★★ |
| 14 | graceful-shutdown | ★★★★ |
| 16 | rate-limiting | ★★★ |
| 19 | fan-in-fan-out | ★★★★ |
| 20 | sync-mutex | ★★★ |
| 21 | worker-pool | ★★★ |
| 32 | sync-once-map | ★★★ |
| 35 | sync-atomic | ★★★ |
| 40 | pubsub | ★★★★ |
| 45 | context-values | ★★★ |
| 48 | semaphore | ★★★ |
| 50 | errgroup-singleflight | ★★★★ |
| 53 | lockfree-ringbuffer | ★★★★★ |
| 72 | rate-limiter | ★★★★ |
| 110 | worker-pool | ★★★ |

### language（16演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 13 | generics | ★★★ |
| 22 | reflection | ★★★ |
| 24 | functional-options | ★★★ |
| 27 | sorting | ★★ |
| 30 | time | ★★ |
| 31 | strings | ★★ |
| 47 | enum | ★★★ |
| 55 | unsafe-memory | ★★★★★ |
| 58 | build-tags | ★★★★ |
| 80 | codegen | ★★★★★ |
| 81 | maps-slices | ★★ |
| 82 | iter-seq | ★★★ |
| 109 | generics-constraints | ★★★ |
| 114 | go-generate | ★★ |
| 118 | go-analysis | ★★★★ |
| 120 | cgo-basics | ★★★★ |

### testing（11演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 02 | table-driven-test | ★★ |
| 10 | test-double | ★★★ |
| 18 | test-helpers | ★★★ |
| 23 | benchmarking | ★★★ |
| 28 | fuzzing | ★★★ |
| 71 | fuzzing | ★★★★ |
| 79 | mockgen | ★★★ |
| 99 | web-testing | ★★ |
| 106 | fuzz-testing | ★★ |
| 116 | testcontainers | ★★★ |
| 121 | advanced-testing | ★★★ |

### io（6演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 09 | file-io | ★★★ |
| 25 | embed | ★★ |
| 26 | io-reader-writer | ★★★ |
| 33 | io-multi-pipe | ★★★ |
| 37 | compress-gzip | ★★★ |
| 119 | io-fs | ★★★ |

### net（17演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 03 | http-handler | ★★★ |
| 12 | http-client | ★★★ |
| 15 | middleware | ★★★ |
| 41 | file-server | ★★ |
| 49 | grpc-basics | ★★★★ |
| 67 | otel-tracing | ★★★★ |
| 68 | prometheus-metrics | ★★★★ |
| 70 | health-check | ★★★ |
| 76 | grpc-streaming | ★★★★ |
| 83 | servemux-routing | ★★★ |
| 87 | grpc-interceptor | ★★★★ |
| 88 | grpc-deadline | ★★★★ |
| 89 | grpc-errors | ★★★ |
| 101 | k8s-basics | ★★★ |
| 102 | k8s-probe | ★★ |
| 115 | websocket-chat | ★★★ |
| 123 | sse | ★★★ |

### encoding（3演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 08 | json-serialization | ★★★ |
| 38 | encoding-xml | ★★ |
| 42 | struct-validation | ★★★ |

### configuration（4演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 17 | config | ★★ |
| 29 | cli-flag | ★★ |
| 77 | feature-flag | ★★★★ |
| 122 | go-work | ★★★ |

### error-handling（5演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 04 | error-handling | ★★★ |
| 39 | retry-backoff | ★★★ |
| 111 | retry-backoff | ★★★ |
| 112 | circuit-breaker | ★★★ |
| 124 | panic-recover | ★★★ |

### os（5演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 36 | os-exec | ★★★ |
| 43 | file-watch | ★★★ |
| 69 | graceful-shutdown | ★★★ |
| 100 | docker-multiarch | ★★ |
| 125 | docker-sdk | ★★★ |

### performance（8演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 51 | gc-profiling | ★★★★ |
| 52 | sync-pool-zero-alloc | ★★★★ |
| 57 | advanced-benchmark | ★★★★ |
| 61 | gmp-scheduler | ★★★★★ |
| 62 | escape-analysis | ★★★★★ |
| 64 | latency-analysis | ★★★★ |
| 65 | trace-profiling | ★★★★★ |
| 117 | pgo | ★★★ |

### security（6演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 56 | constant-time-comparison | ★★★★ |
| 59 | secure-memory | ★★★★ |
| 93 | jwt-auth | ★★★ |
| 94 | tls-mtls | ★★★★ |
| 95 | oauth2 | ★★★ |
| 96 | secret-management | ★★★ |

### design（8演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 01 | interface-design | ★★ |
| 44 | command-pattern | ★★★ |
| 54 | circuit-breaker | ★★★★ |
| 60 | hook-pattern | ★★★★ |
| 63 | connection-pool | ★★★★ |
| 73 | workflow-orchestration | ★★★★★ |
| 75 | wire-di | ★★★★ |
| 78 | api-pagination | ★★★ |

### database（2演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 06 | sql-transaction | ★★★ |
| 74 | sql-migration | ★★★ |

### datastore（3演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 90 | redis-basics | ★★★ |
| 91 | database-migration | ★★★★ |
| 92 | connection-pool | ★★★ |

### logging（5演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 11 | structured-log | ★★ |
| 66 | structured-logging | ★★★ |
| 107 | slog-basics | ★★ |
| 108 | slog-handler | ★★★ |
| 126 | slogtest | ★★ |

### templating（1演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 34 | template | ★★★ |

### crypto（1演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 46 | crypto-aes | ★★★ |

### observability（3演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 84 | trace-context | ★★★ |
| 85 | context-propagation | ★★★★ |
| 86 | red-metrics | ★★★ |

### web（2演習）

| # | 演習 | 難易度 |
|---|------|--------|
| 97 | web-router | ★★ |
| 98 | web-middleware | ★★ |

## 進捗管理

進捗は `~/.go-practicum/progress.json` に自動保存されます。

```json
{
  "completed": ["01-interface-design", "02-table-driven-test"],
  "in_progress": "03-http-handler"
}
```

- `start` で `in_progress` が設定される
- `verify` で全テスト通過時に完了リストに追加
- ファイルを直接編集すれば手動管理も可能

## 演習の構造

各演習は以下のディレクトリ構成を持ちます：

```
exercdata/
  01-interface-design/
    exercise.json     # 演習メタデータ（タイトル・カテゴリ・難易度・ヒント）
    starter/          # スターターコード（編集対象）
      main.go.txt     # .txt は Go のビルド対象外のための拡張
      go.mod.txt
      ...
    solution/         # 解答例（starter と同じファイル名で完全実装）
      ...
    verify/           # 検証用テスト（go test で実行）
      main_test.go.txt
```

## 開発

```bash
go build -o go-practicum .
# または
go install .
```

### Makefile

```bash
make all       # vet → test → build
make test      # 全テスト実行
make vet       # go vet
make lint      # golangci-lint
make doctor    # 演習データ検証
make clean     # テストキャッシュ削除
```

### 全演習のテスト

```bash
go test ./...
```
