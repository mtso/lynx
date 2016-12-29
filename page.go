package lynx

import (
	"log"
	"errors"
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

type PageGroup []Page

func NewPage(t string, n *Page, c string) *Page {
	return &Page {
		Title: t,
		Next: n,
		// Link: l,
		Content: c,
	}
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

		content := string(buf[:len(buf)])
		newpage := NewPage(file.Name(), prev, content)
		prev = newpage

		pages = append(pages, *newpage)
	}

	return pages
}

func (pages PageGroup) Print() {
	for _, p := range pages {
		// fmt.Println(p.Title, p.Next, p.Content)
	}
}

func (pages PageGroup) ExportTo(dirname string) error {

}


func BuildPagesIn(dirname string, exportDir string) {
	sm := &SiteMap{}

	// Load filenames in dirname
	if err := sm.LoadFilepathsIn(dirname); err != nil {
		log.Fatal(err)
	}

	// Validate export directory
	if !strings.HasPrefix(exportDir, "./") {
		exportDir = "./" + exportDir
	}

	// Make export directory
	err := os.MkdirAll(exportDir, os.ModePerm)
	if err == os.ErrInvalid || err == os.ErrPermission {
		log.Fatal(err)
	}
	
	exportName := ""
	for _, filename := range sm.Pages {
		// load text
		buf, err := ioutil.ReadFile(dirname + "/" + filename)
		if err != nil {
			log.Fatal(err)
		}

		// write html file
		exportName = exportDir + "/" + filename + ".html"
		_ = ioutil.WriteFile(exportName, buf, os.ModePerm)
	}

}