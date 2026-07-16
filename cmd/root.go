package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ehanz12/go-starter/internal/generator"
	"github.com/ehanz12/go-starter/internal/prompt"
	"github.com/spf13/cobra"
)

var (
	initGit    bool
	interactive bool
)

var rootCmd = &cobra.Command{
	Use:   "go-starter [project-name]",
	Short: "🚀 A professional Go backend project scaffolder",
	Long: `go-starter scaffolds production-ready Go REST API projects.

It supports multiple databases (MySQL, PostgreSQL, SQLite), optional
JWT authentication, zerolog logger, Docker support, Makefiles, and more.

Examples:
  go-starter new my-api                    # interactive wizard
  go-starter new my-api --db postgres      # skip DB question
  go-starter new my-api -m example.com/my-api --git
  go-starter list                          # show available features`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		return runInteractive(args[0], "", "")
	},
}

// Execute is the entrypoint called by main.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// runInteractive launches the survey TUI and then scaffolds the project.
func runInteractive(projectName, moduleName, dbChoice string) error {
	if err := validateProjectName(projectName); err != nil {
		return err
	}

	projectPath := filepath.Join(".", projectName)
	if _, err := os.Stat(projectPath); err == nil {
		if _, err := os.Stat(filepath.Join(projectPath, "go.mod")); err == nil {
			return fmt.Errorf("directory %q already contains a go.mod file; choose another name or remove it first", projectName)
		}
	}

	// Launch interactive prompt
	opts, err := prompt.Run(projectName, moduleName, initGit)
	if err != nil {
		return fmt.Errorf("prompt cancelled: %w", err)
	}

	// Override DB if specified via flag
	if dbChoice != "" {
		opts.Database = dbChoice
	}

	spec := generator.ProjectSpec{
		Name:     opts.ProjectName,
		Module:   resolveModuleName(opts.ProjectName, opts.ModuleName),
		DB:       generator.NormaliseDB(opts.Database),
		Features: opts.Features,
		InitGit:  opts.InitGit,
	}

	printBanner(spec)

	if err := generator.CreateProject(spec); err != nil {
		return err
	}

	printSuccess(spec)
	return nil
}

func printBanner(spec generator.ProjectSpec) {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║        🚀  go-starter                ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Printf(" Project : %s\n", spec.Name)
	fmt.Printf(" Module  : %s\n", spec.Module)
	fmt.Printf(" Database: %s\n", spec.DB)
	fmt.Printf(" Features: %s\n", strings.Join(spec.Features, ", "))
	fmt.Printf(" Git Init: %v\n", spec.InitGit)
	fmt.Println()
}

func printSuccess(spec generator.ProjectSpec) {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║        ✅  Project Ready!            ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Printf("\n  cd %s\n", spec.Name)
	fmt.Println("  cp .env.example .env   # configure your env")
	fmt.Println("  go mod tidy            # download dependencies")
	fmt.Println("  go run ./cmd/api/main.go")
	fmt.Println()
}

func resolveModuleName(projectName, moduleName string) string {
	if strings.TrimSpace(moduleName) != "" {
		return moduleName
	}
	return fmt.Sprintf("github.com/yourusername/%s", projectName)
}

func validateProjectName(projectName string) error {
	projectName = strings.TrimSpace(projectName)
	if projectName == "" {
		return fmt.Errorf("project name cannot be empty")
	}
	matched, err := regexp.MatchString(`^[a-zA-Z0-9._-]+$`, projectName)
	if err != nil || !matched {
		return fmt.Errorf("project name must only contain letters, numbers, dots, underscores, or hyphens")
	}
	return nil
}

func init() {
	rootCmd.Flags().BoolVar(&initGit, "git", false, "Initialize a Git repository for the generated project")
}
