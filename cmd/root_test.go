package cmd

import (
	"testing"
)

func TestResolveModuleName(t *testing.T) {
	t.Run("uses default module name when none is provided", func(t *testing.T) {
		got := resolveModuleName("my-api", "")
		want := "github.com/yourusername/my-api"
		if got != want {
			t.Fatalf("resolveModuleName() = %q, want %q", got, want)
		}
	})

	t.Run("uses explicit module name when provided", func(t *testing.T) {
		got := resolveModuleName("my-api", "example.com/my-api")
		want := "example.com/my-api"
		if got != want {
			t.Fatalf("resolveModuleName() = %q, want %q", got, want)
		}
	})
}

func TestValidateProjectName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid lowercase", "my-api", false},
		{"valid with dot", "my.api", false},
		{"valid with underscore", "my_api", false},
		{"valid with numbers", "api2024", false},
		{"rejects spaces", "my api", true},
		{"rejects slash", "my/api", true},
		{"rejects empty", "", true},
		{"rejects special chars", "my@api!", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateProjectName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateProjectName(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}
