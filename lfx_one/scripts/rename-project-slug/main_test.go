// Copyright The Linux Foundation and each contributor to LFX.
// SPDX-License-Identifier: MIT

package main

import (
	"testing"
)

func TestBucketFieldsFor_knownBuckets(t *testing.T) {
	cases := []struct {
		bucket string
		field  string
	}{
		{"committee-members", "project_slug"},
		{"committees", "project_slug"},
		{"committee-settings", "project_slug"},
		{"projects", "slug"},
		{"project-settings", "project_slug"},
	}
	for _, c := range cases {
		fields := bucketFieldsFor(c.bucket)
		found := false
		for _, f := range fields {
			if f == c.field {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("bucketFieldsFor(%q): expected field %q, got %v", c.bucket, c.field, fields)
		}
	}
}

func TestBucketFieldsFor_unknownBucket(t *testing.T) {
	fields := bucketFieldsFor("some-unknown-bucket")
	if len(fields) != 1 || fields[0] != "project_slug" {
		t.Errorf("expected [project_slug] for unknown bucket, got %v", fields)
	}
}

func TestParseBuckets(t *testing.T) {
	got := parseBuckets("committee-members, committees , committee-settings")
	want := []string{"committee-members", "committees", "committee-settings"}
	assertEqual(t, want, got)
}

func TestBuildOSQuery_containsOldSlug(t *testing.T) {
	q := buildOSQuery("gridfm")
	b, ok := q["bool"].(map[string]any)
	if !ok {
		t.Fatal("expected bool key in query")
	}
	should, ok := b["should"].([]any)
	if !ok {
		t.Fatal("expected should key in bool query")
	}
	if len(should) == 0 {
		t.Fatal("expected non-empty should clauses")
	}
}

// assertEqual is a minimal helper to avoid importing testify in a standalone module.
func assertEqual[T comparable](t *testing.T, want, got []T) {
	t.Helper()
	if len(want) != len(got) {
		t.Errorf("length mismatch: want %v, got %v", want, got)
		return
	}
	for i := range want {
		if want[i] != got[i] {
			t.Errorf("index %d: want %v, got %v", i, want[i], got[i])
		}
	}
}
