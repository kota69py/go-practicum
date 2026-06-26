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

## 演習一覧（全76演習）

### concurrency（15演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 05 | concurrency | ★★★★ | WaitGroup, channel fan-out |
| 07 | context | ★★★★ | context.WithCancel, ctx.Done() |
| 14 | graceful-shutdown | ★★★★ | signal.Notify, http.Server.Shutdown |
| 16 | rate-limiting | ★★★ | token bucket, time.Ticker |
| 19 | fan-in-fan-out | ★★★★ | マージチャネル, ワーカー分散 |
| 20 | sync-mutex | ★★★ | Mutex, RWMutex, スレッドセーフデータ構造 |
| 21 | worker-pool | ★★★ | channel-based worker pool, context cancel |
| 32 | sync-once-map | ★★★ | sync.Once, sync.Map, CompareAndSwap |
| 35 | sync-atomic | ★★★ | atomic.Add, atomic.Value, CAS |
| 40 | pubsub | ★★★★ | チャネルベース publish/subscribe |
| 45 | context-values | ★★★ | context.WithValue, リクエストスコープデータ |
| 48 | semaphore | ★★★ | チャネルベース重み付きセマフォ |
| 50 | errgroup-singleflight | ★★★★ | errgroup, singleflight |
| 53 | lockfree-ringbuffer | ★★★★★ | lock-free SPSC, sync/atomic, memory ordering |
| 72 | rate-limiter | ★★★★ | トークンバケット, スライディングウィンドウ, レート制御 |

### language（9演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 13 | generics | ★★★ | 型パラメータ, Map/Filter/Reduce |
| 22 | reflection | ★★★ | reflect.TypeOf/ValueOf, 構造体タグ |
| 24 | functional-options | ★★★ | 関数型オプションパターン |
| 27 | sorting | ★★ | sort.Slice, sort.Search, 重複除去 |
| 30 | time | ★★ | time.Format/Parse, DaysBetween, タイマー |
| 31 | strings | ★★ | WordCount, Capitalize, strings.Builder |
| 47 | enum | ★★★ | 型安全 enum, Stringer, JSON 対応 |
| 55 | unsafe-memory | ★★★★★ | unsafe.Pointer, struct padding, メモリアライメント |
| 58 | build-tags | ★★★★ | //go:build, 条件付きビルド, プラットフォーム分岐 |

### testing（6演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 02 | table-driven-test | ★★ | テーブル駆動テスト |
| 10 | test-double | ★★★ | モック, 呼び出し記録アサート |
| 18 | test-helpers | ★★★ | t.Helper, Golden Files, テストユーティリティ |
| 23 | benchmarking | ★★★ | testing.B, 文字列連結ベンチマーク |
| 28 | fuzzing | ★★★ | Go 1.18+ fuzzing, hex parser |
| 71 | fuzzing-advanced | ★★★★ | testing/fuzz, 不変条件テスト, ラウンドトリップ検証 |

### io（5演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 09 | file-io | ★★★ | CSV リーディング, encoding/csv |
| 25 | embed | ★★ | //go:embed, embed.FS |
| 26 | io-reader-writer | ★★★ | io.Reader/Writer 実装, ROT13 |
| 33 | io-multi-pipe | ★★★ | MultiReader, TeeReader, io.Pipe |

### net（9演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 03 | http-handler | ★★★ | REST API, httptest |
| 12 | http-client | ★★★ | retry with backoff, context timeout |
| 15 | middleware | ★★★ | HTTP ミドルウェアデコレータパターン |
| 41 | file-server | ★★ | http.FileServer, 静的ファイル配信 |
| 49 | grpc-basics | ★★★★ | gRPC, protobuf, UnaryServerInterceptor |
| 67 | otel-tracing | ★★★★ | OpenTelemetry, 分散トレース, Span, 伝搬 |
| 68 | prometheus-metrics | ★★★★ | Prometheus, Counter/Histogram, /metrics |
| 70 | health-check | ★★★ | Readiness/Liveness, Kubernetes Probe, Checker |
| 76 | grpc-streaming | ★★★★ | Server/Client/Bidi Streaming, チャット |

### encoding（3演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 08 | json-serialization | ★★★ | struct tags, MarshalJSON/UnmarshalJSON |
| 38 | encoding-xml | ★★ | xml.Marshal/Unmarshal, 構造体タグ, 属性 |
| 42 | struct-validation | ★★★ | reflect 構造体タグバリデーション |

### configuration（2演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 17 | config | ★★ | 環境変数読み込み, デフォルト値 |
| 29 | cli-flag | ★★ | flag.FlagSet サブコマンド |

### error-handling（2演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 04 | error-handling | ★★★ | sentinel errors, エラーラッピング, custom error type |
| 39 | retry-backoff | ★★★ | 指数バックオフ, ジッター付きリトライ |

### os（3演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 36 | os-exec | ★★★ | exec.Command, stdin pipe, ストリーミング出力 |
| 43 | file-watch | ★★★ | ポーリングベースファイル変更検知 |
| 69 | graceful-shutdown | ★★★ | os/signal, Shutdown, シグナルハンドリング |

### performance（7演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 51 | gc-profiling | ★★★★ | runtime.ReadMemStats, pprof, GCトレース, アロケーション最適化 |
| 52 | sync-pool-zero-alloc | ★★★★ | sync.Pool, testing.B, -benchmem, バッファ再利用 |
| 57 | advanced-benchmark | ★★★★ | testing.Benchmark, 比較計測, ウォームアップ, alloc分析 |
| 61 | gmp-scheduler | ★★★★★ | G/M/P モデル, GOMAXPROCS, Gosched, プリエンプション |
| 62 | escape-analysis | ★★★★★ | エスケープ分析, スタック/ヒープ, -gcflags=-m |
| 64 | latency-analysis | ★★★★ | パーセンタイル, Tail Latency, 負荷テスト, outlier検出 |
| 65 | trace-profiling | ★★★★★ | runtime/trace, 実行トレース, goroutine可視化 |

### security（2演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 56 | constant-time-comparison | ★★★★ | crypto/subtle, HMAC, タイミング攻撃対策, 定数時間比較 |
| 59 | secure-memory | ★★★★ | crypto/rand, メモリゼロ化, パスワードハッシュ, 機密情報保護 |

### design（7演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 01 | interface-design | ★★ | Store インターフェース, 依存性注入 |
| 44 | command-pattern | ★★★ | Command インターフェース, Execute/Undo |
| 54 | circuit-breaker | ★★★★ | サーキットブレーカー, 障害分離, 自動回復 |
| 60 | hook-pattern | ★★★★ | Hook/インターセプターチェーン, パイプライン, Recovery |
| 63 | connection-pool | ★★★★ | 汎用Pool[T], 借用/返却, idle管理, リソース制御 |
| 73 | workflow-orchestration | ★★★★★ | ステートマシン, パイプライン, ロールバック |
| 75 | wire-di | ★★★★ | google/wire, 依存性注入, Provider/Injector |

### database（2演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 06 | sql-transaction | ★★★ | Begin/Commit/Rollback, モックストア |
| 74 | sql-migration | ★★★ | golang-migrate, スキーマ管理, Up/Down |

### logging（2演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 11 | structured-log | ★★ | log/slog, JSON handler, Group |
| 66 | structured-logging | ★★★ | log/slog レベル制御, 構造化ログ, With/コンテキスト |

### templating（1演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 34 | template | ★★★ | text/template, range, 条件分岐 |

### crypto（1演習）

| # | 演習 | 難易度 | トピック |
|---|------|--------|----------|
| 46 | crypto-aes | ★★★ | AES-256-GCM 暗号化・復号 |

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

### 全演習のテスト

```bash
go test ./...
```
