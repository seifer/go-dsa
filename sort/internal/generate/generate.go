package main

import (
	"fmt"
	"os"
	"text/template"
)

var types = []struct {
	Name     string
	FuncName string
}{
	{"int", "Ints"},
	{"int8", "Int8s"},
	{"int16", "Int16s"},
	{"int32", "Int32s"},
	{"int64", "Int64s"},

	{"uint", "Uints"},
	{"uint8", "Uint8s"},
	{"uint16", "Uint16s"},
	{"uint32", "Uint32s"},
	{"uint64", "Uint64s"},

	{"float32", "Float32s"},
	{"float64", "Float64s"},

	{"string", "Strings"},
}

func main() {
	fmt.Printf("Generate %s\n", PACKAGE)

	os.Remove(fmt.Sprintf("../../%s/numeric.go", PACKAGE))
	os.Remove(fmt.Sprintf("../../%s/numeric_test.go", PACKAGE))

	write_main_file()
	write_test_file()
}

func write_main_file() {
	FILE := fmt.Sprintf("../../%s/%s.go", PACKAGE, PACKAGE)

	if err := os.Remove(FILE); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}

	file, err := os.Create(FILE)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("func").Parse(TEMPLATE)
	if err != nil {
		panic(err)
	}

	file.WriteString("// This file should be auto generated\n\n")
	file.WriteString("package " + PACKAGE + "\n")
	for i := range types {
		tmpl.Execute(file, types[i])
	}
	file.Close()
}

func write_test_file() {
	FILE := fmt.Sprintf("../../%s/%s_test.go", PACKAGE, PACKAGE)

	if err := os.Remove(FILE); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}

	file, err := os.Create(FILE)
	if err != nil {
		panic(err)
	}

	file.WriteString("// This file should be auto generated\n\n")
	file.WriteString("package " + PACKAGE + "\n")
	file.WriteString(TEMPLATE_NUMERIC_TEST_FUNC)
	file.Close()
}
