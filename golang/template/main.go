package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func main() {
	funcMap := template.FuncMap{
		"func1": func1,
	}
	tmpl := template.Must(template.New("sample.tmpl").Funcs(funcMap).ParseFiles("./sample.tmpl", "./sample2.tmpl"))
	var buf bytes.Buffer
	args := map[string]interface{}{
		"String": "string",
		"Int":    10,
		"Bool":   false,
		"Array":  []int{1, 2, 3, 4, 5},
	}
	if err := tmpl.Execute(&buf, args); err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}

func func1(s string) string {
	return "func1:" + s
}
