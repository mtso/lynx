package lynx

import (
	"log"
	"io/ioutil"
)

type Page struct {

	// Title of the page.
	Title string

	// Pointer to the next page.
	Next *Page

	// Relative link
	Link string

	// Page content.
	Content string
}

func NewPage(t string, n Page, l string, c string) *Page {
	return &Page {
		Title: t,
		Next: n,
		Link: l,
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
		if isMarkdownFilename()
		buf, err := ioutil.ReadFile(file.Name())
		if err != nil {
			log.Println(err, file.Name())
			continue
		}

		newpage = NewPage(file.Name(), prev, "./" + file.Name(), )


		pages = append(pages, )
	}
}