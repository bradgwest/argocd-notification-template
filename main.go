package main

import (
	"log"
	"os"

	"html/template"
)

func main() {
	// tpl := template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseGlob("*.html"))
	const tpl = `Hello {{.Name}}
`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Name string
	}{
		Name: "Dude",
	}

	err = t.Execute(os.Stdout, data)
	check(err)
}
