package lynx

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Index contains the data properties used in generating index.html
type Index struct {
	Title string

	Pages Pages

	// Description is HTML encoded.
	Description template.HTML

	// Template inaccessible properties
	// used for export.
	template *template.Template

	html []byte
}

// newIndex creates a new Index object
// with specified title
func newIndex(t string, d string, pgs Pages) *Index {
	return &Index{
		Title:       t,
		Description: template.HTML(d),
		Pages:       pgs,
	}
}

// loadTemplate loads an Index template by parsing filepath.
func (i *Index) loadTemplate(filepath string) error {
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return err
	}
	i.template = t
	return nil
}

// Write implements Writer interface.
func (i *Index) Write(in []byte) (n int, err error) {
	i.html = append(i.html, in...)
	return len(in), nil
}

// Read implements the Reader interface.
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
