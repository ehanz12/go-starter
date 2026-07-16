package generator

import (
	"testing"
)

func TestNormaliseDB(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"MySQL", "mysql"},
		{"mysql", "mysql"},
		{"PostgreSQL", "postgres"},
		{"postgres", "postgres"},
		{"postgresql", "postgres"},
		{"SQLite", "sqlite"},
		{"sqlite", "sqlite"},
		{"unknown", "mysql"}, // default fallback
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := NormaliseDB(tt.input)
			if got != tt.want {
				t.Errorf("NormaliseDB(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestSupportedDatabases(t *testing.T) {
	for key, cfg := range SupportedDatabases {
		if cfg.Driver == "" {
			t.Errorf("database %q has empty Driver", key)
		}
		if cfg.Import == "" {
			t.Errorf("database %q has empty Import", key)
		}
		if cfg.Dial == "" {
			t.Errorf("database %q has empty Dial", key)
		}
	}
}
