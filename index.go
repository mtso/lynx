package lynx

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

// index contains the data properties used in generating index.html
type index struct {
	Title string

	Pages Pages

	// Description is HTML encoded.
	Description template.HTML

	// Template inaccessible properties
	// used for export.
	template *template.Template

	html []byte
}

// newindex creates a new index object
// with specified title
func newindex(t string, d string, pgs Pages) *index {
	return &index{
		Title:       t,
		Description: template.HTML(d),
		Pages:       pgs,
	}
}

// loadTemplate loads an index template by parsing filepath.
func (i *index) loadTemplate(filepath string) error {
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return err
	}
	i.template = t
	return nil
}

// Write implements Writer interface.
func (i *index) Write(in []byte) (n int, err error) {
	i.html = append(i.html, in...)
	return len(in), nil
}

// Read implements the Reader interface.
func (i *index) Read(out []byte) (n int, err error) {
	out = append(out, i.html...)
	return len(i.html), nil
}

func (i *index) executeTemplate() error {
	return i.template.Execute(i, i)
}

func (i *index) writeTo(dirname string) error {
	filepath := filepath.Join(dirname, "index.html")
	return ioutil.WriteFile(filepath, i.html, os.ModePerm)
}
