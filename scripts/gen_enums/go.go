package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func genGo(e enum) error {
	code := `
package sdkcommon

type ${NAME} string

const (
`
	for _, m := range e.Members {
		key := "${NAME}" + toPascalCase(m.Key)
		code += fmt.Sprintf("\t%s ${NAME} = \"%s\"\n", key, m.Value)
	}
	code += ")\n"
	code = strings.ReplaceAll(code, "${NAME}", e.Name)

	outputPath := fmt.Sprintf("./dist/go/%s.go", e.Name)
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
