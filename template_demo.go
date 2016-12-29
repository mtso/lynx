package lynx

import (
	"html/template"
	"io/ioutil"

	"log"
	"os"
)

func TemplateDemo() {
	b, err := ioutil.ReadFile("template/index.html")
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Title string
	}{
		Title: "My Page",
	}

	t, err := template.New("index").Parse(string(b))
	err = t.Execute(os.Stdout, data)
}
