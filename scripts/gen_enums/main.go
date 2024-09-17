package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Name         string
	RemovePrefix string   `json:"remove_prefix"`
	Values       []string `json:"values"`
}

func main() {
	enumsDir := "./src/enums"
	files, err := os.ReadDir(enumsDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		data, err := os.ReadFile(fmt.Sprintf("%s/%s", enumsDir, file.Name()))
		if err != nil {
			fmt.Println("error reading file:", err)
			return
		}

		var config Config
		err = json.Unmarshal(data, &config)
		if err != nil {
			fmt.Println("error parsing JSON:", err)
			return
		}
		config.Name = strings.Split(file.Name(), ".")[0]

		err = genGo(config)
		if err != nil {
			fmt.Println("error generating Go code:", err)
			return
		}

		err = genPython(config)
		if err != nil {
			fmt.Println("error generating Python code:", err)
			return
		}

		err = genTypeScript(config)
		if err != nil {
			fmt.Println("error generating TypeScript code:", err)
			return
		}
	}
}
