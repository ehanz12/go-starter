package generator

import (
	"embed"
	"os"
	"strings"
)

//go:embed templates
var templatesFS embed.FS

func GenerateFile(templatePath string, outputPath string, replacements map[string]string) error {

	content, err := templatesFS.ReadFile(templatePath)

	if err != nil {
		return err
	}

	result := string(content)

	for key, value := range replacements {

		result = strings.ReplaceAll(
			result,
			"{{"+key+"}}",
			value,
		)

	}

	return os.WriteFile(outputPath, []byte(result), 0644)
}