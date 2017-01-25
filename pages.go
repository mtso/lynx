package lynx

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Pages is a type that represents a slice collection of Page structs.
// It is the receiver for functions that manipulate Page order.
// The Pages in an Index struct can be iterated through an HTML template.
type Pages []Page

func (p Pages) reverse() Pages {
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
	return p
}

// chronological sorts pages by BirthTime.
func (p Pages) chronological() Pages {
	for i := 0; i < len(p); i++ {
		for j := i; j < len(p); j++ {
			if p[j].isCreatedBefore(p[i]) {
				p[j], p[i] = p[i], p[j]
			}
		}
	}
	return p
}

func (p Pages) reverseChronological() Pages {
	for i := 0; i < len(p); i++ {
		for j := i; j < len(p); j++ {
			if p[j].isCreatedAfter(p[i]) {
				p[j], p[i] = p[i], p[j]
			}
		}
	}
	return p
}

func (p Pages) loadTemplateRaw(src string) error {
	t, err := template.New("post").Parse(src)
	if err != nil {
		return nil
	}

	for i := range p {
		p[i].template = t
	}

	return nil
}

func (p Pages) loadTemplate(filepath string) error {
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return err
	}

	// Execute on Page value by Index
	for i := range p {
		p[i].template = t
	}

	return nil
}

// USE THIS: https://golang.org/pkg/html/template/#hdr-Typed_Strings
func (p Pages) executeTemplate() {
	for i := range p {
		t := p[i].template
		if err := t.Execute(&p[i], p[i]); err != nil {
			log.Println(err)
		}
	}
}

func (pages Pages) exportTo(dirname string) (err error) {

	for _, p := range pages {
		// Skip Pages that have not executed their template
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

func (p Pages) relinkNext() {
	// Point each to next Page in slice
	for i := 0; i < len(p)-1; i++ {
		p[i].Next = &p[i+1]
	}
	// Last page points to nil.
	p[len(p)-1].Next = nil

	// or should the last page point to the first?
	// p[len(p)-1].Next = &p[0]
}
