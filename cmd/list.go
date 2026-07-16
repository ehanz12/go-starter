package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type featureInfo struct {
	name        string
	description string
	flag        string
}

var features = []featureInfo{
	{
		name:        "JWT Authentication",
		description: "JWT middleware with role-based access control (golang-jwt/jwt)",
		flag:        "Selected in interactive wizard",
	},
	{
		name:        "Docker Support",
		description: "Multi-stage Dockerfile (scratch image) + docker-compose with DB service",
		flag:        "Selected in interactive wizard",
	},
	{
		name:        "Logger (Zerolog)",
		description: "Structured JSON/pretty logger with Fiber HTTP request logging",
		flag:        "Selected in interactive wizard",
	},
	{
		name:        "Makefile",
		description: "Targets: run, build, test, lint, docker-up/down, migrate, help",
		flag:        "Selected in interactive wizard",
	},
	{
		name:        "README.md",
		description: "Auto-generated README with tech stack, project structure, and commands",
		flag:        "Selected in interactive wizard",
	},
}

var databases = []featureInfo{
	{name: "MySQL", description: "gorm.io/driver/mysql  — port 3306", flag: "--db mysql"},
	{name: "PostgreSQL", description: "gorm.io/driver/postgres — port 5432", flag: "--db postgres"},
	{name: "SQLite", description: "gorm.io/driver/sqlite  — file-based, no server needed", flag: "--db sqlite"},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available databases and optional features",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Println("╔══════════════════════════════════════════════════════════════╗")
		fmt.Println("║            🗄️   Supported Databases                          ║")
		fmt.Println("╚══════════════════════════════════════════════════════════════╝")
		for _, db := range databases {
			fmt.Printf("  %-12s  %s\n", db.name, db.description)
			fmt.Printf("              Flag: %s\n\n", db.flag)
		}

		fmt.Println("╔══════════════════════════════════════════════════════════════╗")
		fmt.Println("║            ✨   Optional Features                            ║")
		fmt.Println("╚══════════════════════════════════════════════════════════════╝")
		for i, f := range features {
			fmt.Printf("  %d. %-25s %s\n", i+1, f.name, f.description)
		}

		fmt.Println()
		fmt.Println("  All features are selected via the interactive wizard:")
		fmt.Println("    go-starter new my-api")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
