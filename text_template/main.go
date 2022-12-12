package main

import (
	"os"
	"text/template"
)

func main() {
	//exampleCallFunction()
	//exampleTemplateFromFile()
	//exampleUsingMoreThaOneTemplate()
	exampleUsingMoreThanOneFile()
}

type model struct {
	Name    string
	Count   int
	GetName func(model) string
}

func getName(m model) string {
	return m.Name
}

func example1() {
	model := model{"index", 10, getName}
	temp, err := template.New("index").Parse("model name is {{ .Name }} and count is {{ .Count }}\n")

	if err != nil {
		panic(err)
	}

	err = temp.Execute(os.Stdout, model)

	if err != nil {
		panic(err)
	}
}

func exampleTrim() {
	model := model{"index", 0, getName}
	tmpl, err := template.New("index").Parse("model name is {{ .Name -}} and count is {{- .Count }}\n")

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, model)

	if err != nil {
		panic(err)
	}
}

func exampleCallFunction() {
	model := model{"index2", 0, getName}
	tmpl, err := template.New("index").Parse("model name is {{ call .GetName .}}\n")

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, model)

	if err != nil {
		panic(err)
	}
}

func exampleTemplateFromFile() {
	model := model{"hello", 0, getName}
	tmpl, err := template.New("map.yaml").ParseFiles("map.yaml")

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, model)

	if err != nil {
		panic(err)
	}
}

func exampleTemplateFromFile2() {
	model := model{"hello", 0, getName}
	tmpl, err := template.ParseFiles("map.yaml")

	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "map.yaml", model)

	if err != nil {
		panic(err)
	}
}

func exampleUsingMoreThaOneTemplate() {
	model := model{"eae", 0, getName}
	tmpl := template.New("t")

	_, err := tmpl.New("t1").Parse("word {{ .Count }}\n")

	if err != nil {
		panic(err)
	}

	t1, err := tmpl.New("t2").Parse("hello {{ .Name }}\n {{template \"t1\" .}} ")

	if err != nil {
		panic(err)
	}

	err = t1.Execute(os.Stdout, model)

	if err != nil {
		panic(err)
	}
}

func exampleUsingMoreThanOneFile() {
	model := model{"test", 2, getName}
	tmpl, err := template.ParseFiles("map.yaml", "secret.yaml")

	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "map.yaml", model)

	if err != nil {
		panic(err)
	}
}
