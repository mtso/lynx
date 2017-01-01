package lynx

import (
	"html/template"
	"path/filepath"
	"io/ioutil"
	"os"
)

type Index struct {

	Title string

	Pages Pages

	Description string

	// Template in-accessible properties
	// used for export
	template *template.Template

	html []byte
}

// Creates a new index object
// with specified title
func newIndex(t string, d string, pgs Pages) *Index {
	return &Index {
		Title: t,
		Description: d,
		Pages: pgs,
	}
}

// Load template by parsing filepath
func (i *Index) loadTemplate(filepath string) error {
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return err
	}
	i.template = t
	return nil
}

// Implement Writer interface
func (i *Index) Write(in []byte) (n int, err error) {
	i.html = append(i.html, in...)
	return len(in), nil
}

// Implement Reader interface
func (i *Index) Read(out []byte) (n int, err error) {
	out = append(out, i.html...)
	return len(i.html), nil
}

func (i *Index) executeTemplate() error {
	return i.template.Execute(i, i)
}

func (i *Index) writeTo(dirname string) error {
	filepath := filepath.Join(dirname, "index.html")
	return ioutil.WriteFile(filepath, i.html, os.ModePerm)
}