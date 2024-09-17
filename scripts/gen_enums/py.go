package main

import (
	"fmt"
	"os"
	"strings"
)

func genPython(e enum) error {
	code := "import enum\n\n\n"
	code += "class ${NAME}(enum.Enum):\n"
	for _, m := range e.Members {
		key := toScreamingSnakeCase(m.Key)
		code += fmt.Sprintf("    %s = \"%s\"\n", key, m.Value)
	}
	code = strings.ReplaceAll(code, "${NAME}", e.Name)

	// Write the TypeScript code to a file
	outputPath := fmt.Sprintf("./dist/py/%s.py", e.Name)
	err := os.WriteFile(outputPath, []byte(code), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}
