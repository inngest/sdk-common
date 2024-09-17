package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/emicklei/proto"
)

type member struct {
	Key   string
	Value string
}

type enum struct {
	Members []member
	Name    string
}

func main() {
	enumsDir := "./src/enums"
	files, err := os.ReadDir(enumsDir)
	if err != nil {
		fmt.Println("error reading directory:", err)
		return
	}

	for _, file := range files {
		data, err := os.ReadFile(fmt.Sprintf("%s/%s", enumsDir, file.Name()))
		if err != nil {
			fmt.Println("error reading file:", err)
			return
		}

		reader := strings.NewReader(string(data))
		parser := proto.NewParser(reader)
		definition, err := parser.Parse()
		if err != nil {
			fmt.Println("error parsing proto file:", err)
			return
		}

		var protoEnum *proto.Enum
		proto.Walk(definition,
			proto.WithEnum(func(m *proto.Enum) {
				if protoEnum != nil {
					fmt.Printf("multiple enums found in proto file %s\n", file.Name())
					return
				}

				protoEnum = m
			}),
		)
		if protoEnum == nil {
			fmt.Println("no Mode message found in proto file")
			return
		}

		e := enum{Name: protoEnum.Name, Members: make([]member, 0)}
		for _, element := range protoEnum.Elements {
			if field, ok := element.(*proto.EnumField); ok {
				key := field.Name
				value := field.ValueOption.Constant.Source
				e.Members = append(e.Members, member{Key: key, Value: value})
			}
		}

		err = genGo(e)
		if err != nil {
			fmt.Println("error generating Go code:", err)
			return
		}

		err = genPython(e)
		if err != nil {
			fmt.Println("error generating Python code:", err)
			return
		}

		err = genTypeScript(e)
		if err != nil {
			fmt.Println("error generating TypeScript code:", err)
			return
		}
	}
}
