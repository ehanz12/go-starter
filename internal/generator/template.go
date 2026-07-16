package generator

import (
	"os"
	"strings"
)

func GenerateFile(templatePath string, outputPath string, replacements map[string]string) error {

	content, err := os.ReadFile(templatePath)

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