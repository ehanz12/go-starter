package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ProjectSpec holds all parameters needed to scaffold a project.
type ProjectSpec struct {
	Name     string
	Module   string
	DB       string   // normalised driver key: mysql | postgres | sqlite
	Features []string // feature names from prompt
	InitGit  bool
}

// CreateProject scaffolds the entire project based on the given spec.
func CreateProject(spec ProjectSpec) error {
	dbCfg, ok := SupportedDatabases[spec.DB]
	if !ok {
		dbCfg = SupportedDatabases["mysql"]
	}

	// ---- 1. Create directories -------------------------------------------
	folders := []string{
		spec.Name,
		filepath.Join(spec.Name, "cmd", "api"),
		filepath.Join(spec.Name, "config"),
		filepath.Join(spec.Name, "internal"),
		filepath.Join(spec.Name, "internal", "handlers"),
		filepath.Join(spec.Name, "internal", "middleware"),
		filepath.Join(spec.Name, "internal", "models"),
		filepath.Join(spec.Name, "internal", "repositories"),
		filepath.Join(spec.Name, "internal", "requests"),
		filepath.Join(spec.Name, "internal", "responses"),
		filepath.Join(spec.Name, "internal", "routes"),
		filepath.Join(spec.Name, "internal", "services"),
		filepath.Join(spec.Name, "internal", "utils"),
		filepath.Join(spec.Name, "internal", "validators"),
		filepath.Join(spec.Name, "database", "migrations"),
		filepath.Join(spec.Name, "docs"),
		filepath.Join(spec.Name, "logs"),
		filepath.Join(spec.Name, "storage"),
		filepath.Join(spec.Name, "scripts"),
		filepath.Join(spec.Name, "pkg"),
	}

	for _, folder := range folders {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			return fmt.Errorf("mkdir %q: %w", folder, err)
		}
		fmt.Println(" 📁", folder)
	}

	// ---- 2. Build replacement map ----------------------------------------
	replacements := map[string]string{
		"PROJECT_NAME": spec.Name,
		"MODULE_NAME":  spec.Module,
		"DB_DRIVER":    dbCfg.Driver,
		"DB_PORT":      dbCfg.Port,
		"DB_IMPORT":    dbCfg.Import,
		"DB_DIAL":      dbCfg.Dial,
		// Docker Compose placeholders
		"DB_IMAGE":       dbCfg.DockerImage,
		"DB_ENV":         dbCfg.DockerEnv,
		"DB_DATA_PATH":   dbCfg.DockerData,
		"DB_HEALTHCHECK": dbCfg.HealthCheck,
	}

	// ---- 3. Core files (always generated) --------------------------------
	coreFiles := []struct{ tpl, out string }{
		{"templates/main.go.tpl", filepath.Join(spec.Name, "cmd", "api", "main.go")},
		{"templates/env.tpl", filepath.Join(spec.Name, ".env")},
		{"templates/env.tpl", filepath.Join(spec.Name, ".env.example")},
		{"templates/gitignore.tpl", filepath.Join(spec.Name, ".gitignore")},
		{"templates/database.go.tpl", filepath.Join(spec.Name, "database", "database.go")},
		{"templates/route.go.tpl", filepath.Join(spec.Name, "internal", "routes", "route.go")},
		{"templates/config.go.tpl", filepath.Join(spec.Name, "config", "config.go")},
		{"templates/handlers.go.tpl", filepath.Join(spec.Name, "internal", "handlers", "handlers.go")},
	}

	for _, f := range coreFiles {
		if err := GenerateFile(f.tpl, f.out, replacements); err != nil {
			return fmt.Errorf("generate %q: %w", f.out, err)
		}
		fmt.Println(" 📄", f.out)
	}

	// ---- 4. Optional feature files ---------------------------------------
	hasFeature := func(name string) bool {
		for _, f := range spec.Features {
			if strings.EqualFold(f, name) {
				return true
			}
		}
		return false
	}

	if hasFeature("JWT Authentication") {
		out := filepath.Join(spec.Name, "internal", "middleware", "jwt.go")
		if err := GenerateFile("templates/jwt_middleware.go.tpl", out, replacements); err != nil {
			return fmt.Errorf("generate jwt middleware: %w", err)
		}
		fmt.Println(" 🔐", out)
	}

	if hasFeature("Logger (Zerolog)") {
		out := filepath.Join(spec.Name, "internal", "middleware", "logger.go")
		if err := GenerateFile("templates/logger.go.tpl", out, replacements); err != nil {
			return fmt.Errorf("generate logger: %w", err)
		}
		fmt.Println(" 📋", out)
	}

	if hasFeature("Docker Support") {
		dockerFiles := []struct{ tpl, out string }{
			{"templates/Dockerfile.tpl", filepath.Join(spec.Name, "Dockerfile")},
			{"templates/docker-compose.yml.tpl", filepath.Join(spec.Name, "docker-compose.yml")},
		}
		for _, f := range dockerFiles {
			if err := GenerateFile(f.tpl, f.out, replacements); err != nil {
				return fmt.Errorf("generate %q: %w", f.out, err)
			}
			fmt.Println(" 🐳", f.out)
		}
	}

	if hasFeature("Makefile") {
		out := filepath.Join(spec.Name, "Makefile")
		if err := GenerateFile("templates/Makefile.tpl", out, replacements); err != nil {
			return fmt.Errorf("generate makefile: %w", err)
		}
		fmt.Println(" 🛠️", out)
	}

	if hasFeature("README.md") {
		out := filepath.Join(spec.Name, "README.md")
		if err := GenerateFile("templates/README.md.tpl", out, replacements); err != nil {
			return fmt.Errorf("generate readme: %w", err)
		}
		fmt.Println(" 📖", out)
	}

	// ---- 5. go mod init --------------------------------------------------
	if err := InitGoModule(spec.Name, spec.Module); err != nil {
		return fmt.Errorf("go mod init: %w", err)
	}
	fmt.Println(" 📦 go mod init done")

	// ---- 6. git init (optional) -----------------------------------------
	if spec.InitGit {
		gitCmd := exec.Command("git", "init")
		gitCmd.Dir = spec.Name
		if err := gitCmd.Run(); err != nil {
			return fmt.Errorf("git init: %w", err)
		}
		fmt.Println(" 🌿 Git repository initialised")
	}

	return nil
}
