package lynx

import (
	"log"
	"path/filepath"
	"io/ioutil"
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
}

func NewPage(t string, n *Page, c string) *Page {
	return &Page {
		Title: t,
		Next: n,
		// Link: l,
		Content: c,
	}
}

func LoadPagesIn(dirname string) []Page {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	pages := make([]Page, 0)
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

	return pages
}