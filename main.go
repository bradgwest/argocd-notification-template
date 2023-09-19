package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"text/template"
)

/*
todo:
* support yaml
* support sprig
* rename app
*/

func main() {
	tpl := flag.String("template", "Hello {{.app.metadata.name}}", "A template to render")
	app := flag.String("app", `{"app": {"metadata": {"name": "something"}}}`, "JSON Application spec")

	var jsn map[string]interface{}
	json.Unmarshal([]byte(*app), &jsn)

	// tpl := template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseGlob("*.html"))

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(*tpl)
	check(err)

	err = t.Execute(os.Stdout, jsn)
	check(err)
}
