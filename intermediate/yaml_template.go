package main

import (
	"os"
	"text/template"
)

type PodData struct {
	Name      string
	Namespace string
	Image     string
	Port      int
}

func main() {
	const podTemplate = `
apiVersion: v1
kind: Pod
metadata:
  name: {{.Name}}
  namespace: {{.Namespace}}
spec:
  containers:
  - name: {{.Name}}-container
    image: {{.Image}}
    ports:
    - containerPort: {{.Port}}
`

	data := PodData{
		Name:      "my-app",
		Namespace: "default",
		Image:     "nginx:latest",
		Port:      80,
	}

	t := template.Must(template.New("pod").Parse(podTemplate))
	t.Execute(os.Stdout, data)
}
