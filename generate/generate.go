package main

import (
	"regexp"
	"text/template"
	"os"
	"strings"
	"fmt"
)

var header = `package pval

`

var typeNames = `string
int
uint8
uint16
uint32
uint64
uintptr
byte
int8
int16
int32
int64
rune
float32
float64
complex64
complex128`

var templateStr = `
func {{.TypeNameTitle}}(val {{.TypeName}}) *{{.TypeName}} {
	return &val
}

func {{.TypeNameTitle}}s(vals ...{{.TypeName}}) []*{{.TypeName}} {
	var result []*{{.TypeName}}
	for _, val := range vals {
		result = append(result, &val)
	}
	return result
}


`

type templateData struct {
	TypeName string
	TypeNameTitle string
}

func main() {
	var re = regexp.MustCompile(`(?m)^\s*(?P<typeName>\S+).*$`)
	templ, err := template.New("funcs").Parse(templateStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(header)
	re.ReplaceAllStringFunc(typeNames, func(typeName string) string {
		err := templ.Execute(os.Stdout, templateData{typeName, strings.Title(typeName)})
		if err != nil {
			panic(err)
		}
		return ""
	})
}
