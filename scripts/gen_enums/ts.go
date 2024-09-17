package main

import (
	"fmt"
	"os"
	"strings"
)

func genTypeScript(config Config) error {
	code := "export const ${NAME} = {\n"
	for _, value := range config.Values {
		if config.RemovePrefix != "" {
			value = strings.TrimPrefix(value, config.RemovePrefix)
		}
		key := toPascalCase(value)
		code += fmt.Sprintf("  %s: \"%s\",\n", key, value)
	}
	code += "} as const;\n\n"
	code += "export type ${NAME} = (typeof ${NAME})[keyof typeof ${NAME}];\n"
	code = strings.ReplaceAll(code, "${NAME}", config.Name)

	// Write the TypeScript code to a file
	outputPath := fmt.Sprintf("./dist/ts/%s.ts", config.Name)
	err := os.WriteFile(outputPath, []byte(code), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}
