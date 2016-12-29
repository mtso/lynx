package lynx

import (
	"html/template"
)

type Index struct {

	Title string

	Pages Pages

	template *template.Template

	html []byte
}

// Creates a new index object
// with specified title
func NewIndex(t string, pgs Pages) *Index {
	return &Index {
		Title: t,
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
	i.html = append(p.html, in...)
	return len(in), nil
}

// Implement Reader interface
func (i *Index) Read(out []byte) (n int, err error) {
	out = append(out, i.html...)
	return len(i.html), nil
}

func (i *Index) executeTemplate() error {
	return i.template.Execute(&i, i)
}