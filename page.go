package lynx

import (
	"log"
	// "errors"
	"path/filepath"
	// "bufio"
	// "io"
	"io/ioutil"
	// "os"
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
	return &Page {
		Title: t,
		Next: n,
		// Link: l,
		Content: c,
		html: make([]byte, 0),
	}
}

// Implement Writer interface
func (p *Page) Write(in []byte) (n int, err error) {
	for _, b := range in {
		p.html = append(p.html, b)
	}
	return len(in), nil
}

func LoadPagesIn(dirname string) PageGroup {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
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

		// content := string(buf)
		content := string(buf[:len(buf)])
		newpage := NewPage(file.Name(), prev, content)
		prev = newpage

		pages = append(pages, *newpage)
	}

	return pages
}

func (pages PageGroup) ExportTo(dirname string) error {

	t, err := template.ParseFiles("template/post-demo.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range pages {
		err = t.Execute(&p, p)
		if err != nil {
			log.Println(err)
		}

	}

	return err
}
