package cmd

import (
	"github.com/ehanz12/go-starter/internal/generator"
	"github.com/spf13/cobra"
)

var (
	moduleFlag string
	dbFlag     string
)

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Go project (interactive wizard)",
	Long: `Scaffold a new production-ready Go REST API project.

Without flags, an interactive wizard is launched so you can choose
the database driver and optional features (JWT, Docker, Logger, etc.).

Examples:
  go-starter new my-api
  go-starter new my-api --module example.com/my-api
  go-starter new my-api --db postgres --git
  go-starter new my-api --db sqlite -m github.com/me/my-api`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInteractive(args[0], moduleFlag, generator.NormaliseDB(dbFlag))
	},
}

func init() {
	newCmd.Flags().StringVarP(&moduleFlag, "module", "m", "", "Module name (default: github.com/yourusername/<project-name>)")
	newCmd.Flags().StringVarP(&dbFlag, "db", "d", "", "Database driver: mysql | postgres | sqlite (skips DB question)")
	newCmd.Flags().BoolVar(&initGit, "git", false, "Initialize a Git repository for the generated project")
	rootCmd.AddCommand(newCmd)
}
