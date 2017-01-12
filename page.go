package lynx

import (
	"fmt"
	"github.com/mtso/readability"
	md "github.com/russross/blackfriday"
	"github.com/shibukawa/extstat"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	contentTag = `{{define "post_content"}}#{{end}}`
)

// Ref: "Capitalize all words in titles of publications and documents,
// except a, an, the, at, by, for, in, of, on, to, up, and, as, but, or, and nor."
// http://grammar.yourdictionary.com/capitalization/rules-for-capitalization-in-titles.html
var titleExceptions = strings.NewReplacer(
	" A ", " a ",
	" An ", " an ",
	" The ", " the ",
	" At ", " at ",
	" By ", " by ",
	" For ", " for ",
	" In ", " in ",
	" Of ", " of ",
	" On ", " on ",
	" To ", " to ",
	" Up ", " up ",
	" And ", " and ",
	" As ", " as ",
	" But ", " but ",
	" Or ", " or ",
	" Nor ", " nor ",
)

// Page represents the data and properties of a single post.
type Page struct {

	// Title of the Page.
	Title string

	// Pointer to the next Page.
	Next *Page

	// Relative link
	RelativeLink string

	// Datetime of last modification
	ModTime time.Time

	// Datetime of the file's creation
	BirthTime time.Time

	// Page content.
	Content string

	// Flesch-Kincaid reading level
	FleschKinkaid string

	html []byte

	template *template.Template

	// temporary template for Content body
	contentTemplate string
}

func newPage(t string, n *Page, c string, modTime time.Time, rel string, ct string, birthTime time.Time, fk string) *Page {
	return &Page{
		Title:           t,
		Next:            n,
		ModTime:         modTime,
		Content:         c,
		html:            make([]byte, 0),
		RelativeLink:    rel,
		contentTemplate: ct,
		BirthTime:       birthTime,
		FleschKinkaid:   fk,
	}
}

// Write implements the Writer interface for use with the HTML template execution.
func (p *Page) Write(in []byte) (n int, err error) {
	p.html = append(p.html, in...)
	return len(in), nil
}

// Read implements the Reader interface for use with the HTML template execution.
func (p *Page) Read(out []byte) (n int, err error) {
	out = append(out, p.html...)
	return len(p.html), nil
}

// String implements the Stringer interface by returning the Title of the Page.
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

func loadPagesIn(dirname string) (Pages, error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	Pages := make(Pages, 0)
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

		// Init Page properties
		title := titleFromFilename(file.Name())
		lowercase := strings.ToLower(title)
		dashedTitle := strings.Replace(lowercase, " ", "-", -1)

		// Relative link will be the clean directory name
		// that eventually contains the index.html file
		// without .html extension
		root := "."
		rel_link := filepath.Join(root, dashedTitle)

		// Parse markdown
		article := stripFrontMatterFrom(buf)
		html := md.MarkdownCommon(article)
		content := string(html[:len(html)])

		// Calculate rough Flesch Kinkaid
		// (atm reads in html tags too :/)
		artstr := string(article[:len(article)])
		fleschKinkaid := read.Fk(artstr)
		fkstr := fmt.Sprintf("%.1f", fleschKinkaid)

		// Define a string containing the html representation
		// of parsed markdown
		rawcontentTemplate := strings.Replace(contentTag, "#", content, 1)

		// Get birth time with extstat
		morestats := extstat.New(stats)
		birthtime := morestats.BirthTime

		// Override file birthtime with custom time
		if hasFrontMatter {
			if customtime, ok := frontmatter["date"]; ok {
				birthtime = customtime.(time.Time)
			}
		}

		newPage := newPage(
			title,
			prev,
			content,
			stats.ModTime(),
			rel_link,
			rawcontentTemplate,
			birthtime,
			fkstr,
		)
		prev = newPage // Assign previous pointer to current Page

		Pages = append(Pages, *newPage)
	}

	return Pages, nil
}

func titleFromFilename(filename string) (t string) {
	t = stripExt(filename)
	t = strings.Replace(t, "-", " ", -1)

	t = strings.Title(t)
	t = titleExceptions.Replace(t)
	return
}
