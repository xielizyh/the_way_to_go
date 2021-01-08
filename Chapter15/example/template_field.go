package main

import (
	"fmt"
	"os"
	"text/template"
)

type person struct {
	Name                string
	nonExportedAgeField string
}

func templateField() {
	t := template.New("hello")
	t, _ = t.Parse("hello {{.Name}}!")
	p := person{Name: "Mary", nonExportedAgeField: "32"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
