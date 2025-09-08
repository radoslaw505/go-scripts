package main

import (
	"os"
	"text/template"
)

type User struct {
	Name  string
	Age   int
	Admin bool
}

func main() {
	// Define the template string
	const tmpl = `
👋 Hello, {{.Name}}!
{{if .Admin}}🔐 You have admin access.{{else}}🔒 You are a regular user.{{end}}
You are {{.Age}} years old.

{{/* Loop over a slice of users */}}
Team Members:
{{range .Team}}
- {{.Name}} ({{if .Admin}}admin{{else}}user{{end}})
{{end}}
`

	// Create the data structure
	data := struct {
		Name  string
		Age   int
		Admin bool
		Team  []User
	}{
		Name:  "Radosław",
		Age:   30,
		Admin: true,
		Team: []User{
			{"Alice", 28, false},
			{"Bob", 35, true},
			{"Charlie", 22, false},
		},
	}

	// Parse and execute the template
	t := template.Must(template.New("profile").Parse(tmpl))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
