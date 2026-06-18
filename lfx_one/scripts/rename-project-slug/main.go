// Copyright The Linux Foundation and each contributor to LFX.
// SPDX-License-Identifier: MIT

// rename-project-slug renames a project slug across LFX V2 data stores:
// OpenSearch (resources index) and NATS JetStream KV buckets.
//
// Usage:
//
//	go run . <old-slug> <new-slug> [flags]
//
// Examples:
//
//	# Preview changes (default) against both stores
//	go run . gridfm opengridfm
//
//	# Apply changes to OpenSearch only
//	OPENSEARCH_URL=http://... go run . gridfm opengridfm --target=opensearch --dry-run=false
//
//	# Apply changes to NATS KV only
//	NATS_URL=nats://... go run . gridfm opengridfm --target=nats --dry-run=false --concurrency=20
//
//	# Apply changes to both stores
//	go run . gridfm opengridfm --dry-run=false
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/url"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	opensearchgo "github.com/opensearch-project/opensearch-go/v2"
	"golang.org/x/sync/errgroup"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var errSlugMismatch = errors.New("project_slug does not match old slug")

// bucketSlugFields maps a KV bucket name to the JSON field name(s) that hold
// the project slug. Only the field listed here is rewritten — other fields
// named "slug" that refer to a different entity (e.g. an org slug) are left
// alone. Buckets absent from the map default to {"project_slug"}.
var bucketSlugFields = map[string][]string{
	"committee-members":  {"project_slug"},
	"committees":         {"project_slug"},
	"committee-settings": {"project_slug"},
	// The "projects" bucket stores the project entity whose *own* slug is the
	// field being renamed. Confirm json tag in the real project-service before
	// running --dry-run=false against this bucket.
	"projects":         {"slug"},
	"project-settings": {"project_slug"},
}

var (
	targetFlag    = flag.String("target", "both", "which stores to migrate: opensearch, nats, or both")
	dryRun        = flag.Bool("dry-run", true, "preview changes without writing (pass --dry-run=false to apply)")
	debug         = flag.Bool("debug", false, "enable debug logging")
	natsURL       = flag.String("nats-url", getEnvOrDefault("NATS_URL", nats.DefaultURL), "NATS server URL")
	natsBuckets   = flag.String("nats-buckets", "committee-members,committees,committee-settings,projects,project-settings", "comma-separated NATS KV bucket names to migrate")
	concurrency   = flag.Int("concurrency", 50, "max concurrent NATS KV record updates per bucket")
	opensearchURL = flag.String("opensearch-url", getEnvOrDefault("OPENSEARCH_URL", "http://localhost:9200"), "OpenSearch base URL")
	oldSlugFlag   = flag.String("old-slug", "", "current slug (alternative to first positional arg)")
	newSlugFlag   = flag.String("new-slug", "", "new slug (alternative to second positional arg)")
)

type bucketStats struct {
	Bucket  string
	Total   int
	Updated int
	Skipped int
	Failed  int
}

func main() {
	flag.Parse()

	level := slog.LevelInfo
	if *debug {
		level = slog.LevelDebug
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})))

	if err := run(flag.Args()); err != nil {
		log.Fatalf("rename-project-slug failed: %v", err)
	}
}

func run(slugArgs []string) error {
	oldSlug := strings.TrimSpace(*oldSlugFlag)
	newSlug := strings.TrimSpace(*newSlugFlag)

	hasFlagSlugs := oldSlug != "" || newSlug != ""
	hasPosArgs := len(slugArgs) > 0

	if hasFlagSlugs && hasPosArgs {
		return fmt.Errorf("provide slugs either as positional args OR via --old-slug/--new-slug flags, not both")
	}
	if hasPosArgs {
		if len(slugArgs) != 2 {
			return fmt.Errorf("expected exactly 2 positional args (<old-slug> <new-slug>), got %d", len(slugArgs))
		}
		oldSlug = strings.TrimSpace(slugArgs[0])
		newSlug = strings.TrimSpace(slugArgs[1])
	}
	if oldSlug == "" || newSlug == "" {
		return fmt.Errorf("usage: go run . <old-slug> <new-slug> [flags]\n       or use --old-slug and --new-slug flags")
	}
	if oldSlug == newSlug {
		return fmt.Errorf("old-slug and new-slug must differ")
	}
	if *concurrency < 1 {
		return fmt.Errorf("--concurrency must be at least 1")
	}

	target := strings.ToLower(strings.TrimSpace(*targetFlag))
	if target != "both" && target != "opensearch" && target != "nats" {
		return fmt.Errorf("--target must be one of: opensearch, nats, both")
	}

	ctx := context.Background()

	slog.InfoContext(ctx, "Starting rename-project-slug",
		"old_slug", oldSlug,
		"new_slug", newSlug,
		"target", target,
		"dry_run", *dryRun,
	)

	if target == "opensearch" || target == "both" {
		if err := runOpenSearch(ctx, oldSlug, newSlug); err != nil {
			return fmt.Errorf("opensearch migration failed: %w", err)
		}
	}

	if target == "nats" || target == "both" {
		bucketList := parseBuckets(*natsBuckets)
		if err := runNATS(ctx, oldSlug, newSlug, bucketList); err != nil {
			return fmt.Errorf("nats migration failed: %w", err)
		}
	}

	return nil
}

// ── OpenSearch ────────────────────────────────────────────────────────────────

func runOpenSearch(ctx context.Context, oldSlug, newSlug string) error {
	cfg := opensearchgo.Config{
		Addresses: []string{*opensearchURL},
	}

	client, err := opensearchgo.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("failed to create OpenSearch client: %w", err)
	}

	slog.InfoContext(ctx, "OpenSearch migration",
		"url", redactURL(*opensearchURL),
		"index", "resources",
		"dry_run", *dryRun,
	)

	// Build the bool query shared between audit and update_by_query.
	query := buildOSQuery(oldSlug)

	if *dryRun {
		return osAudit(ctx, client, query, oldSlug)
	}
	return osUpdateByQuery(ctx, client, query, oldSlug, newSlug)
}

func buildOSQuery(oldSlug string) map[string]any {
	return map[string]any{
		"bool": map[string]any{
			"filter": []any{
				map[string]any{"term": map[string]any{"latest": true}},
			},
			"should": []any{
				map[string]any{"term": map[string]any{"data.project_slug": oldSlug}},
				map[string]any{"term": map[string]any{"data.slug": oldSlug}},
				map[string]any{"term": map[string]any{"tags": "project_slug:" + oldSlug}},
				map[string]any{"term": map[string]any{"object_ref": "project:" + oldSlug}},
				map[string]any{"term": map[string]any{"parent_refs": "project:" + oldSlug}},
			},
			"minimum_should_match": 1,
		},
	}
}

func osAudit(ctx context.Context, client *opensearchgo.Client, query map[string]any, oldSlug string) error {
	body, err := jsonBody(map[string]any{
		"size":             0,
		"track_total_hits": true,
		"query":            query,
	})
	if err != nil {
		return err
	}

	res, err := client.Search(
		client.Search.WithContext(ctx),
		client.Search.WithIndex("resources"),
		client.Search.WithBody(body),
	)
	if err != nil {
		return fmt.Errorf("search request failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		raw, _ := io.ReadAll(res.Body)
		return fmt.Errorf("search error %s: %s", res.Status(), raw)
	}

	var result struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode search response: %w", err)
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("OpenSearch Audit (dry-run)")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Slug:               %s\n", oldSlug)
	fmt.Printf("Index:              %s\n", "resources")
	fmt.Printf("Records matched:    %d\n", result.Hits.Total.Value)
	fmt.Println("(pass --dry-run=false to apply the update)")
	fmt.Println(strings.Repeat("=", 50))

	return nil
}

func osUpdateByQuery(ctx context.Context, client *opensearchgo.Client, query map[string]any, oldSlug, newSlug string) error {
	// Painless script is reproduced verbatim from LFXV2-2254 comment.
	painlessSource := `
def oldSlug=params.oldSlug;
def newSlug=params.newSlug;
boolean changed=false;
def data=ctx._source.get('data');
if (data instanceof Map) {
  if (oldSlug.equals(data.get('project_slug'))) { data.put('project_slug', newSlug); changed=true; }
  if (oldSlug.equals(data.get('slug'))) { data.put('slug', newSlug); changed=true; }
}
def tags=ctx._source.get('tags');
if (tags instanceof List) {
  for (int i=0; i<tags.size(); i++) {
    String tag=(String) tags.get(i);
    if (('project_slug:'+oldSlug).equals(tag)) { tags.set(i, 'project_slug:'+newSlug); changed=true; }
  }
}
String objectRef=(String) ctx._source.get('object_ref');
if (objectRef!=null && ('project:'+oldSlug).equals(objectRef)) { ctx._source.put('object_ref', 'project:'+newSlug); changed=true; }
def parentRefs=ctx._source.get('parent_refs');
if (parentRefs instanceof List) {
  for (int i=0; i<parentRefs.size(); i++) {
    String ref=(String) parentRefs.get(i);
    if (('project:'+oldSlug).equals(ref)) { parentRefs.set(i, 'project:'+newSlug); changed=true; }
  }
}
String ft=(String) ctx._source.get('fulltext');
if (ft!=null && ft.contains(oldSlug)) { ctx._source.put('fulltext', ft.replace(oldSlug, newSlug)); changed=true; }
def aliases=ctx._source.get('name_and_aliases');
if (aliases instanceof List) {
  for (int i=0; i<aliases.size(); i++) {
    String alias=(String) aliases.get(i);
    if (oldSlug.equals(alias)) { aliases.set(i, newSlug); changed=true; }
  }
}
if (!changed) { ctx.op='noop'; }
`

	body, err := jsonBody(map[string]any{
		"conflicts": "proceed",
		"query":     query,
		"script": map[string]any{
			"lang":   "painless",
			"source": strings.TrimSpace(painlessSource),
			"params": map[string]any{
				"oldSlug": oldSlug,
				"newSlug": newSlug,
			},
		},
	})
	if err != nil {
		return err
	}

	res, err := client.UpdateByQuery(
		[]string{"resources"},
		client.UpdateByQuery.WithContext(ctx),
		client.UpdateByQuery.WithBody(body),
	)
	if err != nil {
		return fmt.Errorf("update_by_query request failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		raw, _ := io.ReadAll(res.Body)
		return fmt.Errorf("update_by_query error %s: %s", res.Status(), raw)
	}

	var result struct {
		Total            int `json:"total"`
		Updated          int `json:"updated"`
		VersionConflicts int `json:"version_conflicts"`
		Noops            int `json:"noops"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode update_by_query response: %w", err)
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("OpenSearch Update Complete")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Index:              %s\n", "resources")
	fmt.Printf("Total examined:     %d\n", result.Total)
	fmt.Printf("Updated:            %d\n", result.Updated)
	fmt.Printf("Noops:              %d\n", result.Noops)
	fmt.Printf("Version conflicts:  %d\n", result.VersionConflicts)
	fmt.Println(strings.Repeat("=", 50))

	return nil
}

// ── NATS KV ───────────────────────────────────────────────────────────────────

func runNATS(ctx context.Context, oldSlug, newSlug string, buckets []string) error {
	nc, err := nats.Connect(*natsURL,
		nats.Timeout(10*time.Second),
		nats.MaxReconnects(3),
		nats.ReconnectWait(2*time.Second),
	)
	if err != nil {
		return fmt.Errorf("failed to connect to NATS at %s: %w", redactNATSURL(*natsURL), err)
	}
	defer nc.Close()

	slog.InfoContext(ctx, "Connected to NATS",
		"url", redactNATSURL(nc.ConnectedUrl()),
		"buckets", buckets,
		"dry_run", *dryRun,
	)

	js, err := jetstream.New(nc)
	if err != nil {
		return fmt.Errorf("failed to create JetStream context: %w", err)
	}

	var grandTotal, grandUpdated, grandSkipped, grandFailed, bucketErrors int
	for _, bucket := range buckets {
		stats, err := migrateBucket(ctx, js, bucket, oldSlug, newSlug)
		if err != nil {
			if errors.Is(err, jetstream.ErrBucketNotFound) {
				slog.WarnContext(ctx, "bucket not found, skipping", "bucket", bucket)
			} else {
				slog.ErrorContext(ctx, "bucket migration failed", "bucket", bucket, "error", err)
				bucketErrors++
			}
			continue
		}
		grandTotal += stats.Total
		grandUpdated += stats.Updated
		grandSkipped += stats.Skipped
		grandFailed += stats.Failed
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	if *dryRun {
		fmt.Println("NATS KV Audit (dry-run, no writes)")
	} else {
		fmt.Println("NATS KV Update Complete (all buckets)")
	}
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Total records:    %d\n", grandTotal)
	if *dryRun {
		fmt.Printf("Would update:     %d\n", grandUpdated)
	} else {
		fmt.Printf("Updated:          %d\n", grandUpdated)
	}
	fmt.Printf("Skipped:          %d\n", grandSkipped)
	fmt.Printf("Failed:           %d\n", grandFailed)
	fmt.Printf("Bucket errors:    %d\n", bucketErrors)
	fmt.Println(strings.Repeat("=", 50))

	if bucketErrors > 0 {
		return fmt.Errorf("%d bucket(s) failed to open or list — migration incomplete", bucketErrors)
	}
	if grandFailed > 0 {
		return fmt.Errorf("%d records failed to update across all buckets", grandFailed)
	}
	return nil
}

func migrateBucket(ctx context.Context, js jetstream.JetStream, bucket, oldSlug, newSlug string) (*bucketStats, error) {
	kvStore, err := js.KeyValue(ctx, bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to open KV bucket %q: %w", bucket, err)
	}

	fields := bucketFieldsFor(bucket)

	slog.InfoContext(ctx, "Scanning bucket", "bucket", bucket, "slug_fields", fields)

	keys, err := kvStore.ListKeys(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list keys in bucket %q: %w", bucket, err)
	}
	defer keys.Stop() //nolint:errcheck

	var recordKeys []string
	for key := range keys.Keys() {
		if strings.HasPrefix(key, "lookup/") || strings.HasPrefix(key, "slug/") {
			continue
		}
		recordKeys = append(recordKeys, key)
	}

	slog.InfoContext(ctx, "Found records in bucket", "bucket", bucket, "count", len(recordKeys))

	if *dryRun {
		slog.InfoContext(ctx, "DRY RUN MODE - no writes will be made", "bucket", bucket)
	}

	stats := &bucketStats{Bucket: bucket, Total: len(recordKeys)}
	var statsMu sync.Mutex
	var processed atomic.Int64

	g, gCtx := errgroup.WithContext(ctx)
	g.SetLimit(*concurrency)

	for _, key := range recordKeys {
		key := key
		g.Go(func() error {
			err := processKVRecord(gCtx, kvStore, key, fields, oldSlug, newSlug, *dryRun)

			statsMu.Lock()
			if err != nil {
				if errors.Is(err, errSlugMismatch) {
					stats.Skipped++
				} else {
					slog.ErrorContext(gCtx, "failed to process record",
						"bucket", bucket, "key", key, "error", err)
					stats.Failed++
				}
			} else {
				stats.Updated++
			}
			statsMu.Unlock()

			if n := processed.Add(1); n%1000 == 0 || int(n) == stats.Total {
				statsMu.Lock()
				u, sk, f := stats.Updated, stats.Skipped, stats.Failed
				statsMu.Unlock()
				slog.InfoContext(gCtx, "Progress",
					"bucket", bucket,
					"processed", n, "total", stats.Total,
					"updated", u, "skipped", sk, "failed", f,
				)
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return stats, err
	}

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Printf("Bucket: %s\n", bucket)
	fmt.Printf("  Total:        %d\n", stats.Total)
	if *dryRun {
		fmt.Printf("  Would update: %d\n", stats.Updated)
	} else {
		fmt.Printf("  Updated:      %d\n", stats.Updated)
	}
	fmt.Printf("  Skipped:      %d\n", stats.Skipped)
	fmt.Printf("  Failed:       %d\n", stats.Failed)

	return stats, nil
}

func processKVRecord(ctx context.Context, kvStore jetstream.KeyValue, key string, fields []string, oldSlug, newSlug string, dryRun bool) error {
	entry, err := kvStore.Get(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to get entry: %w", err)
	}

	raw := make(map[string]json.RawMessage)
	if err := json.Unmarshal(entry.Value(), &raw); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	matched := false
	for _, field := range fields {
		val, ok := raw[field]
		if !ok {
			continue
		}
		var s string
		if err := json.Unmarshal(val, &s); err != nil {
			continue
		}
		if s == oldSlug {
			matched = true
			break
		}
	}

	if !matched {
		slog.DebugContext(ctx, "no matching slug field, skipping", "key", key)
		return errSlugMismatch
	}

	slog.DebugContext(ctx, "updating record slug fields",
		"key", key,
		"fields", fields,
		"old_slug", oldSlug,
		"new_slug", newSlug,
		"dry_run", dryRun,
	)

	if dryRun {
		return nil
	}

	// Apply writes with optimistic-lock retry (3 attempts).
	maxRetries := 3
	var updateErr error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		// Re-read the latest entry on each attempt to get the current revision.
		if attempt > 1 {
			entry, err = kvStore.Get(ctx, key)
			if err != nil {
				return fmt.Errorf("failed to re-fetch entry: %w", err)
			}
			raw = make(map[string]json.RawMessage)
			if err := json.Unmarshal(entry.Value(), &raw); err != nil {
				return fmt.Errorf("failed to unmarshal re-fetched entry: %w", err)
			}
			// Abort if another process already made the change.
			anyMatch := false
			for _, field := range fields {
				if val, ok := raw[field]; ok {
					var s string
					if json.Unmarshal(val, &s) == nil && s == oldSlug {
						anyMatch = true
						break
					}
				}
			}
			if !anyMatch {
				slog.DebugContext(ctx, "slug changed by concurrent process, skipping", "key", key)
				return errSlugMismatch
			}
		}

		newSlugJSON, _ := json.Marshal(newSlug)
		for _, field := range fields {
			if val, ok := raw[field]; ok {
				var s string
				if json.Unmarshal(val, &s) == nil && s == oldSlug {
					raw[field] = newSlugJSON
				}
			}
		}

		// Refresh updated_at if present.
		if _, ok := raw["updated_at"]; ok {
			ts, _ := json.Marshal(time.Now().UTC().Format(time.RFC3339Nano))
			raw["updated_at"] = ts
		}

		updated, marshalErr := json.Marshal(raw)
		if marshalErr != nil {
			return fmt.Errorf("failed to marshal updated record: %w", marshalErr)
		}

		_, updateErr = kvStore.Update(ctx, key, updated, entry.Revision())
		if updateErr == nil {
			break
		}

		if attempt < maxRetries {
			slog.WarnContext(ctx, "optimistic lock failed, retrying",
				"key", key, "attempt", attempt, "error", updateErr)
			time.Sleep(time.Duration(attempt*100) * time.Millisecond)
		}
	}

	if updateErr != nil {
		return fmt.Errorf("failed to update after %d attempts: %w", maxRetries, updateErr)
	}

	return nil
}

// bucketFieldsFor returns the JSON field names that hold a project slug for the
// given bucket name. Falls back to {"project_slug"} for unknown buckets.
func bucketFieldsFor(bucket string) []string {
	if fields, ok := bucketSlugFields[bucket]; ok {
		return fields
	}
	return []string{"project_slug"}
}

// ── helpers ───────────────────────────────────────────────────────────────────

func jsonBody(v any) (io.Reader, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}
	return bytes.NewReader(b), nil
}

func parseBuckets(s string) []string {
	var out []string
	for _, b := range strings.Split(s, ",") {
		b = strings.TrimSpace(b)
		if b != "" {
			out = append(out, b)
		}
	}
	return out
}

func redactURL(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return "<invalid>"
	}
	if u.User != nil {
		u.User = url.User("REDACTED")
	}
	return u.String()
}

func redactNATSURL(raw string) string { return redactURL(raw) }

func getEnvOrDefault(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
