package main

// `go run template.go`

// Import libraries
import (
	"fmt"
	HtmlTemplate "html/template" // generates HTML output safe against code injection
	"os"
	"text/template"
)

func simple1() {
	type Inventory struct {
		Material string
		Count    uint
	}

	sweaters := Inventory{"wool", 17}

	// Create a New template and populate it with Parse contents. Must just checks for parse error and panics
	tmpl := template.Must(template.New("test").Parse("{{.Count}} items are made of {{.Material}}\n"))
	if err := tmpl.Execute(os.Stdout, sweaters); err != nil {
		panic(err)
	}
}

func file1() {
	type entry struct {
		Name string
		Done bool
	}

	type ToDo struct {
		User string
		List []entry
	}

	// Parse data -- omitted for brevity
	todos := ToDo{"justin", []entry{entry{"Practice go", false}, entry{"Eat dirt", true}}}

	// Files are provided as a slice of strings.
	paths := []string{
		"template1.html",
	}

	fmt.Println(todos)

	t := HtmlTemplate.Must(HtmlTemplate.New("template1.html").ParseFiles(paths...))
	err := t.Execute(os.Stdout, todos)
	if err != nil {
		panic(fmt.Errorf("Template Execute error: %s", err))
	}
}

func main() {
	simple1()
	file1()

	// Read a template from file an execute it
}
