package lynx

import (
	"os"
	"path/filepath"
)

func Generate() error {

	// Make public directory to store files
	if err := mkdirIfNone("public"); err != nil {
		return err
	}

	// Load content pages
	pages, err := LoadPagesIn("content")
	if err != nil {
		return err
	}

	// Save content pages
	err = pages.ExportTo("public")
	if err != nil {
		return err
	}

	// Generate index page
	index := NewIndex("Blog", pages)
	indexTemplatePath := filepath.Join("template", "index-demo.html")
	if err := index.loadTemplate(indexTemplatePath); err != nil {
		return err
	}

	if err := index.executeTemplate(); err != nil {
		return err
	}

	if err := index.writeTo("public"); err != nil {
		return err
	}

	return nil
}

// Makes a directory if none exists
func mkdirIfNone(dirname string) error {
	err := os.MkdirAll(dirname, os.ModePerm)
	if err == os.ErrInvalid || err == os.ErrPermission {
		return err
	}
	return nil
}