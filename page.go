package lynx

import (
	"log"
	"path/filepath"
	"io/ioutil"
	"os"
	"time"
	"html/template"
)

type Page struct {

	// Title of the page.
	Title string

	// Pointer to the next page.
	Next *Page

	// Relative link
	// Link string

	ModTime time.Time

	// Page content.
	Content string

	html []byte

	template *template.Template
}

type Pages []Page

func NewPage(t string, n *Page, c string, modTime time.Time) *Page {
	return &Page{
		Title: t,
		Next:  n,
		ModTime: modTime,
		Content: c,
		html:    make([]byte, 0),
	}
}

// Implement Writer interface
func (p *Page) Write(in []byte) (n int, err error) {
	p.html = append(p.html, in...)
	return len(in), nil
}

// Implement Reader interface
func (p *Page) Read(out []byte) (n int, err error) {
	out = append(out, p.html...)
	return len(p.html), nil
}

func (p Page) String() string {
	return p.Title
}

func LoadPagesIn(dirname string) (Pages, error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	pages := make(Pages, 0)
	var prev *Page = nil
	for _, file := range files {
		// if !(filepath.Ext(file.Name()) == ".md")
		if !isMarkdownExtension(file.Name()) {
			continue
		}

		path := filepath.Join(dirname, file.Name())

		// Get file info
		stats, err := os.Stat(path) //file.Stat()
		if err != nil {
			log.Println(err)
			continue
		}

		// Read file content
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			log.Println(err)
			continue
		}

		content := string(buf[:len(buf)])
		
		title := stripExt(file.Name())
		newpage := NewPage(title, prev, content, stats.ModTime())
		prev = newpage

		pages = append(pages, *newpage)
	}

	return pages, nil
}

func (pages Pages) loadTemplate(filepath string) error {
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return err
	}

	// Execute on page value by index
	for i := range pages {
		pages[i].template = t
	}

	return nil
}

func (pages Pages) executeTemplate() {
	for i := range pages {
		t := pages[i].template
		if err := t.Execute(&pages[i], pages[i]); err != nil {
			log.Println(err)
		}
	}
}

func (pages Pages) ExportTo(dirname string) (err error) {

	for _, p := range pages {
		// Skip pages that have not executed their template
		if len(p.html) == 0 {
			continue
		}

		filepath := filepath.Join(dirname, stripExt(p.Title)) + ".html"
		err = ioutil.WriteFile(filepath, p.html, os.ModePerm)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return
}
