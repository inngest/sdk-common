package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func genGo(config Config) error {
	code := `
package sdkcommon

type ${NAME} string

const (
`
	for _, value := range config.Values {
		key := value
		if config.RemovePrefix != "" {
			key = strings.TrimPrefix(key, config.RemovePrefix)
		}
		key = config.Name + toPascalCase(key)
		code += fmt.Sprintf("\t%s %s = \"%s\"\n", key, config.Name, value)
	}
	code += ")\n"
	code = strings.ReplaceAll(code, "${NAME}", config.Name)

	outputPath := fmt.Sprintf("./dist/go/%s.go", config.Name)
	err := os.WriteFile(outputPath, []byte(code), 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	// Run go fmt on the generated file
	cmd := exec.Command("go", "fmt", outputPath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error running go fmt: %w", err)
	}

	return nil
}
