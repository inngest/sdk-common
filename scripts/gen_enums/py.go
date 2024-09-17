package main

import (
	"fmt"
	"os"
	"strings"
)

func genPython(config Config) error {
	code := "import enum\n\n\n"
	code += "class ${NAME}(enum.Enum):\n"
	for _, value := range config.Values {
		if config.RemovePrefix != "" {
			value = strings.TrimPrefix(value, config.RemovePrefix)
		}
		key := toScreamingSnakeCase(value)

		code += fmt.Sprintf("    %s = \"%s\"\n", key, value)
	}
	code = strings.ReplaceAll(code, "${NAME}", config.Name)

	// Write the TypeScript code to a file
	outputPath := fmt.Sprintf("./dist/py/%s.py", config.Name)
	err := os.WriteFile(outputPath, []byte(code), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}
