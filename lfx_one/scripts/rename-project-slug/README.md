# rename-project-slug

A standalone Go migration script that renames a project slug across the two LFX V2
data stores that do not self-heal during a Core Services slug migration:

- **OpenSearch** — the `resources` index (dry-run audits the match count, apply
  runs a painless `_update_by_query` that rewrites every slug occurrence).
- **NATS JetStream KV** — the source-of-truth buckets for V2 services (dry-run
  logs what would change, apply writes each updated record with optimistic-lock
  retries).

## Prerequisites

- Go 1.24+
- Network access (direct or via `kubectl port-forward`) to OpenSearch and NATS

## Build

```bash
go build -o bin/rename-project-slug .
```

Or run directly without building:

```bash
go run . <old-slug> <new-slug> [flags]
```

## Usage

```text
go run . <old-slug> <new-slug> [flags]
```

### Required arguments

| Argument | Description |
|----------|-------------|
| `<old-slug>` | The current project slug to replace |
| `<new-slug>` | The new project slug to set |

Alternatively, use `--old-slug` / `--new-slug` flags.

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--target` | `both` | Stores to migrate: `opensearch`, `nats`, or `both` |
| `--dry-run` | `true` | Preview changes without writing; pass `--dry-run=false` to apply |
| `--debug` | `false` | Enable debug-level structured logging |

#### OpenSearch flags

| Flag / Env var | Default | Description |
|----------------|---------|-------------|
| `--opensearch-url` / `OPENSEARCH_URL` | `http://localhost:9200` | OpenSearch base URL |
| `--opensearch-index` / `OPENSEARCH_INDEX` | `resources` | Index name |
| `--opensearch-username` / `OPENSEARCH_USERNAME` | _(none)_ | Basic-auth username |
| `--opensearch-password` / `OPENSEARCH_PASSWORD` | _(none)_ | Basic-auth password |
| `--insecure-skip-tls-verify` | `false` | Skip TLS verification (useful for port-forwarded clusters) |

#### NATS flags

| Flag / Env var | Default | Description |
|----------------|---------|-------------|
| `--nats-url` / `NATS_URL` | `nats://localhost:4222` | NATS server URL |
| `--nats-buckets` | see below | Comma-separated KV bucket names to migrate |
| `--concurrency` | `10` | Max concurrent record updates per bucket |

Default NATS buckets: `committee-members,committees,committee-settings,projects,project-settings`

> **Note on the `projects` bucket:** the real project-service is not always locally
> available. Before running `--dry-run=false` with this bucket, verify the JSON slug
> field name in the project-service. If the bucket does not exist the script logs a
> warning and continues.

## Examples

```bash
# Audit (dry-run, both stores) — safe to run any time
go run . gridfm opengridfm

# Audit OpenSearch only
OPENSEARCH_URL=http://localhost:9200 go run . gridfm opengridfm --target=opensearch

# Apply OpenSearch changes
OPENSEARCH_URL=http://localhost:9200 \
  go run . gridfm opengridfm --target=opensearch --dry-run=false

# Apply NATS KV changes
NATS_URL=nats://localhost:4222 \
  go run . gridfm opengridfm --target=nats --dry-run=false --concurrency=20

# Apply both stores
OPENSEARCH_URL=http://localhost:9200 \
NATS_URL=nats://localhost:4222 \
  go run . gridfm opengridfm --dry-run=false

# Restrict NATS migration to specific buckets
go run . gridfm opengridfm \
  --target=nats --dry-run=false \
  --nats-buckets=committee-members,committees
```

## Port-forwarding (typical staging workflow)

```bash
# Terminal 1 — OpenSearch
kubectl port-forward -n lfx-v2 svc/opensearch 9200:9200

# Terminal 2 — NATS
kubectl port-forward -n lfx-v2 svc/nats 4222:4222

# Terminal 3 — run the script
OPENSEARCH_URL=http://localhost:9200 \
NATS_URL=nats://localhost:4222 \
  go run . <old-slug> <new-slug> --dry-run=false
```

## What the OpenSearch migration rewrites

For every document in the `resources` index that matches on any of these fields:

- `data.project_slug` equals `<old-slug>`
- `data.slug` equals `<old-slug>`
- `tags` contains `project_slug:<old-slug>`
- `object_ref` equals `project:<old-slug>`
- `parent_refs` contains `project:<old-slug>`

The painless script rewrites all matched fields to `<new-slug>` and marks the
document as a noop if nothing changed (`conflicts: proceed`).

## What the NATS KV migration rewrites

Per-bucket field mapping (which JSON field holds the project slug):

| Bucket | Field rewritten |
|--------|----------------|
| `committee-members` | `project_slug` |
| `committees` | `project_slug` |
| `committee-settings` | `project_slug` |
| `projects` | `slug` |
| `project-settings` | `project_slug` |
| _(any other bucket)_ | `project_slug` (default) |

Index/alias keys (prefixed `lookup/` or `slug/`) are skipped. Each record update
uses an optimistic-lock retry (3 attempts with backoff).

## Running tests

```bash
go test ./...
```
