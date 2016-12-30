package lynx

import (
	// "bytes"
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
			if p[j].isModifiedBefore(p[i]) {
				p[j], p[i] = p[i], p[j]
			}
		}
	}
	return p
}

func (p Pages) ReverseChronological() Pages {
	for i := 0; i < len(p); i++ {
		for j := i; j < len(p); j++ {
			if p[j].isModifiedAfter(p[i]) {
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
		base := filepath.Base(p.RelativeLink)
		filepath := filepath.Join(dirname, base)

		err = ioutil.WriteFile(filepath, p.html, os.ModePerm)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return
}