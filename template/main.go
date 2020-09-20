package main

import (
	"os"
	"text/template"
	"time"
)

type Todo struct {
	Name        string
	Description string
}

func base() {
	td := Todo{"Test templates", "Let's test a template to see the magic."}

	t, err := template.New("todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}
}

var ttt = `<ul>
    {{range .}}
    	<li>{{.}}</li>
    {{end}}
</ul>
`

func loops() {

	todos := []string{"tang", "abc", "cde"}
	tem, _ := template.New("tttt").Parse(ttt)

	excute(tem, todos)

}

var ttt2 = `<div>
	{{$1 := .Name}}{{$2 := .Age}}
  {{$1}} {{$2}}
</div>`

type A struct {
	Name string
	Age  int
}

func variables() {

	todos := A{
		"JIM",
		123,
	}

	t, _ := template.New("tttt").Parse(ttt2)

	excute(t, todos)

}

var tf = `{{if .HasPermission}}<h3>Feature A</h3>
{{else}}
    <h3>No  Feature A</h3>
{{end}}`

func (a A) HasPermission() bool {
	return false
}

func funcT() {

	todos := A{
		"JIM",
		123,
	}

	t, _ := template.New("tttt").Parse(tf)

	excute(t, todos)

}

func main() {

	time.Now()
	// loops()
	funcT()
}

func excute(t *template.Template, v interface{}) {

	err := t.Execute(os.Stdout, v)
	if err != nil {
		panic(err)
	}
}
