package main

import (
	"os"
	"text/template"
)

type Payload struct {
	UserID   int
	Username string
	Email    string
	Active   bool
}

func main() {
	const jsonTemplate = `
{
  "user_id": {{.UserID}},
  "username": "{{.Username}}",
  "email": "{{.Email}}",
  "active": {{.Active}}
}
`

	data := Payload{
		UserID:   101,
		Username: "radoslaw",
		Email:    "rado@example.com",
		Active:   true,
	}

	t := template.Must(template.New("json").Parse(jsonTemplate))
	t.Execute(os.Stdout, data)
}
