---
name: lfx-v2-backend
description: >
  Expert guide for developing on the LFX Self-Service backend platform (historically called
  LFX v2 and LFX One — these names still appear throughout the codebase and repos).
  Use this skill whenever the user is working in any lfx-v2-* repository, asking about
  LFX Self-Service / LFX v2 / LFX One backend services, adding new fields or resources
  to the platform, making backend changes to support UI updates, or debugging platform
  service issues. Trigger for non-technical users making data model changes, frontend
  engineers touching the backend, and experienced backend engineers alike. Trigger on
  any mention of: lfx-v2, LFX One, LFX Self-Service, project-service, committee-service,
  meeting-service, voting-service, survey-service, indexer-service, query-service,
  fga-sync, access-check, NATS indexing, OpenFGA tuples, OpenSearch indexing,
  lfx-v2-helm, lfx-v2-argocd, or any LFX v2 resource type.
---

# LFX Self-Service Backend Development Guide

> **Naming note**: The platform is now called **LFX Self-Service**. It was previously
> called **LFX v2** and then **LFX One**. All three names appear throughout the codebase,
> repos (`lfx-v2-*`), and documentation — they all refer to the same platform.

The LFX Self-Service platform is a microservices backend built in Go. It uses fine-grained
access control (OpenFGA), event-driven messaging (NATS), and full-text search (OpenSearch)
as its core infrastructure. All services live in `lfx-v2-*` GitHub repositories.

## What to Read for Your Task

| Task | Reference |
| --- | --- |
| Repo map, deployment overview, local dev setup, where to start | [references/getting-started.md](references/getting-started.md) |
| NATS subject naming, service-to-service communication, KV storage | [references/nats-messaging.md](references/nats-messaging.md) |
| Goa DSL conventions, `make apigen`, adding fields, ETag / If-Match pattern | [references/goa-patterns.md](references/goa-patterns.md) |
| Indexer message payload, `IndexingConfig` schema, OpenSearch doc structure | [references/indexer-patterns.md](references/indexer-patterns.md) |
| OpenFGA tuples, generic fga-sync handlers, permission inheritance, debugging access | [references/fga-patterns.md](references/fga-patterns.md) |
| Native vs wrapper service types, which template to follow | [references/service-types.md](references/service-types.md) |
| Query service API, how it queries OpenSearch and checks access via fga-sync | [references/query-service.md](references/query-service.md) |
| Service Helm chart — deployment, HTTPRoute, Heimdall ruleset, KV buckets, secrets | [references/helm-chart.md](references/helm-chart.md) |
| Building a new resource service end-to-end | [references/new-service.md](references/new-service.md) |

---

## Platform Services

These four services are the infrastructure layer. They are generic — you rarely change them
when adding a new resource type.

| Service | Repo | Role |
| --- | --- | --- |
| **indexer-service** | `lfx-v2-indexer-service` | Receives NATS events, writes documents to OpenSearch |
| **query-service** | `lfx-v2-query-service` | HTTP search API over OpenSearch with access control filtering |
| **fga-sync** | `lfx-v2-fga-sync` | Syncs access tuples to OpenFGA; handles access check requests |
| **access-check** | `lfx-v2-access-check` | HTTP wrapper for access checks (validates JWT, calls fga-sync) |

## Resource Services

These own a specific resource and expose a REST API. There are two types — see
[references/service-types.md](references/service-types.md) for the full breakdown.

| Type | Examples | Pattern |
| --- | --- | --- |
| **Native** | `project-service`, `committee-service` | Owns data in NATS JetStream KV |
| **Wrapper** | `meeting-service`, `voting-service`, `survey-service` | Proxies to external system |

---

## The Two-Message Rule

Every write operation (create, update, delete) **must** publish an index message.
If the resource type has its own OpenFGA type, it must also publish an access message:

1. **Index message** (`lfx.index.{type}`) → indexer-service → updates OpenSearch
2. **Access message** (`lfx.fga-sync.update_access`) → fga-sync → updates OpenFGA permissions
   (only if the resource has an FGA type)

See [references/nats-messaging.md](references/nats-messaging.md) for subject naming and
[references/fga-patterns.md](references/fga-patterns.md) for the access message format.

---

## Adding a Field to an Existing Resource

The most common task. Steps at a glance:

1. Add the field to the domain model in `internal/domain/model/`
2. Add the field attribute to the Goa design in `cmd/{service}/design/`, then run `make apigen` — see [references/goa-patterns.md](references/goa-patterns.md)
3. Wire it through the payload-to-domain and domain-to-response converter functions
4. Ensure the field is in the `data` payload of the NATS index message
5. If it should be searchable, update `IndexingConfig` — see [references/indexer-patterns.md](references/indexer-patterns.md)

No OpenSearch schema migration needed — `data` is a `flat_object` and handles new fields dynamically.

---

## Code Conventions

- **Commit titles**: `[JIRA-TICKET] type: description` — e.g. `[LFXV2-1228] fix: send ActionUpdated for update events`
- **Actions**: Use `ActionCreated`, `ActionUpdated`, `ActionDeleted` constants — never hardcode strings
- **Errors**: Typed errors from `pkg/errors` — e.g. `errors.NewBadRequest(...)`, `errors.NewNotFound(...)`
- **Context**: Pass `context.Context` through all calls; never store it in a struct
- **Logging**: Structured JSON via `slog`; `DebugContext` for success, `WarnContext` for recoverable issues
- **Tests**: Colocate `_test.go` with source; `testify` for assertions; mocks in `internal/infrastructure/mock/`
- **Generated code**: Never edit files in `gen/` — overwritten by `make apigen`
- **OpenTelemetry**: New services must include the full OTEL stack (tracing, metrics, logs via `slog-otel`)

---

## Local Development

See [local-development.md](../local-development.md) for environment setup with OrbStack and Helm.

Common environment variables:

```bash
NATS_URL=nats://localhost:4222
OPENSEARCH_URL=http://localhost:9200
JWKS_URL=http://localhost:4457/.well-known/jwks
LFX_ENVIRONMENT=lfx.
PORT=8080
```
