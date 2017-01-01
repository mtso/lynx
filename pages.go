package lynx

import (
	"html/template"
	"log"
	"path/filepath"
	"io/ioutil"
	"os"
)

type Pages []Page

func (p Pages) Reverse() Pages {
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
	return p
}

func (p Pages) Chronological() Pages {
	for i := 0; i < len(p); i++ {
		for j := i; j < len(p); j++ {
			if p[j].isCreatedBefore(p[i]) {
				p[j], p[i] = p[i], p[j]
			}
		}
	}
	return p
}

func (p Pages) ReverseChronological() Pages {
	for i := 0; i < len(p); i++ {
		for j := i; j < len(p); j++ {
			if p[j].isCreatedAfter(p[i]) {
				p[j], p[i] = p[i], p[j]
			}
		}
	}
	return p
}

func (pages Pages) loadTemplate(filepath string) error {
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return err
	}

	// Execute on page value by index
	for i := range pages {
		// Clone the base template
		// This allows us to use the clone to parse
		// this page's ContentTemplate containing markdown
		tc, err := t.Clone()
		if notNil(err) {
			continue
		}

		// Attach this page's content template
		// to its base `post` template
		t, err := tc.Parse(pages[i].ContentTemplate)
		if notNil(err) {
			continue
		}
		pages[i].template = t
	}

	return nil
}

// USE THIS: https://golang.org/pkg/html/template/#hdr-Typed_Strings
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

		// Build filepath from base of relative link
		rel_link := filepath.Base(p.RelativeLink)
		dir_path := filepath.Join(dirname, rel_link)

		rel_location := filepath.Join(rel_link, "index.html")
		full_filepath := filepath.Join(dirname, rel_location)

		// Make directories for permalink path
		mkdirIfNone(dir_path)

		err = ioutil.WriteFile(full_filepath, p.html, os.ModePerm)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return
}

func (p Pages) RelinkNext() {
	// Point each to next Page in slice
	for i := 0; i < len(p) - 1; i++ {
		p[i].Next = &p[i+1]
	}
	// Last page points to none
	p[len(p)-1].Next = nil

	// or first?
	// p[len(p)-1].Next = &p[0]
}