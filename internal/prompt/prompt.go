package prompt

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// ProjectOptions holds all the user-selected options for project generation
type ProjectOptions struct {
	ProjectName string
	ModuleName  string
	Database    string // mysql, postgres, sqlite
	Features    []string
	InitGit     bool
	Author      string
}

var databaseChoices = []string{
	"MySQL",
	"PostgreSQL",
	"SQLite",
}

var featureChoices = []string{
	"JWT Authentication",
	"Docker Support",
	"Logger (Zerolog)",
	"Makefile",
	"README.md",
}

// Run shows an interactive TUI prompt and returns the chosen options
func Run(projectName, moduleName string, autoGit bool) (*ProjectOptions, error) {
	opts := &ProjectOptions{
		ProjectName: projectName,
		ModuleName:  moduleName,
		InitGit:     autoGit,
	}

	// If project name not supplied via arg, ask for it
	if opts.ProjectName == "" {
		if err := survey.AskOne(&survey.Input{
			Message: "Project name:",
			Default: "my-api",
		}, &opts.ProjectName, survey.WithValidator(survey.Required)); err != nil {
			return nil, err
		}
	}

	// Default module name
	defaultModule := fmt.Sprintf("github.com/yourusername/%s", opts.ProjectName)
	if opts.ModuleName == "" {
		if err := survey.AskOne(&survey.Input{
			Message: "Module name:",
			Default: defaultModule,
		}, &opts.ModuleName); err != nil {
			return nil, err
		}
		if strings.TrimSpace(opts.ModuleName) == "" {
			opts.ModuleName = defaultModule
		}
	}

	// Database selection
	if err := survey.AskOne(&survey.Select{
		Message: "Select database driver:",
		Options: databaseChoices,
		Default: "MySQL",
	}, &opts.Database); err != nil {
		return nil, err
	}

	// Feature selection
	if err := survey.AskOne(&survey.MultiSelect{
		Message: "Select features to include:",
		Options: featureChoices,
		Default: []string{"JWT Authentication", "Logger (Zerolog)", "README.md"},
	}, &opts.Features); err != nil {
		return nil, err
	}

	// Git init (only ask if not already set via flag)
	if !opts.InitGit {
		if err := survey.AskOne(&survey.Confirm{
			Message: "Initialize Git repository?",
			Default: true,
		}, &opts.InitGit); err != nil {
			return nil, err
		}
	}

	return opts, nil
}

// HasFeature checks if a feature string is selected
func HasFeature(features []string, name string) bool {
	for _, f := range features {
		if strings.EqualFold(f, name) {
			return true
		}
	}
	return false
}
