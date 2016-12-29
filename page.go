package lynx

import (
	"log"
	"path/filepath"
	"io/ioutil"
	"os"
	"html/template"
)

type Page struct {

	// Title of the page.
	Title string

	// Pointer to the next page.
	Next *Page

	// Relative link
	// Link string

	// Page content.
	Content string

	html []byte
}

type PageGroup []Page

func NewPage(t string, n *Page, c string) *Page {
	return &Page{
		Title: t,
		Next:  n,
		// Link: l,
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

func LoadPagesIn(dirname string) (PageGroup, error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	pages := make(PageGroup, 0)
	var prev *Page = nil
	for _, file := range files {
		// if !(filepath.Ext(file.Name()) == ".md")
		if !isMarkdownExtension(file.Name()) {
			continue
		}

		buf, err := ioutil.ReadFile(filepath.Join(dirname, file.Name()))
		if err != nil {
			log.Println(err)
			continue
		}

		content := string(buf[:len(buf)])
		newpage := NewPage(file.Name(), prev, content)
		prev = newpage

		pages = append(pages, *newpage)
	}

	return pages, nil
}

// Makes `public` directory to contain the static
// files for hosting
func genPublicDir() error {
	err := os.MkdirAll("public", os.ModePerm)
	if err == os.ErrInvalid || err == os.ErrPermission {
		return err
	}
	return nil
}

func (pages PageGroup) ExportTo(dirname string) error {

	// Init a new template by parsing post-demo file
	t, err := template.ParseFiles("template/post-demo.html")
	if err != nil {
		return err
	}

	if err = genPublicDir(); err != nil {
		return err
	}

	for _, p := range pages {
		// Execute post-demo template with Page object
		err = t.Execute(&p, p)
		if err != nil {
			log.Println(err)
			continue
		}

		filepath := filepath.Join(dirname, stripExt(p.Title)) + ".html"
		err = ioutil.WriteFile(filepath, p.html, os.ModePerm)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return err
}
