# go-practicum 技術領域ギャップ分析

> 現状 **128演習** のカバレッジを評価し、不足領域を洗い出す。

---

## 凡例

| 記号 | 意味 |
|------|------|
| ✅ カバー済み | 専用の演習が存在する |
| 🔶 部分的 | 関連演習はあるが深掘り不足 |
| ❌ 未カバー | 該当する演習がない |

---

## 1. Go 1.21–1.25 新機能

| 機能 | 登場 | 状況 | 該当演習 |
|------|------|------|----------|
| `iter.Seq` / range-over-func | 1.23 | ✅ | 82-iter-seq |
| `maps`, `slices` パッケージ | 1.21 | ✅ | 81-maps-slices |
| `cmp.Ordered` / `cmp.Compare` | 1.21 | ✅ | 109-generics-constraints |
| `net/http.ServeMux` methodパターン | 1.22 | ✅ | 83-servemux-routing |
| `log/slog` | 1.21 | ✅ | 107-slog-basics, 108-slog-handler, 126-slogtest |
| PGO (Profile Guided Optimization) | 1.21 | ✅ | 117-pgo |
| `unique.Make` (interning) | 1.23 | ❌ | — |
| `math/rand/v2` | 1.22 | ❌ | — |
| `testing.TestDeps` / testing.Short 高度 | 1.23 | ❌ | — |
| `go test -fuzz` のCI継続的実行 | 1.18 | 🔶 | 28/71/106 は学習用、CI未導入 |

---

## 2. Observability（可観測性）— 8演習

| 領域 | 状況 | 該当演習 |
|------|------|----------|
| 構造化ログ (slog) | ✅ | 66-structured-logging, 107-slog-basics, 108-slog-handler, 126-slogtest |
| OpenTelemetry トレーシング | ✅ | 67-otel-tracing, 84-trace-context |
| OTel SDK初期化 (バッチ/サンプリング/OTLP) | ✅ | **128-otel-sdk-init** |
| Context伝搬 (HTTP/gRPC/slog 連携) | 🔶 | 85-context-propagation, 87-grpc-interceptor |
| Prometheus メトリクス | ✅ | 68-prometheus-metrics |
| REDメトリクス設計 | ✅ | 86-red-metrics |
| ヘルスチェック | ✅ | 70-health-check, 102-k8s-probe |
| OTel + slog の1リクエスト連携 | 🔶 | 個別演習はあるが統合演習なし |

---

## 3. gRPC 深堀 — 5演習

| 領域 | 状況 | 該当演習 |
|------|------|----------|
| gRPC 基礎 (unary) | ✅ | 49-grpc-basics |
| gRPC ストリーミング | ✅ | 76-grpc-streaming |
| インターセプター (認証/ログ/トレース) | ✅ | 87-grpc-interceptor |
| デッドライン伝搬 | ✅ | 88-grpc-deadline |
| gRPC エラーハンドリング | ✅ | 89-grpc-errors |
| ヘルスプロトコル (`grpc.health.v1`) | ❌ | — |
| ロードバランシング (xds) | ❌ | — |
| `protobuf` 深堀 (`protojson`, `protowire`, `protoreflect`) | ❌ | — |

---

## 4. データベース / ストレージ — 7演習

| 領域 | 状況 | 該当演習 |
|------|------|----------|
| SQL トランザクション | ✅ | 06-sql-transaction |
| SQL マイグレーション | ✅ | 74-sql-migration, 91-database-migration |
| コネクションプール | ✅ | 63-connection-pool, 92-connection-pool |
| Redis 基礎 | ✅ | 90-redis-basics |
| テストコンテナ | ✅ | 116-testcontainers |
| Redis パイプライン / Lua / Pub/Sub | 🔶 | 90-redis-basics は基礎のみ |
| GCS / S3 互換ストレージ (presigned URL) | ❌ | — |
| 楽観的ロック / 悲観的ロック | ❌ | — |

---

## 5. セキュリティ — 8演習

| 領域 | 状況 | 該当演習 |
|------|------|----------|
| 定数時間比較 | ✅ | 56-constant-time-comparison |
| セキュアメモリ | ✅ | 59-secure-memory |
| JWT 発行・検証 | ✅ | 93-jwt-auth |
| TLS / mTLS | ✅ | 94-tls-mtls |
| OAuth2 認可コードフロー | ✅ | 95-oauth2 |
| シークレット管理 | ✅ | 96-secret-management |
| HTTP セキュリティ (CSRF/CORS/CSP) | ✅ | **127-http-security** |
| OIDC id_token 検証 | ❌ | OAuth2 はあるが OIDC は未カバー |
| Let's Encrypt / autocert | ❌ | — |
| SBOM / 依存関係スキャン | 🔶 | CI に govulncheck 導入済みだが演習なし |

---

## 6. アーキテクチャパターン — 7演習

| 領域 | 状況 | 該当演習 |
|------|------|----------|
| クリーンアーキテクチャ | ✅ | 103-clean-arch |
| CQRS | ✅ | 104-cqrs |
| イベント駆動 | ✅ | 105-event-driven |
| ヘキサゴナルアーキテクチャ | ✅ | 113-hexagonal-arch |
| Command パターン | ✅ | 44-command-pattern |
| Hook パターン | ✅ | 60-hook-pattern |
| Singleflight (Thundering herd対策) | ✅ | 50-errgroup-singleflight |
| Wire DI | ✅ | 75-wire-di |
| Circuit breaker | ✅ | 54-circuit-breaker, 112-circuit-breaker |
| Saga パターン | ❌ | — |
| Backpressure 制御 | ❌ | — |
| Graceful degradation | ❌ | — |

---

## 7. ツーリング

| ツール | 状況 | 該当演習・設定 |
|--------|------|----------------|
| `go test -fuzz` | ✅ | 28-fuzzing, 71-fuzzing, 106-fuzz-testing |
| `go tool pprof` / `go tool trace` | ✅ | 65-trace-profiling |
| `go test -bench` + benchstat | 🔶 | 57-advanced-benchmark, CI未導入 |
| `go vet` + golangci-lint (13 linters) | ✅ | `.golangci.yml` に gosec/gocritic 等含む |
| Docker multi-stage build | ✅ | `Dockerfile` (alpine→scratch) |
| `goreleaser` | ✅ | `.goreleaser.yaml` + release workflow |
| CI カバレッジ計測 | ✅ | `codecov` Action + README バッジ |
| `govulncheck` | ✅ | CI security job |
| Dependabot | ✅ | `.github/dependabot.yml` |
| `go test -fuzz` の CI 継続的実行 | ❌ | — |
| `go-licenses` ライセンスチェック | ❌ | — |
| カスタム go vet アナライザ | ❌ | — |

---

## 8. HTTP 深堀 — 13演習

| 領域 | 状況 | 該当演習 |
|------|------|----------|
| HTTP ハンドラ基礎 | ✅ | 03-http-handler |
| HTTP クライアント | ✅ | 12-http-client |
| ミドルウェアパターン | ✅ | 15-middleware |
| ServeMux ルーティング | ✅ | 83-servemux-routing |
| Web ルーター | ✅ | 97-web-router |
| Web ミドルウェア | ✅ | 98-web-middleware |
| Web テスト | ✅ | 99-web-testing |
| WebSocket | ✅ | 115-websocket-chat |
| Server-Sent Events | ✅ | 123-sse |
| HTTP セキュリティ (CSRF/CORS/CSP) | ✅ | **127-http-security** |
| レート制限 | ✅ | 16-rate-limiting |
| `httputil.ReverseProxy` | ❌ | — |
| HTTP/2 サーバプッシュ | ❌ | — |
| TLS 終端 + ACME (`autocert`) | ❌ | — |

---

## 総評

128演習まで拡充した結果、**主要8領域のカバレッジは大きく向上**した。
特にセキュリティ（8演習）、Observability（8演習）、gRPC（5演習）、HTTP（13演習）は実戦レベルの厚みが出ている。

### それでも不足している領域

| 領域 | ギャップ |
|------|----------|
| **分散システム** | Saga、Backpressure、Graceful degradation |
| **ストレージ応用** | GCS/S3 presigned URL、Redis 高度（Pipeline/Lua）、楽観的/悲観的ロック |
| **OIDC** | OAuth2 はあるが id_token 検証は未カバー |
| **CI 高度化** | Fuzzing 継続実行、Benchmark 比較、ライセンスチェック |
| **プロトコルバッファ** | protojson/protowire/protoreflect の深掘り |

### 次の一手

もし追加するとすれば:

```
distributed-systems  (Saga / Backpressure / Graceful degradation)  → 3演習
storage-advanced     (GCS/S3 / Redis pipeline / locking)           → 3演習
ci-advanced          (fuzz CI / benchstat / license check)         → 1設定 + 2演習
oidc-deep            (id_token verification, PKCE)                 → 1演習
protobuf-deep        (protojson, protowire, protoreflect)         → 1演習
```
