package lynx

import (
	"log"
	"path/filepath"
	"io/ioutil"
	"os"
	"time"
	"html/template"
	"strings"
	md "github.com/russross/blackfriday"
	"github.com/shibukawa/extstat"
)

const (
	contentTag = `{{define "post_content"}}#{{end}}`
)

type Page struct {

	// Title of the page.
	Title string

	// Pointer to the next page.
	Next *Page

	// Relative link
	RelativeLink string

	// Datetime of last modification
	ModTime time.Time

	// Datetime of the file's creation
	BirthTime time.Time

	// Page content.
	Content string
	ContentTemplate string

	html []byte

	template *template.Template
}

func NewPage(t string, n *Page, c string, modTime time.Time, rel string, ct string, birthTime time.Time) *Page {
	return &Page{
		Title: t,
		Next:  n,
		ModTime: modTime,
		Content: c,
		html:    make([]byte, 0),
		RelativeLink: rel,
		ContentTemplate: ct,
		BirthTime: birthTime,
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

// Used for sorting by time
func (p *Page) isModifiedBefore(right Page) bool {
	return p.ModTime.Before(right.ModTime)
}

func (p *Page) isModifiedAfter(right Page) bool {
	return p.ModTime.After(right.ModTime)
}

func (p *Page) isCreatedBefore(right Page) bool {
	return p.BirthTime.Before(right.BirthTime)
}

func (p *Page) isCreatedAfter(right Page) bool {
	return p.BirthTime.After(right.BirthTime)
}

func LoadPagesIn(dirname string) (Pages, error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	pages := make(Pages, 0)
	var prev *Page = nil
	for _, file := range files {
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

		// Parse front matter
		hasFrontMatter := true
		frontmatter, err := parseFrontMatterIn(buf)
		if err != nil {
			hasFrontMatter = false
		}

		// Init page properties
		title := titleFromFilename(file.Name())
		rel_link := filepath.Join(".", stripExt(file.Name()) + ".html")

		// Parse markdown
		html := md.MarkdownCommon(stripFrontMatterFrom(buf))
		content := string(html[:len(html)])
		
		// Define a string containing the html representation
		// of parsed markdown
		rawContentTemplate := strings.Replace(contentTag, "#", content, 1)

		// Get birth time with extstat
		morestats := extstat.New(stats)
		birthtime := morestats.BirthTime

		// Override file birthtime with custom time
		if hasFrontMatter {
			if customtime, ok := frontmatter["date"]; ok {
				birthtime = customtime.(time.Time)
			} 
		}

		newpage := NewPage(
			title, 
			prev, 
			content, 
			stats.ModTime(), 
			rel_link,
			rawContentTemplate,
			birthtime,
		)
		prev = newpage // Assign previous pointer to current page

		pages = append(pages, *newpage)
	}

	return pages, nil
}

func titleFromFilename(filename string) string {
	t := stripExt(filename)
	return strings.Replace(t, "-", " ", -1)
}
