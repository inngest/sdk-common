package main

import (
	"fmt"
	"os"
	"strings"
)

func genTypeScript(e enum) error {
	code := "export const ${NAME} = {\n"
	for _, m := range e.Members {
		key := toPascalCase(m.Key)
		code += fmt.Sprintf("  %s: \"%s\",\n", key, m.Value)
	}
	code += "} as const;\n\n"
	code += "export type ${NAME} = (typeof ${NAME})[keyof typeof ${NAME}];\n"
	code = strings.ReplaceAll(code, "${NAME}", e.Name)

	// Write the TypeScript code to a file
	outputPath := fmt.Sprintf("./dist/ts/%s.ts", e.Name)
	err := os.WriteFile(outputPath, []byte(code), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}
