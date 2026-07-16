package generator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTemplatesFS(t *testing.T) {
	requiredTemplates := []string{
		"templates/main.go.tpl",
		"templates/env.tpl",
		"templates/gitignore.tpl",
		"templates/database.go.tpl",
		"templates/route.go.tpl",
		"templates/config.go.tpl",
		"templates/handlers.go.tpl",
		"templates/jwt_middleware.go.tpl",
		"templates/logger.go.tpl",
		"templates/Dockerfile.tpl",
		"templates/docker-compose.yml.tpl",
		"templates/Makefile.tpl",
		"templates/README.md.tpl",
	}

	for _, tpl := range requiredTemplates {
		t.Run(tpl, func(t *testing.T) {
			_, err := templatesFS.ReadFile(tpl)
			if err != nil {
				t.Errorf("failed to read embedded template %q: %v", tpl, err)
			}
		})
	}
}

func TestGenerateFile(t *testing.T) {
	// Create a temporary file path
	tmpDir, err := os.MkdirTemp("", "go-starter-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	outPath := filepath.Join(tmpDir, "test_main.go")
	replacements := map[string]string{
		"PROJECT_NAME": "test-project",
		"MODULE_NAME":  "github.com/test/test-project",
	}

	// Generate a file using one of the templates
	err = GenerateFile("templates/main.go.tpl", outPath, replacements)
	if err != nil {
		t.Fatalf("GenerateFile failed: %v", err)
	}

	// Verify the file was written
	content, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	contentStr := string(content)
	if len(contentStr) == 0 {
		t.Error("generated file is empty")
	}
}
