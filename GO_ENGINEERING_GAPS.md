# go-practicum 技術領域ギャップ分析

> Google トップのGoエンジニア目線で、現状の80演習に不足している技術領域をまとめる。

---

## 1. Go 1.21–1.25 新機能

現行演習バージョン (Go 1.23) の標準機能が未活用。

| 機能 | 登場 | なぜ重要 |
|------|------|---------|
| `iter.Seq` / range-over-func | 1.23 | カスタムイテレータ、DB結果セットの遅延評価 |
| `unique.Make` (interning) | 1.23 | メモリ最適化の実戦パターン |
| `maps`, `slices` パッケージ | 1.21 | 標準化されたコレクション操作 |
| `cmp.Ordered` / `cmp.Compare` | 1.21 | ジェネリクスを使った比較 |
| `math/rand/v2` | 1.22 | 高速・安全な乱数 |
| `net/http.ServeMux` methodパターン | 1.22 | `GET /api/users/{id}` の標準記法 |
| `testing.TestDeps` / testing.Short 高度 | 1.23 | テスト基盤の制御 |

---

## 2. Observability（可観測性）

現状: 3演習 (`otel-tracing`, `prometheus-metrics`, `structured-logging`)。以下が不足。

| 領域 | 内容 |
|------|------|
| OTel SDK初期化 | `OTEL_SDK_DISABLED` / バッチプロセッサ / サンプリング設定 |
| Context伝搬 | gRPC interceptor + HTTP middleware + slog を1リクエストで連携 |
| REDメトリクス設計 | Rate / Errors / Duration を自前実装 |
| ヘルスチェック拡充 | readiness / liveness / startup probe の実戦運用 |

Google SRE の第一原則は「観測可能であること」。

---

## 3. gRPC 深堀

現状: 2演習 (`grpc-basics`, `grpc-streaming`)。以下が不足。

- インターセプターチェーン（auth + logging + rate-limit + tracing の重ね合わせ）
- ヘルスプロトコル（`grpc.health.v1.Health`）
- ロードバランシング（`xds` 名前解決、pick_first / round_robin）
- デッドライン伝搬（クライアント→サーバ→下流サービス）
- `google.golang.org/protobuf` の深い使い方（`protojson`, `protowire`, `reflect/protoreflect`）

---

## 4. データベース / ストレージ

現状: 2演習 (`sql-transaction`, `sql-migration`)。以下が不足。

- NoSQL（Redis パイプライン / Luaスクリプト / Pub/Sub）
- GCS / S3 互換ストレージ（presigned URL、`cloud.google.com/go/storage`）
- コネクションプールチューニング（`MaxOpenConns`, `MaxIdleConns`, `ConnMaxLifetime`）
- 楽観的ロック / 悲観的ロック（`SELECT FOR UPDATE`）
- テストコンテナ（`testcontainers-go` を使ったインテグレーションテスト）

---

## 5. セキュリティ

現状: 2演習 (`constant-time-comparison`, `secure-memory`)。深刻な不足。

| 領域 | 内容 |
|------|------|
| OAuth2 / OIDC | `golang.org/x/oauth2`, id_token 検証 |
| JWT | `golang-jwt/jwt/v5` による発行と検証 |
| TLS / mTLS | `crypto/tls` の設定、Let's Encrypt 取得 |
| Secret管理 | Secret Manager / Vault / env からの適切な読み込み |
| HTTPセキュリティ | CSRF / CORS / CSP ヘッダー設定 |
| 依存関係 | `go.sum` / SBOM / 脆弱性スキャン |

---

## 6. アーキテクチャパターン

現状: 基本的なデザインパターンはある。以下が不足。

- CQRS / イベントソーシング
- Saga パターン（分散トランザクション）
- Backpressure（`time.Ticker` 以外の制御）
- Graceful degradation（部分障害時の設計）
- Thundering herd 対策（singleflight はあるが、実戦的な coalsecing）

---

## 7. ツーリング

現状使えていない重要なツール。

| ツール | 用途 |
|--------|------|
| `go test -fuzz` + CI | ファジングによる回帰テスト |
| `go tool pprof` / `go tool trace` | 本番プロファイルの読み方まで含む |
| `go tool cover` + HTML | CIでのカバレッジ閾値 enforcement |
| `go vet` カスタムアナライザ | `go-critic` 等のサードパーティ linter |
| Docker multi-stage build | `CGO_ENABLED=0`, `distroless`, `scratch` |
| `goreleaser` | リリース自動化、Homebrew tap |

---

## 8. HTTP 深堀

現状: 3演習 (`http-handler`, `http-client`, `middleware`)。以下が不足。

- `httputil.ReverseProxy`（リバースプロキシ実装）
- HTTP/2 サーバプッシュ / ストリーミング
- Server-Sent Events（`text/event-stream`）
- レート制限の実戦（スライディングウィンドウ + Redis）
- TLS 終端 + ACME（`autocert` / Let's Encrypt 自動取得）

---

## 総評

現状の80演習は「Go言語そのものの習得」としては良くできている。Google エンジニアとして不足しているのは主に **Observability × 本番運用 × セキュリティ** の3軸。

「動くコードを書く」から「本番で動かし続けるコードを書く」への飛躍に必要なのは、標準ライブラリ外のエコシステム知識と分散システム設計判断力。

### 追加推奨カテゴリ

```
infra         (Docker/k8s/Graceful shutdown with lifecycle hooks)
security      (OAuth2/JWT/TLS/mTLS)
observability (OTel SDK/Context propagation/RED)
datastore     (Redis/GCS/connection pooling)
integration   (testcontainers/contract testing)
```

この5領域に各3〜5演習を追加するとカバレッジが大幅に向上する。
